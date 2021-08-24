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
	router.HandleFunc("/categories", authorizer.Authorize(handler.GetVehicles(f, l))).Methods("GET")
	router.HandleFunc("/categories", authorizer.Authorize(handler.AddVehicles(f, l))).Methods("POST")
	router.HandleFunc("/categories", authorizer.Authorize(handler.UpdateVehicles(f, l))).Methods("PATCH")
	router.HandleFunc("/categories", authorizer.Authorize(handler.DeleteVehicles(f, l))).Methods("DELETE")

	router.HandleFunc("/demands", authorizer.Authorize(handler.GetDemands(f, l))).Methods("GET")
	router.HandleFunc("/demands", authorizer.Authorize(handler.AddDemands(f, l))).Methods("POST")
	router.HandleFunc("/demands", authorizer.Authorize(handler.UpdateDemands(f, l))).Methods("PATCH")
	router.HandleFunc("/demands", authorizer.Authorize(handler.DeleteDemands(f, l))).Methods("DELETE")

	router.HandleFunc("/auth", handler.Login(f, l)).Methods("POST")

	return router
}
