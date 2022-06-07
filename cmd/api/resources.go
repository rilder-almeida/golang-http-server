package main

import (
	"net/http"

	"github.com/arquivei/foundationkit/app"
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

	app.RegisterShutdownHandler(
		&app.ShutdownHandler{
			Name:     "http_server",
			Priority: shutdownPriorityHTTP,
			Handler:  httpServer.Shutdown,
			Policy:   app.ErrorPolicyAbort,
		})

	// TODO RETORNAR O SERVER COM O WRAPPER DE METRICS

	return httpServer
}
