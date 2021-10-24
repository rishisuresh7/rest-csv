package vehicle

import (
	"fmt"
	"strconv"

	"rest-csv/builder"
	"rest-csv/constant"
	"rest-csv/models"
	"rest-csv/repository"
)

type Vehicle interface {
	GetVehicles(map[string]string) ([]models.Vehicle, error)
	AddVehicles(item []models.Vehicle) (int64, error)
	UpdateVehicles(item []models.Vehicle) (int64, error)
	DeleteVehicles(id []int64) (int64, error)
}

type vehicle struct {
	vehicleBuilder builder.Vehicles
	queryExecutor  repository.QueryExecutor
}

func NewVehicle(cb builder.Vehicles, qe repository.QueryExecutor) Vehicle {
	return &vehicle{
		vehicleBuilder: cb,
		queryExecutor:   qe,
	}
}

func stringToInteger(value string) int64 {
	num, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return -1
	}

	return num
}

func (c *vehicle) GetVehicles(filters map[string]string) ([]models.Vehicle, error) {
	vehicleType := filters["vehicleType"]
	delete(filters, "vehicleType")
	query := c.vehicleBuilder.GetVehicles(filters)
	sqlRows, err := c.queryExecutor.Query(query)
	if err != nil {
		return nil, fmt.Errorf("GetVehicles: unable to get details: %s", err)
	}

	rows, err := c.queryExecutor.ParseRows(sqlRows)
	if err != nil {
		return nil, fmt.Errorf("GetVehicles: unable to parse details: %s", err)
	}

	res := make([]models.Vehicle, 0)
	for _, row := range rows {
		res = append(res , c.getVehicleRow(vehicleType, row))
	}

	return res, nil
}

func (c *vehicle) getVehicleRow(vehicleType string, row []string) models.Vehicle {
	item := models.Item{
		Id:          stringToInteger(row[0]),
		Sqn:         row[1],
		VehicleType: row[2],
		BaNo:        row[3],
		Type:        row[4],
	}

	if vehicleType == constant.AVehicle {
		item.Remarks = row[19]
		return models.Vehicle{
			Item:             item,
			Kilometers:       stringToInteger(row[5]),
			EngineHours:      stringToInteger(row[6]),
			Efc:              stringToInteger(row[7]),
			TM1:              row[8],
			TM2:              row[9],
			CMSIn:            row[10],
			CMSOut:           row[11],
			WorkshopIn:       row[12],
			WorkshopOut:      row[13],
			MR1:              row[14],
			MR2:              row[15],
			FDFiring:         row[16],
			SeriesInspection: row[17],
			Trg:              row[18],
		}
	} else {
		item.Remarks = row[10]
		return models.Vehicle{
			Item:        item,
			Kilometers:  stringToInteger(row[5]),
			CMSIn:       row[6],
			CMSOut:      row[7],
			WorkshopIn:  row[8],
			WorkshopOut: row[9],
		}
	}
}

func (c *vehicle) AddVehicles(items []models.Vehicle) (int64, error) {
	query := c.vehicleBuilder.AddVehicles(items)
	res, err := c.queryExecutor.Exec(query)
	if err != nil {
		return -1, fmt.Errorf("AddVehicles: unable to write data: %s", err)
	}

	noOfRows, err := res.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("AddVehicles: unable to parse result: %s", err)
	}

	return noOfRows, nil
}

func (c *vehicle) UpdateVehicles(items []models.Vehicle) (int64, error) {
	query := c.vehicleBuilder.UpdateVehicles(items)
	rows, err := c.queryExecutor.Exec(query)
	if err != nil {
		return -1, fmt.Errorf("UpdateVehicles: unable to update data: %s", err)
	}

	noOfRows, err := rows.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("UpdateVehicles: unable to parse result: %s", err)
	}

	return noOfRows, nil
}

func (c *vehicle) DeleteVehicles(ids []int64) (int64, error) {
	query := c.vehicleBuilder.DeleteVehicles(ids)
	res, err := c.queryExecutor.Exec(query)
	if err != nil {
		return -1, fmt.Errorf("DeleteVehicles: unable to delete: %s", err)
	}

	noOfRows, err := res.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("DeleteVehicles: unable parse delete result: %s", err)
	}

	return noOfRows, nil
}
