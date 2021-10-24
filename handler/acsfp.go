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

func GetACSFPItems(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filters := map[string]string{}
		acsfp := f.ACSFP()
		res, err := acsfp.GetItems(filters)
		if err != nil {
			l.Errorf("GetACSFPItems: unable to read data from acsfp items: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: res}.Send(w)
	}
}

func AddACSFPItems(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload []models.ACSFP
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			l.Errorf("AddACSFPItems: unable to decode payload: %s", err)
			response.Error{Error: "invalid request payload"}.ClientError(w)
			return
		}

		if len(payload) == 0 {
			l.Errorf("AddACSFPItems: payload cannot be empty")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		acsfp := f.ACSFP()
		res, err := acsfp.AddItems(payload)
		if err != nil {
			l.Errorf("AddACSFPItems: unable to read data from acsfp items: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: fmt.Sprintf("%d ACSFP item(s) added successfully", res)}.Send(w)
	}
}

func UpdateACSFPItems(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload []models.ACSFP
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			l.Errorf("UpdateACSFPItems: unable to decode payload: %s", err)
			response.Error{Error: "invalid request payload"}.ClientError(w)
			return
		}

		if len(payload) == 0 {
			l.Errorf("UpdateACSFPItems: payload cannot be empty")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		if payload[0].Id == 0 {
			l.Errorf("UpdateACSFPItems: id cannot be empty")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		acsfp := f.ACSFP()
		res, err := acsfp.UpdateItems(payload)
		if err != nil {
			l.Errorf("UpdateACSFPItems: unable to read data from acsfp items: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: fmt.Sprintf("%d ACSFP item(s) updated successfully", res)}.Send(w)
	}
}

func DeleteACSFPItems(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload models.Ids
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			l.Errorf("DeleteACSFPItems: invalid request payload")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		if len(payload.Ids) == 0 {
			l.Errorf("DeleteACSFPItems: no ids to delete")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		acsfp := f.ACSFP()
		res, err := acsfp.DeleteItems(payload.Ids)
		if err != nil {
			l.Errorf("DeleteACSFPItems: unable to delete data from ACSFP: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: fmt.Sprintf("%d ACSFP item(s) deleted successfully", res)}.Send(w)
	}
}