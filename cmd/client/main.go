package main

import (
	"log"
	"net/http"
	"time"

	"github.com/addetz/totally-real-multi-cloud-payments/handlers"
)

func main() {
	timeout := 5 * time.Second
	srv := &http.Server{
		Handler:      handlers.NewClientRouter(),
		Addr:         ":8080",
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	// Start Server
	go func() {
		log.Println("Starting client...")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful Shutdown
	handlers.Shutdown(srv)
}
