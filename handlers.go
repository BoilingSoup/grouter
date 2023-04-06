package main

import "net/http"

func homeGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You sent a GET request to the HOME route.\n"))
}

func homePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You sent a POST request to the HOME route.\n"))
}
