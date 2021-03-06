package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	handler := WrapHandlerWithLogging(NewHandler())
	handler = WrapHandlerWithMetrics(handler)

	mux := http.NewServeMux()
	mux.Handle("/nfe/v1", handler)
	mux.Handle("/metrics", promhttp.Handler())

	srv := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	go srv.ListenAndServe()

	<-ctx.Done()

	stop()
	log.Printf("HTTP server is Shuttingdown")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(timeoutCtx); err != nil {
		log.Fatalf("Failed to shutdown HTTP server: %v", err)
	}
}
