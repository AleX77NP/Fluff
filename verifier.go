package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	// "os"
	"strings"
)

// Save your secret in env variable

// var secret = os.Getenv("WEBHOOK_SECRET")
var secret = "secret123"

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
	hex.Decode(actual, []byte(signature[5:]))

	return hmac.Equal(SignBody(secret, body), actual)
}