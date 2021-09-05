package main

import (
	"log"
	"net/http"
	"io"
	"fmt"
	"github.com/golang/glog"
)

var (
	token = "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"
	repo = "AleX77NP/Express-test-app.git"
)

func main() {

	client := NewGithubClient(token)

	eventHandler := EventHandler{
		repo: repo,
		client: client,
	}

	http.HandleFunc("/hook", HookHandler(eventHandler))
	log.Fatal(http.ListenAndServe(":7654", nil))
}

func HookHandler(eh EventHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		hook, err := ParseHook([]byte(secret), req, w, eh)

		w.Header().Set("Content-type", "application/json")

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("Could not parse web hook", err)
			log.Printf("Not able to process hook ('%s')", err)
			return
		}

		glog.Infof("web hook event %v", hook.Event)
		
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "{}")
		return
	}
}
