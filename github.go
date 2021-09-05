package main

import (
	"context"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GithubClient struct {
	accessToken string
	*github.Client
}

func NewGithubClient(accessToken string) *GithubClient{
	accessToken = strings.Trim(accessToken, "\n ")
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	client := github.NewClient(oauth2.NewClient(context.Background(), tokenSource))

	return &GithubClient{
		accessToken,
		client,
	}
}

