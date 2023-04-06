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

	srv := http.Server{
		Addr:    ":" + *port,
		Handler: mux,
	}

	fmt.Println("Serving application on port", *port)
	log.Fatal(srv.ListenAndServe())
}
