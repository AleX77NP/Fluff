package main

import "fmt"

var (
	githubUrl = "https://github.com"
)

func getDirectory() string {
	return fmt.Sprintf("/tmp/project")
}
