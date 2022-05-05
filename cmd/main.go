package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting server on port 8000")
	server := NewServer()
	log.Fatal(http.ListenAndServe(":8000", server))
}
