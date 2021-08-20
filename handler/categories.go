package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"rest-csv/factory"
	"rest-csv/models"
	"rest-csv/response"
)

func ListCategories(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		category := f.Category("")
		res, err := category.GetCategoryItems()
		if err != nil {
			l.Errorf("GetCategoryItems: unable to read data from category: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: res}.Send(w)
	}
}

func GetCategoryItems(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name, ok := vars["name"]
		if !ok {
			l.Errorf("GetCategoryItems: could not read name from path params")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		category := f.Category(name)
		res, err := category.GetCategoryItems()
		if err != nil {
			l.Errorf("GetCategoryItems: unable to read data from category: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: res}.Send(w)
	}
}

func AddCategoryItem(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name, ok := vars["name"]
		if !ok {
			l.Errorf("AddCategoryItem: could not read name from path params")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		var payload []models.Item
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			l.Errorf("AddCategoryItem: invalid request payload")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		if len(payload) == 0 {
			l.Errorf("AddCategoryItem: payload cannot be empty")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		category := f.Category(name)
		err = category.AddCategoryItem(payload)
		if err != nil {
			l.Errorf("AddCategoryItem: unable to write data into category: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: "item added successfully"}.Send(w)
	}
}

func DeleteCategoryItem(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name, ok := vars["name"]
		if !ok {
			l.Errorf("DeleteCategoryItem: could not read name from path params")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		var payload models.Ids
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			l.Errorf("DeleteCategoryItem: invalid request payload")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		if len(payload.Ids) == 0 {
			l.Errorf("DeleteCategoryItem: no ids to delete")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		category := f.Category(name)
		err = category.DeleteCategoryItem(payload.Ids)
		if err != nil {
			l.Errorf("DeleteCategoryItem: unable to delete data from category: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: "item deleted successfully"}.Send(w)
	}
}

func UpdateCategoryItem(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name, ok := vars["name"]
		if !ok {
			l.Errorf("UpdateCategoryItem: could not read name from path params")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		var payload []models.Item
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			l.Errorf("UpdateCategoryItem: invalid request payload")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		if len(payload) == 0 {
			l.Errorf("UpdateCategoryItem: id cannot be empty")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		category := f.Category(name)
		err = category.UpdateCategoryItem(payload)
		if err != nil {
			l.Errorf("UpdateCategoryItem: unable to delete data from category: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: "item updated successfully"}.Send(w)
	}
}
