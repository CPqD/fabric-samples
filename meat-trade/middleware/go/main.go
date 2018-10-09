package main

import (
	"log"
	"net/http"
)

// main method: makes route and instantiate the server
func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":7500", router))
}
