package builder

import (
	"fmt"
	"strings"

	"rest-csv/models"
)

type ACSFPBuilder interface {
	GetItems(filters map[string]string) string
	UpdateItem(items []models.ACSFP) string
	AddItem(items []models.ACSFP) string
	DeleteItem(ids []int64) string
}

type acsfp struct{}

func NewACSFPBuilder() ACSFPBuilder {
	return &acsfp{}
}

func (a *acsfp) GetItems(filters map[string]string) string {
	return "SELECT * FROM acsfp"
}

func (a *acsfp) UpdateItem(items []models.ACSFP) string {
	item := items[0]
	return fmt.Sprintf(`UPDATE acsfp SET name = '%s', qty_auth = %d, qty_held = %d, regd_number = '%s',
			year_of_proc = %d, cost = %f, qty_ser = %d, fwd_to = '%s', demand_details = '%s', remarks = '%s'
			WHERE id = %d`,
		item.Name, item.QuantityAuth, item.QuantityHeld, item.RegisteredNumber, item.YearOfProc, item.Cost,
		item.QuantityServed, item.ForwardTo, item.DemandDetails, item.Remarks, item.Id)
}

func (a *acsfp) AddItem(items []models.ACSFP) string {
	var values []string
	for _, item := range items {
		value := fmt.Sprintf(`(NULL, '%s', %d, %d, '%s', %d, %f, %d, '%s', '%s', '%s')`,
			item.Name, item.QuantityAuth, item.QuantityHeld, item.RegisteredNumber, item.YearOfProc, item.Cost,
			item.QuantityServed, item.ForwardTo, item.DemandDetails, item.Remarks)

		values = append(values, value)
	}

	return fmt.Sprintf(`INSERT INTO acsfp(id, name, qty_auth, qty_held, regd_number, year_of_proc,
			cost, qty_ser, fwd_to, demand_details, remarks)
			VALUES %s`, strings.Join(values, ", "))
}

func (a *acsfp) DeleteItem(ids []int64) string {
	queryString := "( "
	for i := range ids {
		if i != len(ids)-1 {
			queryString += fmt.Sprintf("%d, ", ids[i])
		}
	}

	queryString = queryString + fmt.Sprintf("%d )", ids[len(ids)-1])

	return fmt.Sprintf("DELETE FROM acsfp WHERE id IN %s", queryString)
}
