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
	router.HandleFunc("/vehicles/{vehicleType}", authorizer.Authorize(handler.GetVehicles(f, l))).Methods("GET")
	router.HandleFunc("/vehicles/{vehicleType}", authorizer.Authorize(handler.AddVehicles(f, l))).Methods("POST")
	router.HandleFunc("/vehicles/{vehicleType}", authorizer.Authorize(handler.UpdateVehicles(f, l))).Methods("PATCH")
	router.HandleFunc("/vehicles/{vehicleType}", authorizer.Authorize(handler.DeleteVehicles(f, l))).Methods("DELETE")

	router.HandleFunc("/demands", authorizer.Authorize(handler.GetDemands(f, l))).Methods("GET")
	router.HandleFunc("/demands", authorizer.Authorize(handler.AddDemands(f, l))).Methods("POST")
	router.HandleFunc("/demands", authorizer.Authorize(handler.UpdateDemands(f, l))).Methods("PATCH")
	router.HandleFunc("/demands", authorizer.Authorize(handler.DeleteDemands(f, l))).Methods("DELETE")

	router.HandleFunc("/acsfp", authorizer.Authorize(handler.GetACSFPItems(f, l))).Methods("GET")
	router.HandleFunc("/acsfp", authorizer.Authorize(handler.AddACSFPItems(f, l))).Methods("POST")
	router.HandleFunc("/acsfp", authorizer.Authorize(handler.UpdateACSFPItems(f, l))).Methods("PATCH")
	router.HandleFunc("/acsfp", authorizer.Authorize(handler.DeleteACSFPItems(f, l))).Methods("DELETE")

	router.HandleFunc("/notifications", authorizer.Authorize(handler.GetNotifications(f, l))).Methods("GET")
	router.HandleFunc("/notifications", authorizer.Authorize(handler.UpdateAlerts(f, l))).Methods("PATCH")
	router.HandleFunc("/alerts", authorizer.Authorize(handler.CreateAlerts(f, l))).Methods("POST")
	router.HandleFunc("/alerts", authorizer.Authorize(handler.ModifyAlerts(f, l))).Methods("PATCH")
	router.HandleFunc("/alerts", authorizer.Authorize(handler.DeleteAlerts(f, l))).Methods("DELETE")
	router.HandleFunc("/alerts", authorizer.Authorize(handler.GetAlerts(f, l))).Methods("GET")
	router.HandleFunc("/tabs", authorizer.Authorize(handler.Tabs(f, l))).Methods("GET")

	router.HandleFunc("/export/{viewName}", authorizer.Authorize(handler.Export(f, l))).Methods("GET")
	router.HandleFunc("/import/{viewName}", authorizer.Authorize(handler.Import(f, l))).Methods("POST")

	router.HandleFunc("/auth", handler.Login(f, l)).Methods("POST")

	return router
}
