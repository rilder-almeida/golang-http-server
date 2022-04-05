package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewServer()
	log.Fatal(http.ListenAndServe(":5000", server))
}
