package handler

import (
	"net/http"

	"github.com/gorilla/mux"
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

func GetCategoryItems(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name, ok := vars["name"]
		if !ok {
			l.Errorf("GetCategory: could not read name from path params")
			response.Error{Error: "bad request"}.ClientError(w)
			return
		}

		category := f.Category(name)
		res, err := category.GetCategoryItems()
		if err != nil {
			l.Errorf("GetCategory: unable to read data from category: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: res}.Send(w)
	}
}
