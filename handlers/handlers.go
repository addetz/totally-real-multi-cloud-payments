package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func New() *mux.Router{
	// Create Server and Route Handlers
	r := mux.NewRouter()

	r.HandleFunc("/", Root)
	r.HandleFunc("/health", Health)
	r.HandleFunc("/readiness", Readiness)

	return r
}

func Root(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request for root handler\n")
	w.Write([]byte("Welcome to the totally real multi-cloud payments engine!"))
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func Readiness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
