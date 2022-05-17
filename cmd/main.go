package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// func main() {
// 	server := NewServer()
// 	log.Fatal(http.ListenAndServe(":8000", server))
// }

// // https://millhouse.dev/posts/graceful-shutdowns-in-golang-with-signal-notify-context
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	server := WrapServerWithLogging(NewServer())

	var srv http.Server
	srv.Addr = ":8000"
	srv.Handler = server

	go srv.ListenAndServe()

	<-ctx.Done()

	stop()
	log.Printf("HTTP server is Shutingdown")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(timeoutCtx); err != nil {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}

// // https://pkg.go.dev/net/http#Server.Shutdown
// func main() {

// 	var srv http.Server

// 	server := NewServer()

// 	idleConnsClosed := make(chan struct{})

// 	go func() {
// 		sigint := make(chan os.Signal, 1)
// 		signal.Notify(sigint, os.Interrupt)

// 		<-sigint

// 		if err := srv.Shutdown(context.Background()); err != nil {
// 			log.Printf("HTTP server Shutdown: %v", err)
// 		}

// 		close(idleConnsClosed)
// 	}()

// 	srv.Addr = ":8000"
// 	srv.Handler = server

// 	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
// 		log.Fatalf("HTTP server ListenAndServe: %v", err)
// 	}
// }
