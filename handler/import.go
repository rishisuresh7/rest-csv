package handler

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"rest-csv/factory"
	"rest-csv/response"
)

func Import(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		viewKey, ok := vars["viewName"]
		if !ok {
			l.Errorf("Import: could not find 'viewName' from path params")
			response.Error{Error: "bad request"}.ClientError(w)
			return
		}

		viewName, ok := viewMap[strings.ToLower(viewKey)]
		if !ok {
			l.Errorf("Import: invalid value for 'viewName'")
			response.Error{Error: "bad request"}.ClientError(w)
			return
		}

		r.ParseMultipartForm(2 << 10)

		file, fileHeader, err := r.FormFile("file")
		if err != nil || file == nil {
			l.Errorf("Import: unable to read data from request: %s", err)
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		if fileHeader.Header.Get("Content-Type") != "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" ||
			!(strings.HasSuffix(fileHeader.Filename, ".xlsx") || strings.HasSuffix(fileHeader.Filename, ".xls")) {
			l.Errorf("Import: invalid content type: %s", fileHeader.Header.Get("Content-Type"))
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		importer := f.Importer(viewName)
		res, err := importer.ImportView(viewName, file.(*os.File).Name())
		if err != nil {
			l.Errorf("Import: unable to Import data: %s", err)
			response.Error{Error: "internal server error"}.ServerError(w)
			return
		}

		response.Success{Success: fmt.Sprintf("Inserted %d row(s) successfully", res)}.Send(w)
	}
}
