package demand

import (
	"fmt"

	"rest-csv/builder"
	"rest-csv/repository"
)

type Demand interface {
	ListDemands() ([][]string, error)
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

func (d *demand) ListDemands() ([][]string, error) {
	query := d.demandBuilder.ListDemands()
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