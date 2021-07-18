package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	"rest-csv/factory"
	"rest-csv/models"
	"rest-csv/response"
)

func Login(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload models.User
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			l.Errorf("Login: invalid request payload")
			response.Error{Error: "forbidden"}.Forbidden(w)
			return
		}

		auth := f.Auth()
		res, err := auth.Login(payload.Username, payload.Password)
		if err != nil {
			l.Errorf("Login: invalid credentials")
			response.Error{Error: "forbidden"}.Forbidden(w)
			return
		}

		response.Success{Success: res}.Send(w)
	}
}
