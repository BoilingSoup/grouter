package main

import "net/http"

func homeGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You sent a GET request to the HOME route.\n"))
}

func homePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You sent a POST request to the HOME route.\n"))
}

func homePut(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You sent a PUT request to the HOME route.\n"))
}

func homePatch(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You sent a PATCH request to the HOME route.\n"))
}

func homeDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You sent a DELETE request to the HOME route.\n"))
}

func resourcePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You sent a POST request for a RESOURCE.\n"))
}
