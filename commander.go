package main

import (
	"fmt"
	"os/exec"

	"github.com/golang/glog"
)

type Commander struct {
	dir string
}

func NewCommander() Commander {
	commander := Commander{}
	commander.dir = getDirectory()
	return commander
}

func(c Commander) CloneRepository(repoUrl string, ref string) error {
	glog.Infof("Cloning repository: %s in directory: %s", repoUrl, c.dir)

	githubRepoUrl := fmt.Sprintf("%s/%s.git", githubUrl, repoUrl)
	err := exec.Command("git", "clone", githubRepoUrl, c.dir, "--depth", "1").Run()

	if err != nil {
		return fmt.Errorf("Failed to clone repository %s into %s: %v", githubRepoUrl, c.dir, err)
	}

	return nil
}