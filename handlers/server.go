package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Server router and endpoints
func NewServerRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", RootServer)
	r.HandleFunc("/health", Health)
	r.HandleFunc("/readiness", Readiness)

	return r
}

func RootServer(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for root server handler")
	w.Write([]byte("Welcome to the totally real multi-cloud payments server!"))
}
