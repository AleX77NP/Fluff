package main

import (
	"io/ioutil"
	"net/http"
)

type Hook struct {
	Signature string
	Event     string
	Id        string
	Payload   []byte
}

// parse hook and check fields

func ParseHook(secret []byte, request *http.Request) (*Hook, error) {
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

	return &hook, nil
}