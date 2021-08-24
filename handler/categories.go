package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"rest-csv/factory"
	"rest-csv/models"
	"rest-csv/response"
)

func GetVehicles(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		category := f.Category()
		res, err := category.GetVehicles()
		if err != nil {
			l.Errorf("GetCategoryItems: unable to read data from category: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: res}.Send(w)
	}
}

func AddVehicles(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		category := f.Category()
		res, err := category.AddVehicles(payload)
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

		category := f.Category()
		res, err := category.DeleteVehicles(payload.Ids)
		if err != nil {
			l.Errorf("DeleteVehicles: unable to delete data from category: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: fmt.Sprintf("%d item(s) deleted successfully", res)}.Send(w)
	}
}

func UpdateVehicles(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		category := f.Category()
		res, err := category.UpdateVehicles(payload)
		if err != nil {
			l.Errorf("UpdateVehicles: unable to delete data from category: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: fmt.Sprintf("%d item(s) updated successfully", res)}.Send(w)
	}
}
