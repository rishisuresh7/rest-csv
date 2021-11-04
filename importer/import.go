package importer

import (
	"fmt"
	"strconv"

	"github.com/tealeg/xlsx"

	"rest-csv/acsfp"
	"rest-csv/demand"
	"rest-csv/models"
	"rest-csv/repository"
	"rest-csv/vehicle"
)

type Importer interface {
	ImportView(viewName, fileName string) (int64, error)
}

type importer struct {
	inserter interface{}
	exec     repository.QueryExecutor
}

func NewImporter(i interface{}, e repository.QueryExecutor) Importer {
	return &importer{
		exec:     e,
		inserter: i,
	}
}

func (i *importer) ImportView(viewName, fileName string) (int64, error) {
	xlFile, err := xlsx.OpenFile(fileName)
	if err != nil {
		return -1, fmt.Errorf("ImportView: unable to open excel file: %s", err)
	}

	xlSlice, err := xlFile.ToSlice()
	if err != nil {
		return -1, fmt.Errorf("ImportView: unable to convert excel file: %s", err)
	}

	if len(xlSlice[0]) < 1 {
		return -1, fmt.Errorf("ImportView: invalid excel file: missing data")
	}

	var res int64
	switch viewName {
	case "a_vehicles":
		vehicles, err := parseVehicles("A", xlSlice[0][1:])
		if err != nil {
			return -1, fmt.Errorf("ImportView: invalid data: %s", err)
		}

		res, err = i.inserter.(vehicle.Vehicle).AddVehicles(vehicles)
	case "b_vehicles":
		vehicles, err := parseVehicles("B", xlSlice[0][1:])
		if err != nil {
			return -1, fmt.Errorf("ImportView: invalid data: %s", err)
		}

		res, err = i.inserter.(vehicle.Vehicle).AddVehicles(vehicles)
	case "demands":
		demands, err := parseDemands(xlSlice[0][1:])
		if err != nil {
			return -1, fmt.Errorf("ImportView: invalid data: %s", err)
		}

		res, err = i.inserter.(demand.Demand).AddDemands(demands)
	case "acsfp":
		items, err := parseACSFPItems(xlSlice[0][1:])
		if err != nil {
			return -1, fmt.Errorf("ImportView: invalid data: %s", err)
		}

		res, err = i.inserter.(acsfp.ACSFP).AddItems(items)
	}

	if err != nil {
		return -1, fmt.Errorf("ImportView: unable to insert data: %s", err)
	}

	return res, nil
}

func parseDemands(rows [][]string) ([]models.Demand, error) {
	if len(rows[0]) != 11 {
		return nil, fmt.Errorf("missing demands columns")
	}

	demands := make([]models.Demand, 0)
	for _, row := range rows {
		item := models.Demand{
			Item: models.Item{
				Sqn:         row[0],
				VehicleType: row[1],
				BaNo:        row[2],
				Type:        row[3],
			},
			EquipmentDemanded: row[4],
			Depot:             row[5],
			DemandNumber:      row[6],
			DemandDate:        row[7],
			ControlNumber:     row[8],
			ControlDate:       row[9],
			Status:            row[10],
		}

		demands = append(demands, item)
	}

	return demands, nil
}

func parseACSFPItems(rows [][]string) ([]models.ACSFP, error) {
	if len(rows[0]) != 10 {
		return nil, fmt.Errorf("missing ACSFP columns")
	}

	items := make([]models.ACSFP, 0)
	for _, row := range rows {
		item := models.ACSFP{
			Name:             row[0],
			QuantityAuth:     stringToInteger(row[1]),
			QuantityHeld:     stringToInteger(row[2]),
			RegisteredNumber: row[3],
			YearOfProc:       stringToInteger(row[4]),
			Cost:             stringToFloat(row[5]),
			QuantityServed:   stringToInteger(row[6]),
			ForwardTo:        row[7],
			DemandDetails:    row[8],
			Remarks:          row[9],
		}

		items = append(items, item)
	}

	return items, nil
}

func parseVehicles(vehicleType string, rows [][]string) ([]models.Vehicle, error) {
	if len(rows[0]) < 5 {
		return nil, fmt.Errorf("missing vehicle columns")
	}

	vehicles := make([]models.Vehicle, 0)
	for _, row := range rows {
		veh := models.Vehicle{
			Item: models.Item{
				Sqn:         row[0],
				VehicleType: row[1],
				BaNo:        row[2],
				Type:        row[3],
			},
			Kilometers: stringToInteger(row[4]),
		}
		if vehicleType == "A" {
			if len(rows[0]) != 19 {
				return nil, fmt.Errorf("missing A vehicle columns")
			}

			veh.EngineHours = stringToInteger(row[5])
			veh.Efc = stringToInteger(row[6])
			veh.TM1 = row[7]
			veh.TM2 = row[8]
			veh.CMSIn = row[9]
			veh.CMSOut = row[10]
			veh.WorkshopIn = row[11]
			veh.WorkshopOut = row[12]
			veh.MR1 = row[13]
			veh.MR2 = row[14]
			veh.FDFiring = row[15]
			veh.SeriesInspection = row[16]
			veh.Trg = row[17]
			veh.Remarks = row[18]
		} else {
			if len(rows[0]) != 10 {
				return nil, fmt.Errorf("missing B vehicle columns")
			}

			veh.CMSIn = row[5]
			veh.CMSOut = row[6]
			veh.WorkshopIn = row[7]
			veh.WorkshopOut = row[8]
			veh.Remarks = row[9]
		}

		vehicles = append(vehicles, veh)
	}

	return vehicles, nil
}

func stringToInteger(value string) int64 {
	num, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return -1
	}

	return num
}

func stringToFloat(value string) float64 {
	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return -1
	}

	return num
}
