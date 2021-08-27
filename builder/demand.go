package builder

import (
	"fmt"

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
	item := items[0]
	return fmt.Sprintf(`INSERT INTO demands(id, ba_number, control_number, demand_number, depot,
			equipment_demanded, squadron, status, type, veh_type)
			VALUES(NULL, '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')`,
		item.BaNo, item.ControlNumber, item.DemandNumber, item.Depot, item.EquipmentDemanded,
		item.Sqn, item.Status, item.Type, item.VehicleType)
}

func (d *demand) UpdateDemands(items []models.Demand) string {
	item := items[0]
	return fmt.Sprintf(`UPDATE demands
			SET ba_number = '%s', control_number = '%s', demand_number = '%s', depot = '%s',
				equipment_demanded = '%s', squadron = '%s', status = '%s', type = '%s', veh_type = '%s'
			WHERE id = %d;
			`, item.BaNo, item.ControlNumber, item.DemandNumber, item.Depot, item.EquipmentDemanded,
			item.Sqn, item.Status, item.Type, item.VehicleType, item.Id)
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