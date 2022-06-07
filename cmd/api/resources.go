package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func getHTTPServer() *http.Server {
	r := mux.NewRouter()

	handler := NewHandler()
	r.PathPrefix("/nfe/v1").Handler(handler)
	r.PathPrefix("/metrics").Handler(promhttp.Handler())

	httpAddr := ":" + config.HTTP.Port
	httpServer := &http.Server{
		Addr:    httpAddr,
		Handler: r,
	}

	return httpServer
}
