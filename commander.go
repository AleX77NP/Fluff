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

// clone my repo
func(c Commander) CloneRepository(repoUrl string, ref string) error {
	glog.Infof("Cloning repository: %s in directory: %s", repoUrl, c.dir)

	githubRepoUrl := fmt.Sprintf("%s/%s.git", githubUrl, repoUrl)
	err := exec.Command("git", "clone", githubRepoUrl, c.dir, "--depth", "1").Run()

	if err != nil {
		return fmt.Errorf("Failed to clone repository %s into %s: %v", githubRepoUrl, c.dir, err)
	}

	return nil
}

// pull new changes
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

// if test or run fails, do this, so the old version of app can continie to run
func(c Commander) Revert(head string) error {
	glog.Infof("Undoing commit %s", head)
	cmd := exec.Command("git", "reset", "--hard", "HEAD^")
	cmd.Dir = c.dir
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Failed to revert commit %s, beacuse of: %v", head, err)
	}
	return nil
}

//automated tests
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

// run project 
func (c Commander) Run(runCommand string) error {
	glog.Infof("Starting app...")
	cmd := exec.Command("make", runCommand)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	glog.Infof("Output: %v", string(output))
	return nil
}

// remove project
func(c Commander) Remove() {
	err := os.RemoveAll(c.dir)
	if err != nil {
		glog.Errorf("Failed to clean because %v", err)
	}
}

// cleanup unused stuff
func (c Commander) Cleanup() {
	glog.Infof("Started cleaning up...")
	cmd := exec.Command("make", "fluff-cleanup")
	output, err := cmd.CombinedOutput()
	if err != nil {
		glog.Warningf("Error: %v", err)
	}
	glog.Infof("Output: %v", string(output))
}