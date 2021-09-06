package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang/glog"
	"github.com/google/go-github/github"
)

type Hook struct {
	Signature string
	Event     string
	Id        string
	Payload   []byte
}

// parse hook and check fields

func ParseHook(secret []byte, request *http.Request, w http.ResponseWriter, eh EventHandler) (*Hook, error) {
	hook := Hook{}

	if hook.Signature = request.Header.Get("x-hub-signature"); len(hook.Signature) == 0 {
		return nil, SignatureError
	}
	if hook.Event = request.Header.Get("x-github-event"); len(hook.Event) == 0 {
		return nil, NoEventError
	}
	if hook.Id = request.Header.Get("x-github-delivery"); len(hook.Event) == 0 {
		return nil, NoEventIdError
	}

	body, err := ioutil.ReadAll(request.Body)
	
	if err != nil {
		return nil, err
	}

	if !Verify(secret, hook.Signature, body) {
		return nil, InvalidSignatureError
	}

	event, err := github.ParseWebHook(github.WebHookType(request), body)
	if err != nil {
		fmt.Println("Could not parse web hook", err)
		glog.Warning("Could not parse web hook %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return nil, GithubParseError
	}

	go func ()  {
		switch event := event.(type) {
		case *github.PushEvent:
			err = eh.HandlePushEvent(event)
		}
	}() // background doing

	return &hook, nil
}