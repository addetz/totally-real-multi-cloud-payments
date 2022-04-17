package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Client router and endpoints
func NewClientRouter() *mux.Router{
	r := mux.NewRouter()

	r.HandleFunc("/", RootClient)
	r.HandleFunc("/health", Health)
	r.HandleFunc("/readiness", Readiness)

	return r
}

func RootClient(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for root client handler")
	w.Write([]byte("Welcome to the totally real multi-cloud payments client!"))
}

