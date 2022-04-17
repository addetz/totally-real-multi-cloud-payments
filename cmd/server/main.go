package main

import (
	"log"
	"net/http"
	"time"

	"github.com/addetz/totally-real-multi-cloud-payments/handlers"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	CockroachURL = "cockroachdb://roach:payments-r-us@localhost:26257/payments"
)

func main() {
	timeout := 5 * time.Second
	srv := &http.Server{
		Handler:      handlers.NewServerRouter(),
		Addr:         ":4000",
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	// Start Server
	go func() {
		log.Println("Starting server...")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful Shutdown
	handlers.Shutdown(srv)
}

func runMigrate() {
	// run the migration
	m, err := migrate.New(
		"file://db/migrations",
		CockroachURL)
	if err != nil {
		log.Fatal("Migrate new:", err)
	}
	if err := m.Up(); err != nil {
		log.Fatal("Migrate up:", err)
	}

	log.Print("Migration to Cockroach complete.")
}
