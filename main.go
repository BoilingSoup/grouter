package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "4000", "The port to serve the application.")
	flag.Parse()

	mux := NewServeMux()

	mux.Get("/", http.HandlerFunc(homeGet))
	mux.Post("/", http.HandlerFunc(homePost))
	mux.Put("/", http.HandlerFunc(homePut))
	mux.Patch("/", http.HandlerFunc(homePatch))
	mux.Delete("/", http.HandlerFunc(homeDelete))

	mux.Post("/resource", http.HandlerFunc(resourcePost))

	srv := http.Server{
		Addr:    ":" + *port,
		Handler: mux,
	}

	fmt.Println("Serving application on port", *port)
	log.Fatal(srv.ListenAndServe())
}
