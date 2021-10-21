package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"

	"rest-csv/factory"
	"rest-csv/models"
	"rest-csv/response"
)

func GetNotifications(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		alerts := f.Alerts()
		res, err := alerts.GetNotifications()
		if err != nil {
			l.Errorf("GetNotifications: unable to get notifications: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		queryParams := r.URL.Query()
		if strings.ToLower(queryParams.Get("type")) == "count" {
			response.Success{Success: len(res)}.Send(w)
			return
		}

		response.Success{Success: res}.Send(w)
	}
}

func CreateAlerts(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload models.Alert
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			l.Errorf("CreateAlerts: unable to decode payload: %s", err)
			response.Error{Error: "invalid request payload"}.ClientError(w)
			return
		}

		if payload.Name == "" || payload.BaNo == "" || payload.AlertField == "" || payload.LastValue == "" || payload.NextValue == "" {
			l.Errorf("CreateAlerts: invalid payload, empty fields")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		alerts := f.Alerts()
		err = alerts.CreateAlert(payload)
		if err != nil {
			l.Errorf("CreateAlerts: unable to create alerts: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: "alert created successfully"}.Send(w)
	}
}

func GetAlerts(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		alerts := f.Alerts()
		res, err := alerts.GetAlerts()
		if err != nil {
			l.Errorf("GetAlerts: unable to get alerts: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: res}.Send(w)
	}
}

func UpdateAlerts(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload models.Notification
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			l.Errorf("UpdateAlerts: unable to decode payload: %s", err)
			response.Error{Error: "invalid request payload"}.ClientError(w)
			return
		}

		alerts := f.Alerts()
		err = alerts.UpdateAlert(payload)
		if err != nil {
			l.Errorf("UpdateAlerts: unable to update alert: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: "Alert updated successfully"}.Send(w)
	}
}

func Tabs(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.Success{Success: []string{"Vehicle A", "Vehicle B", "Others", "Demands"}}.Send(w)
	}
}
