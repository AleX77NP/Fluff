package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"os"
	"strings"
)

// Save your secret in env variable
var secret = os.Getenv("WEBHOOK_SECRET")

func SignBody(secret, body []byte) []byte {
	result := hmac.New(sha1.New, secret)
	result.Write(body)
	return []byte(result.Sum(nil))
}

func Verify(secret []byte, signature string, body []byte) bool {
	const prefix = "sha1="
	const length = 45 

	if(len(signature) != length || !strings.HasPrefix(signature, prefix)) {
		return false
	}

	actual := make([]byte, 20)

	return hmac.Equal(SignBody(secret, body), actual)
}