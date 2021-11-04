package export

import (
	"bytes"
	"fmt"

	"github.com/tealeg/xlsx"

	"rest-csv/builder"
	"rest-csv/repository"
)

type Exporter interface {
	ExportView(viewName string) ([]byte, error)
}

type exporter struct {
	exportBuilder builder.ExportBuilder
	exec          repository.QueryExecutor
}

func NewExporter(b builder.ExportBuilder, e repository.QueryExecutor) Exporter {
	return &exporter{
		exportBuilder: b,
		exec:          e,
	}
}

func (e *exporter) ExportView(viewName string) ([]byte, error) {
	query := e.exportBuilder.ExportView(viewName)
	sqlRows, err := e.exec.Query(query)
	if err != nil {
		return nil, fmt.Errorf("ExportView: unable to query database: %s", err)
	}

	rows, err := e.exec.ParseRows(sqlRows)
	if err != nil {
		return nil, fmt.Errorf("ExportView: unable to parse rows: %s", err)
	}

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("A VEH")
	if err != nil {
		return nil, fmt.Errorf("ExportView: unable to add a new sheet: %s", err)
	}

	row := sheet.AddRow()
	header := e.getHeader(viewName)
	row.WriteSlice(&header, -1)

	for _, stringRow := range rows {
		row := sheet.AddRow()
		validValues := stringRow[1:]
		row.WriteSlice(&validValues, -1)
	}

	b := bytes.Buffer{}
	_ = file.Write(&b)

	return b.Bytes(), nil
}

func (e *exporter) getHeader(tableName string) []string {
	var res []string
	switch tableName {
	case "a_vehicles":
		res = []string{"SQN", "VEHICLE TYPE", "BA NUMBER", "TYPE", "KILOMETERS", "ENGINE HOURS",
			"EFC", "TM 1", "TM 2", "CMS IN", "CMS OUT", "WORKSHOP IN", "WORKSHOP OUT", "MR 1",
			"MR 2", "FD FIRING", "SERIES INSPECTION", "TRG/OP", "REMARKS"}
	case "b_vehicles":
		res = []string{"SQN", "VEHICLE TYPE", "BA NUMBER", "TYPE", "KILOMETERS", "CMS IN", "CMS OUT",
			"WORKSHOP IN", "WORKSHOP OUT", "REMARKS"}
	case "demands":
		res = []string{"SQN", "VEHICLE TYPE", "BA NUMBER", "TYPE", "EQUIPMENT DEMANDED", "DEPOT",
			"DEMAND NUMBER", "DEMAND DATE", "CONTROL NUMBER", "CONTROL DATE", "STATUS"}
	case "acsfp":
		res = []string{"NAME", "QTY AUTH", "QTY HELD", "REGD NUMBER", "YEAR OF PROC", "COST",
			"QTY SER", "FWD TO", "DEMAND DETAILS", "REMARKS"}
	}

	return res
}
