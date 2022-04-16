package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	CockroachURL = "cockroachdb://roach:payments-r-us@localhost:26257/payments"
)

func main() {
	log.Print("Welcome to the totally real multi-cloud payments engine!")

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
