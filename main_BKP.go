package main

import (
	"log"
	"net/http"
)

func main_BKP() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}
