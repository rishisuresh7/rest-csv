package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Success struct {
	Success interface{} `json:"success"`
}

type Error struct {
	Error interface{} `json:"error"`
}

func (s Success) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(s)
	if err != nil {
		return fmt.Errorf("Send: unable to encode to JSON: %s", err)
	}

	return nil
}

func (e Error) ClientError(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	err := json.NewEncoder(w).Encode(e)
	if err != nil {
		return fmt.Errorf("ClientError: unable to encode to JSON: %s", err)
	}

	return nil
}

func (e Error) Forbidden(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusForbidden)
	err := json.NewEncoder(w).Encode(e)
	if err != nil {
		return fmt.Errorf("Forbidden: unable to encode to JSON: %s", err)
	}

	return nil
}

func (e Error) ServerError(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	err := json.NewEncoder(w).Encode(e)
	if err != nil {
		return fmt.Errorf("ServerError: unable to encode to JSON: %s", err)
	}

	return nil
}
