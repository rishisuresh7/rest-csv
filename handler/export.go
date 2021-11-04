package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"rest-csv/factory"
	"rest-csv/response"
)

var viewMap = map[string]string{
	"avehicle": "a_vehicles",
	"bvehicle": "b_vehicles",
	"demands":  "demands",
	"acsfp":    "acsfp",
}

func Export(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		viewKey, ok := vars["viewName"]
		if !ok {
			l.Errorf("Export: could not find 'viewName' from path params")
			response.Error{Error: "bad request"}.ClientError(w)
			return
		}

		viewName, ok := viewMap[strings.ToLower(viewKey)]
		if !ok {
			l.Errorf("Export: invalid value for 'viewName'")
			response.Error{Error: "bad request"}.ClientError(w)
			return
		}

		exporter := f.Exporter()
		res, err := exporter.ExportView(viewName)
		if err != nil {
			l.Errorf("Export: unable to export data: %s", err)
			response.Error{Error: "internal server error"}.ServerError(w)
			return
		}

		w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet;charset=UTF-8")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s.xlsx", viewName))
		w.WriteHeader(200)
		w.Write(res)
	}
}
