package builder

import (
	"fmt"
	"strings"

	"rest-csv/models"
)

type Demand interface {
	GetDemands(filters map[string]string) string
	AddDemands(items []models.Demand) string
	UpdateDemands(items []models.Demand) string
	DeleteDemands(ids []int64) string
}

type demand struct{}

func NewDemand() Demand {
	return &demand{}
}

func (d *demand) GetDemands(filters map[string]string) string {
	queryFilters := ""
	for key, value := range filters {
		queryFilters += fmt.Sprintf(" AND %s = '%s'", key, value)
	}

	return "SELECT * FROM demands WHERE 1=1 " + queryFilters
}

func (d *demand) AddDemands(items []models.Demand) string {
	var values []string
	for _, item := range items {
		value := fmt.Sprintf(`(NULL, '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')`,
			item.Sqn, item.VehicleType, item.BaNo, item.Type, item.EquipmentDemanded,
			item.Depot, item.DemandNumber, item.DemandDate, item.ControlNumber, item.ControlDate, item.Status)

		values = append(values, value)
	}

	return fmt.Sprintf(`INSERT INTO demands(id, squadron, veh_type, ba_number, type, equipment_demanded,
			depot, demand_number, demand_date, control_number, control_date, status)
			VALUES %s`, strings.Join(values, ", "))
}

func (d *demand) UpdateDemands(items []models.Demand) string {
	item := items[0]
	return fmt.Sprintf(`UPDATE demands
			SET ba_number = '%s', control_number = '%s', demand_number = '%s', depot = '%s',
				equipment_demanded = '%s', squadron = '%s', status = '%s', type = '%s', veh_type = '%s',
				demand_date = '%s', control_date = '%s'
			WHERE id = %d;
			`, item.BaNo, item.ControlNumber, item.DemandNumber, item.Depot, item.EquipmentDemanded,
		item.Sqn, item.Status, item.Type, item.VehicleType, item.DemandDate, item.ControlDate, item.Id)
}

func (d *demand) DeleteDemands(ids []int64) string {
	queryString := "( "
	for i := range ids {
		if i != len(ids)-1 {
			queryString += fmt.Sprintf("%d, ", ids[i])
		}
	}

	queryString = queryString + fmt.Sprintf("%d )", ids[len(ids)-1])

	return fmt.Sprintf("DELETE FROM demands WHERE id IN %s", queryString)
}
