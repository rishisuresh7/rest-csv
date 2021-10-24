package demand

import (
	"fmt"
	"strconv"

	"rest-csv/builder"
	"rest-csv/models"
	"rest-csv/repository"
)

type Demand interface {
	GetDemands(filters map[string]string) ([]models.Demand, error)
	AddDemands(demands []models.Demand) (int64, error)
	UpdateDemands(demands []models.Demand) (int64, error)
	DeleteDemands(ids []int64) (int64, error)
}

type demand struct {
	demandBuilder builder.Demand
	queryExecutor repository.QueryExecutor
}

func NewDemand(b builder.Demand, qe repository.QueryExecutor) Demand {
	return &demand{
		demandBuilder: b,
		queryExecutor: qe,
	}
}

func stringToInteger(value string) int64 {
	num, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return -1
	}

	return num
}

func (d *demand) GetDemands(filters map[string]string) ([]models.Demand, error) {
	query := d.demandBuilder.GetDemands(filters)
	sqlRows, err := d.queryExecutor.Query(query)
	if err != nil {
		return nil, fmt.Errorf("ListDemands: unable to query demands: %s", err)
	}

	rows, err := d.queryExecutor.ParseRows(sqlRows)
	if err != nil {
		return nil, fmt.Errorf("ListDemands: unable to parse rows: %s", err)
	}

	res := make([]models.Demand, 0)
	for _, row := range rows {
		demand := models.Demand{
			Item:              models.Item{
				Id:          stringToInteger(row[0]),
				Sqn:         row[1],
				VehicleType: row[2],
				BaNo:        row[3],
				Type:        row[4],
			},
			EquipmentDemanded: row[5],
			Depot:             row[6],
			DemandNumber:      row[7],
			DemandDate:        row[8],
			ControlNumber:     row[9],
			ControlDate:       row[10],
			Status:            row[11],
		}
		res = append(res, demand)
	}

	return res, nil
}

func (d *demand) AddDemands(demands []models.Demand) (int64, error) {
	query := d.demandBuilder.AddDemands(demands)
	rows, err := d.queryExecutor.Exec(query)
	if err != nil {
		return -1, fmt.Errorf("AddDemands: unable to insert demands: %s", err)
	}

	noOfRows, err := rows.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("AddDemands: unable to parse result: %s", err)
	}

	return noOfRows, nil
}

func (d *demand) UpdateDemands(demands []models.Demand) (int64, error) {
	query := d.demandBuilder.UpdateDemands(demands)
	rows, err := d.queryExecutor.Exec(query)
	if err != nil {
		return -1, fmt.Errorf("UpdateDemands: unable to upadate demands: %s", err)
	}

	noOfRows, err := rows.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("UpdateDemands: unable to parse result: %s", err)
	}

	return noOfRows, nil
}

func (d *demand) DeleteDemands(ids []int64) (int64, error) {
	query := d.demandBuilder.DeleteDemands(ids)
	rows, err := d.queryExecutor.Exec(query)
	if err != nil {
		return -1, fmt.Errorf("DeleteDemands: unable to delete demands: %s", err)
	}

	noOfRows, err := rows.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("DeleteDemands: unable to parse result: %s", err)
	}

	return noOfRows, nil
}