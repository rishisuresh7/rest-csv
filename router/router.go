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
	router.HandleFunc("/categories/{name}", handler.GetCategoryItems(f, l)).Methods("GET")
	router.HandleFunc("/categories/{name}", handler.AddCategoryItem(f, l)).Methods("POST")
	router.HandleFunc("/categories/{name}", handler.UpdateCategoryItem(f, l)).Methods("PATCH")
	router.HandleFunc("/categories/{name}", handler.DeleteCategoryItem(f, l)).Methods("DELETE")

	return router
}
