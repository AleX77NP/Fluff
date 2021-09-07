package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/golang/glog"
)

var (
	token1 = os.Getenv("WEBHOOK_ACCESS_TOKEN")
	token = "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"
	//repo = "AleX77NP/node-test"
)

func main() {

	flag.Set("logtostderr", "true")
	flag.Parse()

	client := NewGithubClient(token)

	eventHandler := EventHandler{
		//repo: repo,
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
