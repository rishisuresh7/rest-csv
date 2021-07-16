package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"rest-csv/factory"
	"rest-csv/response"
)

func ListCategories(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		categories := []string{
			"Tanks",
			"Heavy vehicles",
			"Others",
		}

		response.Success{Success: categories}.Send(w)
	}
}
