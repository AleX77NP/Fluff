package main

import "fmt"

var (
	githubUrl = "https://github.com"
)

func getDirectoryStaging() string {
	return fmt.Sprintf("/tmp/staging")
}

func getDirectoryProduction() string {
	return fmt.Sprintf("/tmp/production")
}
