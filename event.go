package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/google/go-github/github"
)

var (
	MasterRef string
)

func init() {
	flag.StringVar(&MasterRef, "master-ref", "refs/heads/master", "The ref with post-commit targets. Defaults to refs/heads/master")
}

type EventHandler struct {
	client *GithubClient
	repo string
	masterRef string
}

func(e *EventHandler) HandlePushEvent(event *github.PushEvent) error {
	glog.Infof("Recieved push event")

	if err := checkPush(event); err != nil {
		return err
	}

	head, fName := *event.HeadCommit.ID, *event.Repo.FullName

	glog.Infof("Head and Repo name %v %v", head, fName)

	commander := NewCommander()

	commander.CloneRepository(fName, event.GetRef())

	if MasterRef != event.GetRef() {
		glog.Infof("Not a master ref, but %s", event.GetRef())
		return nil
	}

	commander.Pull()


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