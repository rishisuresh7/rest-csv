package handler

import (
	"net/http"

	"rest-csv/response"
)

func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.Success{Success: "I'm alive"}.Send(w)
	}
}
