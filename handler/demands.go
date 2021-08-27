package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"

	"rest-csv/factory"
	"rest-csv/models"
	"rest-csv/response"
)

func GetDemands(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filters := map[string]string{}
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
				case "vehtype":
					filters["veh_type"] = val[0]
				case "squ":
					filters["squadron"] = val[0]
				default:
					l.Errorf("GetDemands: unable to read data : invalid filters")
					response.Error{Error: "invalid request"}.ClientError(w)
					return
				}
			}
		}

		demand := f.Demand()
		res, err := demand.GetDemands(filters)
		if err != nil {
			l.Errorf("GetDemands: unable to read data from demands: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: res}.Send(w)
	}
}

func AddDemands(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload []models.Demand
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			l.Errorf("AddDemands: invalid request payload: %s", err)
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		if len(payload) == 0 {
			l.Errorf("AddDemands: payload cannot be empty")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		category := f.Demand()
		res, err := category.AddDemands(payload)
		if err != nil {
			l.Errorf("AddDemands: unable to write data: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: fmt.Sprintf("%d item(s) added successfully", res)}.Send(w)
	}
}

func DeleteDemands(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload models.Ids
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			l.Errorf("DeleteDemands: invalid request payload")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		if len(payload.Ids) == 0 {
			l.Errorf("DeleteDemands: no ids to delete")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		category := f.Demand()
		res, err := category.DeleteDemands(payload.Ids)
		if err != nil {
			l.Errorf("DeleteDemands: unable to delete data from category: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: fmt.Sprintf("%d item(s) deleted successfully", res)}.Send(w)
	}
}

func UpdateDemands(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload []models.Demand
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			l.Errorf("UpdateDemands: invalid request payload: %s", err)
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		if len(payload) == 0 {
			l.Errorf("UpdateDemands: payload cannot be empty")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		if payload[0].Id == 0 {
			l.Errorf("UpdateDemands: id cannot be empty")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		category := f.Demand()
		res, err := category.UpdateDemands(payload)
		if err != nil {
			l.Errorf("UpdateDemands: unable to delete data from category: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: fmt.Sprintf("%d item(s) updated successfully", res)}.Send(w)
	}
}
