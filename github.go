package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang/glog"
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

func(c *GithubClient) PostStatus(fName, head, job string, status, target string) error {
	// Request url
	url := fmt.Sprintf("https://api.github.com/repos/%s/statuses/%s", fName, head)

	// Request body
	data := map[string]string{}
	data["context"] = "ci-cd/fluff" + target
	data["state"] = status
	data["target_url"] = c.getBaseUrl() + job
	data["description"] = "Fluff test"
	dataToSend, _ := json.Marshal(data)

	// send request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(dataToSend))
	if err != nil {
		return fmt.Errorf("Failed to create status create request, error: %v", err)
	}

	// get response
	res, err := c.Do(context.Background(), req, nil)
	if err != nil {
		return err
	} else if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("Failed to post status, error: %v with status code: %v", err, res.StatusCode)
	}

	glog.Infof("Success setting status of %s/%s to %s link: %s", fName, status, target, data["target_url"])
	return nil
}


func(c *GithubClient) getBaseUrl() string {
	if len(c.accessToken) == 0 {
		return fmt.Sprintf("https://github.com")
	}
	return fmt.Sprintf("https://%s@github.com", c.accessToken)
}

