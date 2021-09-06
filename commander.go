package main

import (
	"fmt"
	"os"
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

func(c Commander) Pull() error {
	glog.Infof("Pulling from master")
	cmd := exec.Command("git", "pull")
	cmd.Dir = c.dir
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Failed to pull, beacuse of: %v", err)
	}
	return nil
}

func(c Commander) Revert(head string) error {
	glog.Infof("Undoing commit %s", head)
	cmd := exec.Command("git", "reset", "--hard", "HEAD^")
	cmd.Dir = c.dir
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Failed to revert commit %s, beacuse of: %v", head, err)
	}
	return nil
}

func(c Commander) TestRepository() error {
	glog.Infof("Testing app with command %s", c.dir)
	cmd := exec.Command("make", "fluff-test")
	//cmd.Dir = c.dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	glog.Infof("Test output: %v", string(output))
	return nil
}

func(c Commander) Clean() {
	err := os.RemoveAll(c.dir)
	if err != nil {
		glog.Errorf("Failed to clean because %v", err)
	}
}