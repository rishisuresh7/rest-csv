package vehicle

import (
	"fmt"

	"rest-csv/builder"
	"rest-csv/models"
	"rest-csv/repository"
)

type Vehicle interface {
	GetVehicles(map[string]string) ([][]string, error)
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

func (c *vehicle) GetVehicles(filters map[string]string) ([][]string, error) {
	query := c.vehicleBuilder.GetVehicles(filters)
	rows, err := c.queryExecutor.Query(query)
	if err != nil {
		return nil, fmt.Errorf("GetCategoryItems: unable to get details: %s", err)
	}

	res, err := c.queryExecutor.ParseRows(rows)
	if err != nil {
		return nil, fmt.Errorf("GetCategoryItems: unable to parse details: %s", err)
	}

	return res, nil
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
