package main

import (
	"github.com/golang/glog"
	"github.com/google/go-github/github"
)

type EventHandler struct {
	client *GithubClient
	repo string
	masterRef string
}

func(e *EventHandler) HandlePushEvent(event *github.PushEvent) error {
	glog.Infof("Recieved push event")

	head, fName := *event.HeadCommit.ID, *event.Repo.FullName

	glog.Infof("Head and Repo name %v %v", head, fName)

	return nil
}