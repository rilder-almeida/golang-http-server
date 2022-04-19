package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewServer()
	log.Fatal(http.ListenAndServe(":8000", server))
}
