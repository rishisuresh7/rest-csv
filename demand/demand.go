package demand

import (
	"fmt"

	"rest-csv/builder"
	"rest-csv/models"
	"rest-csv/repository"
)

type Demand interface {
	GetDemands() ([][]string, error)
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

func (d *demand) GetDemands() ([][]string, error) {
	query := d.demandBuilder.GetDemands()
	rows, err := d.queryExecutor.Query(query)
	if err != nil {
		return nil, fmt.Errorf("ListDemands: unable to query demands: %s", err)
	}

	res, err := d.queryExecutor.ParseRows(rows)
	if err != nil {
		return nil, fmt.Errorf("ListDemands: unable to parse rows: %s", err)
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