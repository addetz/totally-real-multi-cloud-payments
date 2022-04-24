package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/addetz/totally-real-multi-cloud-payments/models"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func Readiness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func Shutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}

func NewPayment(w http.ResponseWriter, req *http.Request) {
	p, err := models.NewPaymentRequest(req)
	if err != nil {
		r := models.Response{
			StatusCode: http.StatusBadRequest,
			Body:       nil,
			Error:      fmt.Errorf("error decoding body:%s", err.Error()),
		}
		r.Write(w)
	}
	r := models.Response{
		StatusCode: http.StatusOK,
		Body: models.PaymentResponse{
			ID:      p.TransactionID,
			Message: "payment request received",
		},
		Error: nil,
	}
	r.Write(w)
	return
}
