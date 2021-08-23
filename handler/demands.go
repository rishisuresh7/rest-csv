package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"rest-csv/factory"
	"rest-csv/response"
)

func ListDemands(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		demand := f.Demand()
		res, err := demand.ListDemands()
		if err != nil {
			l.Errorf("ListDemands: unable to read data from demands: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Success: res}.Send(w)
	}
}
