package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hook", HookHandler)
	log.Fatal(http.ListenAndServe(":7654", nil))
}