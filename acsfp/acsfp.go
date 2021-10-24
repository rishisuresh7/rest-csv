package acsfp

import (
	"fmt"
	"strconv"

	"rest-csv/builder"
	"rest-csv/models"
	"rest-csv/repository"
)

type ACSFP interface {
	GetItems(filters map[string]string) ([]models.ACSFP, error)
	AddItems(items []models.ACSFP) (int64, error)
	UpdateItems(items []models.ACSFP) (int64, error)
	DeleteItems(ids []int64) (int64, error)
}

type acsfp struct {
	acsfpBuilder builder.ACSFPBuilder
	exec repository.QueryExecutor
}

func NewACSFP(b builder.ACSFPBuilder, e repository.QueryExecutor) ACSFP {
	return &acsfp{
		acsfpBuilder: b,
		exec: e,
	}
}

func stringToFloat(value string) float64 {
	num, err := strconv.ParseFloat(value, 10)
	if err != nil {
		return -1
	}

	return num
}

func stringToInteger(value string) int64 {
	num, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return -1
	}

	return num
}

func (a *acsfp) GetItems(filters map[string]string) ([]models.ACSFP, error) {
	query := a.acsfpBuilder.GetItems(filters)
	sqlRows, err := a.exec.Query(query)
	if err != nil {
		return nil, fmt.Errorf("GetItems: unable to get details: %s", err)
	}

	rows, err := a.exec.ParseRows(sqlRows)
	if err != nil {
		return nil, fmt.Errorf("GetItems: unable to parse details: %s", err)
	}

	res := make([]models.ACSFP, 0)
	for _, row := range rows {
		item := models.ACSFP{
			Id:               stringToInteger(row[0]),
			Name:             row[1],
			QuantityAuth:     stringToInteger(row[2]),
			QuantityHeld:     stringToInteger(row[3]),
			RegisteredNumber: row[4],
			YearOfProc:       stringToInteger(row[5]),
			Cost:             stringToFloat(row[6]),
			QuantityServed:   stringToInteger(row[7]),
			ForwardTo:        row[8],
			DemandDetails:    row[9],
			Remarks:          row[10],
		}
		res = append(res , item)
	}

	return res, nil
}

func (a acsfp) AddItems(items []models.ACSFP) (int64, error) {
	query := a.acsfpBuilder.AddItem(items)
	res, err := a.exec.Exec(query)
	if err != nil {
		return -1, fmt.Errorf("AddItems: unable to write data: %s", err)
	}

	noOfRows, err := res.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("AddItems: unable to parse result: %s", err)
	}

	return noOfRows, nil
}

func (a acsfp) UpdateItems(items []models.ACSFP) (int64, error) {
	query := a.acsfpBuilder.UpdateItem(items)
	rows, err := a.exec.Exec(query)
	if err != nil {
		return -1, fmt.Errorf("UpdateItems: unable to update data: %s", err)
	}

	noOfRows, err := rows.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("UpdateItems: unable to parse result: %s", err)
	}

	return noOfRows, nil
}

func (a acsfp) DeleteItems(ids []int64) (int64, error) {
	query := a.acsfpBuilder.DeleteItem(ids)
	res, err := a.exec.Exec(query)
	if err != nil {
		return -1, fmt.Errorf("DeleteItems: unable to delete: %s", err)
	}

	noOfRows, err := res.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("DeleteItems: unable parse delete result: %s", err)
	}

	return noOfRows, nil
}
