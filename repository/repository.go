package repository

import (
	"database/sql"
	"fmt"
)

type QueryExecutor interface {
	Exec(query string, args ...interface{}) (*sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	ParseRows(rows *sql.Rows) ([][]string, error)
}

type queryExecutor struct {
	db *sql.DB
}

func NewExecutor(d *sql.DB) QueryExecutor {
	return &queryExecutor{
		db: d,
	}
}

func (qe *queryExecutor) Exec(query string, args ...interface{}) (*sql.Result, error) {
	return nil, nil
}

func (qe *queryExecutor) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := qe.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("Query: unable to query database: %s", err)
	}

	return rows, nil
}

func (qe *queryExecutor) ParseRows(rows *sql.Rows) ([][]string, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("ParseRows: unable to get columns: %s", err)
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	res := make([][]string, 0)
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, fmt.Errorf("ParseRows: unable to parse columns: %s", err)
		}

		row := make([]string, len(columns))
		for j, value := range values {
			row[j] = toString(value)
		}

		res = append(res, row)
	}

	return res, nil
}

func toString(val interface{}) string {
	if val == nil {
		return ""
	} else {
		return fmt.Sprintf("%v", val)
	}
}
