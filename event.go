package main

import (
	"flag"
	"fmt"

	"github.com/golang/glog"
	"github.com/google/go-github/github"
	"strings"
)

var (
	MasterRefStaging string
	masterRefProduction string
)

func init() {
	flag.StringVar(&MasterRefStaging, "master-ref", "refs/heads/master", "The ref with post-commit targets. Defaults to refs/heads/master")
	flag.StringVar(&masterRefProduction, "tag-ref", "refs/tags/", "Tag ref with production ready commit.")
}

type EventHandler struct {
	client *GithubClient
	//repo      string
	masterRef string
}

func (e *EventHandler) HandlePushEvent(event *github.PushEvent) error {
	glog.Infof("Recieved push event")

	var directory string
	var state string
	var testState string

	if err := checkPush(event); err != nil {
		return err
	}

	head, fName := *event.HeadCommit.ID, *event.Repo.FullName

	glog.Infof("Head and Repo name %v %v", head, fName)

	if event.GetRef() != MasterRefStaging {
		if (strings.Contains(event.GetRef(), masterRefProduction)) {
			glog.Infof("Ready for production build...")
			directory = getDirectoryProduction()
			state = "fluff-run-prod"
			testState = "fluff-test-prod"
		} else {
			glog.Infof("Not a master or tag ref, but %s", event.GetRef())
			return nil
		}
	} else {
		directory = getDirectoryStaging()
		state = "fluff-run-stag"
		testState = "fluff-test-stag"
	}

	commander := NewCommander(directory)

	defer commander.Cleanup()

	commander.CloneRepository(fName, event.GetRef())

	err := e.client.PostStatus(fName, head, head, "pending", "fluff-ci/cd-test")
	if err != nil {
		glog.Warning("Failed to create pending status, error: %v", err)
	}

	err = commander.Pull()
	if err != nil {
		e.client.PostStatus(fName, head, head, "failure", "fluff-ci/cd-test")
		glog.Infof("Failed to pull from master, error: %v", err)
	}

	err = commander.TestRepository(testState)
	if err != nil {
		glog.Warningf("App test failed, error: %v", err)
		e.client.PostStatus(fName, head, head, "failure", "fluff-ci/cd-test")
		err = commander.Revert(head)
		if err != nil {
			glog.Warningf("Failed to revert, error: %v", err)
			return nil
		}
		return nil
	}

	err = commander.Run(state)
	if err != nil {
		glog.Warningf("Failed to run app, error: %v", err)
		e.client.PostStatus(fName, head, head, "failure", "fluff-ci/cd-test")
		err = commander.Revert(head)
		if err != nil {
			glog.Warningf("Failed to revert, error: %v", err)
			return nil
		}
		err = commander.Run(state)
		if err != nil {
			glog.Warningf("Failed to run app, error: %v", err)
			return nil
		}
		return nil
	}

	e.client.PostStatus(fName, head, head, "success", "fluff-ci/cd-test")

	return nil
}

func checkPush(event *github.PushEvent) error {
	if event.Repo == nil {
		return fmt.Errorf("Missing PushEvent.Repo")
	}
	if event.HeadCommit == nil {
		return fmt.Errorf("Missing PushEvent head commit")
	}
	if event.Repo.FullName == nil {
		return fmt.Errorf("Missing PushEvent repo full name")
	}
	if event.HeadCommit.ID == nil {
		return fmt.Errorf("Missing PushEvent head commit ID")
	}
	return nil
}
