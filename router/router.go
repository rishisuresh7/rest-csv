package router

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"rest-csv/factory"
	"rest-csv/handler"
)

func NewRouter(f factory.Factory, l *logrus.Logger) *mux.Router {
	router := mux.NewRouter()
	authorizer := f.NewJWTAuth()

	router.HandleFunc("/health", handler.Health()).Methods("GET")
	router.HandleFunc("/categories", authorizer.Authorize(handler.ListCategories(f, l))).Methods("GET")
	router.HandleFunc("/categories", authorizer.Authorize(handler.DeleteItems(f, l))).Methods("DELETE")
	router.HandleFunc("/demands", authorizer.Authorize(handler.ListDemands(f, l))).Methods("GET")

	router.HandleFunc("/categories/{name}", authorizer.Authorize(handler.GetCategoryItems(f, l))).Methods("GET")
	router.HandleFunc("/categories/{name}", authorizer.Authorize(handler.AddCategoryItem(f, l))).Methods("POST")
	router.HandleFunc("/categories/{name}", authorizer.Authorize(handler.UpdateCategoryItem(f, l))).Methods("PATCH")

	router.HandleFunc("/auth", handler.Login(f, l)).Methods("POST")

	return router
}
