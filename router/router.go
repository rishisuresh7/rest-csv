package router

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"rest-csv/factory"
	"rest-csv/handler"
)

func NewRouter(f factory.Factory, l *logrus.Logger) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", handler.Health()).Methods("GET")
	router.HandleFunc("/categories", handler.ListCategories(f, l)).Methods("GET")

	return router
}
