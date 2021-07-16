package router

import (
	"github.com/gorilla/mux"

	"rest-csv/handler"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", handler.Health()).Methods("GET")

	return router
}
