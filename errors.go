package main

import "errors"

var (
	SignatureError = errors.New("No signature provided!")
	NoEventError = errors.New("No event!")
	NoEventIdError = errors.New("No event id!")
	InvalidSignatureError = errors.New("Invalid signature error!")
)