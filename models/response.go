package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	StatusCode int
	Body       interface{}
	Error      error
}

func (r Response) Write(w http.ResponseWriter) {
	bytes, err := r.getBody()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprint("error during response:", err)))
		return
	}
	w.WriteHeader(r.StatusCode)
	w.Write(bytes)
}

func (r Response) getBody() ([]byte, error) {
	if r.Error != nil {
		return json.Marshal(r.Error)
	}
	return json.Marshal(r.Body)
}
