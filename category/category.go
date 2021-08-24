package category

import (
	"fmt"

	"rest-csv/builder"
	"rest-csv/models"
	"rest-csv/repository"
)

type Category interface {
	GetVehicles() ([][]string, error)
	AddVehicles(item []models.Vehicle) (int64, error)
	UpdateVehicles(item []models.Vehicle) (int64, error)
	DeleteVehicles(id []int64) (int64, error)
}

type category struct {
	categoryBuilder builder.Categories
	queryExecutor   repository.QueryExecutor
}

func NewCategory(cb builder.Categories, qe repository.QueryExecutor) Category {
	return &category{
		categoryBuilder: cb,
		queryExecutor:   qe,
	}
}

func (c *category) GetVehicles() ([][]string, error) {
	query := c.categoryBuilder.GetVehicles()
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

func (c *category) AddVehicles(items []models.Vehicle) (int64, error) {
	query := c.categoryBuilder.AddVehicles(items)
	fmt.Println(query)
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

func (c *category) UpdateVehicles(items []models.Vehicle) (int64, error) {
	query := c.categoryBuilder.UpdateVehicles(items)
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

func (c *category) DeleteVehicles(ids []int64) (int64, error) {
	query := c.categoryBuilder.DeleteVehicles(ids)
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
