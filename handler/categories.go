package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"rest-csv/factory"
	"rest-csv/response"
)

func ListCategories(f factory.Factory, _ *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		category := f.Category("")

		response.Success{Success: category.GetCategories()}.Send(w)
	}
}
