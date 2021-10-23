package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"rest-csv/constant"
	"rest-csv/factory"
	"rest-csv/models"
	"rest-csv/response"
	"rest-csv/utility"
)

var vehicleTypes = []string{constant.AVehicle, constant.BVehicle, constant.Others}

func GetVehicles(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		vehicleType, ok := vars["vehicleType"]
		if !ok {
			l.Errorf("GetVehicles: 'vehicleType' not found in path param")
			response.Error{Error: "bad request"}.ClientError(w)
			return
		}

		if !utility.CheckList(vehicleTypes, strings.ToLower(vehicleType)) {
			l.Errorf("GetVehicles: invalid vehicle type")
			response.Error{Error: "bad request"}.ClientError(w)
			return
		}

		filters := map[string]string{
			"vehicleType": strings.ToLower(vehicleType),
		}
		queries := r.URL.Query()
		if len(queries) > 0 {
			for key, val := range queries {
				if len(val) < 1 {
					continue
				}

				if strings.ToLower(val[0]) == "all" {
					continue
				}

				switch strings.ToLower(key) {
				case "squ":
					filters["squadron"] = val[0]
				case "search":
					filters["search"] = val[0]
				default:
					l.Errorf("GetVehicles: unable to read data : invalid filters")
					response.Error{Error: "invalid request"}.ClientError(w)
					return
				}
			}
		}

		vehicle := f.Vehicles(strings.ToLower(vehicleType))
		res, err := vehicle.GetVehicles(filters)
		if err != nil {
			l.Errorf("GetCategoryItems: unable to read data from vehicle: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: res}.Send(w)
	}
}

func AddVehicles(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		vehicleType, ok := vars["vehicleType"]
		if !ok {
			l.Errorf("AddVehicles: 'vehicleType' not found in path param")
			response.Error{Error: "bad request"}.ClientError(w)
			return
		}

		if !utility.CheckList(vehicleTypes, strings.ToLower(vehicleType)) {
			l.Errorf("AddVehicles: invalid vehicle type")
			response.Error{Error: "bad request"}.ClientError(w)
			return
		}

		var payload []models.Vehicle
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			l.Errorf("AddVehicles: invalid request payload: %s", err)
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		if len(payload) == 0 {
			l.Errorf("AddVehicles: payload cannot be empty")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		vehicle := f.Vehicles(strings.ToLower(vehicleType))
		res, err := vehicle.AddVehicles(payload)
		if err != nil {
			l.Errorf("AddVehicles: unable to write data: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: fmt.Sprintf("%d item(s) added successfully", res)}.Send(w)
	}
}

func DeleteVehicles(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		vehicleType, ok := vars["vehicleType"]
		if !ok {
			l.Errorf("DeleteVehicles: 'vehicleType' not found in path param")
			response.Error{Error: "bad request"}.ClientError(w)
			return
		}

		if !utility.CheckList(vehicleTypes, strings.ToLower(vehicleType)) {
			l.Errorf("DeleteVehicles: invalid vehicle type")
			response.Error{Error: "bad request"}.ClientError(w)
			return
		}

		var payload models.Ids
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			l.Errorf("DeleteVehicles: invalid request payload")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		if len(payload.Ids) == 0 {
			l.Errorf("DeleteVehicles: no ids to delete")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		vehicle := f.Vehicles(strings.ToLower(vehicleType))
		res, err := vehicle.DeleteVehicles(payload.Ids)
		if err != nil {
			l.Errorf("DeleteVehicles: unable to delete data from vehicle: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: fmt.Sprintf("%d item(s) deleted successfully", res)}.Send(w)
	}
}

func UpdateVehicles(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		vehicleType, ok := vars["vehicleType"]
		if !ok {
			l.Errorf("UpdateVehicles: 'vehicleType' not found in path param")
			response.Error{Error: "bad request"}.ClientError(w)
			return
		}

		if !utility.CheckList(vehicleTypes, strings.ToLower(vehicleType)) {
			l.Errorf("UpdateVehicles: invalid vehicle type")
			response.Error{Error: "bad request"}.ClientError(w)
			return
		}

		var payload []models.Vehicle
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			l.Errorf("UpdateVehicles: invalid request payload: %s", err)
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		if len(payload) == 0 {
			l.Errorf("UpdateVehicles: payload cannot be empty")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		if payload[0].Id == 0 {
			l.Errorf("UpdateVehicles: id cannot be empty")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		vehicle := f.Vehicles(strings.ToLower(vehicleType))
		res, err := vehicle.UpdateVehicles(payload)
		if err != nil {
			l.Errorf("UpdateVehicles: unable to delete data from vehicle: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: fmt.Sprintf("%d item(s) updated successfully", res)}.Send(w)
	}
}
