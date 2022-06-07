package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getHTTPServer() *http.Server {
	r := mux.NewRouter()

	r.PathPrefix("/").Handler(nil)

	httpAddr := ":" + config.HTTP.Port
	httpServer := &http.Server{
		Addr:    httpAddr,
		Handler: r,
	}

	return httpServer
}
