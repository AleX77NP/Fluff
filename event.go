package main

import (
	"flag"

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

	head, fName := *event.HeadCommit.ID, *event.Repo.FullName

	glog.Infof("Head and Repo name %v %v", head, fName)

	commander := NewCommander()

	commander.CloneRepository(fName, event.GetRef())

	return nil
}