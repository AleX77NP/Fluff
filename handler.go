package main

import (
	"io"
	"log"
	"net/http"
)


func HookHandler(w http.ResponseWriter, r *http.Request) {
	hook, err := ParseHook([]byte(secret), r)

	w.Header().Set("Content-type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Not able to process hook ('%s')", err)
		return
	}

	log.Printf("Recieved %s", hook.Event)
	
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "{}")
	return
}