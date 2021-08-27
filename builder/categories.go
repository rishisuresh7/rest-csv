package builder

import (
	"fmt"

	"rest-csv/models"
)

type Categories interface {
	GetVehicles(filters map[string]string) string
	AddVehicles(items []models.Vehicle) string
	UpdateVehicles(items []models.Vehicle) string
	DeleteVehicles(ids []int64) string
}

type categories struct{}

func NewCategories() Categories {
	return &categories{}
}

func (c *categories) GetVehicles(filters map[string]string) string {
	queryFilters := ""
	for key, value := range filters {
		if key != "search" {
			queryFilters += fmt.Sprintf(" AND %s = '%s'", key, value)
		} else {
			queryFilters += " AND (ba_number LIKE '%" + value +  "%')"
		}
	}

	return `SELECT * FROM vehicles WHERE 1=1 ` + queryFilters
}

func (c *categories) AddVehicles(items []models.Vehicle) string {
	item := items[0]
	return fmt.Sprintf(`INSERT INTO vehicles(id, squadron, veh_type, ba_number, type, kilometers,
			engine_hours, efc, tm_1, tm_2, cms_in, cms_out, series_inspection, tag_op)
			VALUES(NULL, '%s', '%s', '%s', '%s', %d, %d, %d, %s, %s, %s, %s, '%s', '%s')`,
		item.Sqn, item.VehicleType, item.BaNo, item.Type, item.Kilometers, item.EngineHours, item.Efc,
		toSQLDate(item.TM1), toSQLDate(item.TM2), toSQLDate(item.CMSIn), toSQLDate(item.CMSOut),
		item.SeriesInspection, item.Tag)
}

func toSQLDate(val string) string {
	return `STR_TO_DATE('` + val + `', '%Y-%m-%d')`
}

func (c *categories) UpdateVehicles(items []models.Vehicle) string {
	item := items[0]
	return fmt.Sprintf(`UPDATE vehicles
			SET squadron = '%s', veh_type = '%s', ba_number = '%s', type = '%s',
				kilometers = %d, engine_hours = %d, efc = %d, tm_1 = %s, tm_2 = %s,
				cms_in = %s, cms_out = %s, series_inspection = '%s', tag_op = '%s'
			WHERE id = %d;
			`, item.Sqn, item.VehicleType, item.BaNo, item.Type, item.Kilometers, item.EngineHours, item.Efc,
		toSQLDate(item.TM1), toSQLDate(item.TM2), toSQLDate(item.CMSIn), toSQLDate(item.CMSOut),
		item.SeriesInspection, item.Tag, item.Id)
}

func (c *categories) DeleteVehicles(ids []int64) string {
	queryString := "( "
	for i := range ids {
		if i != len(ids)-1 {
			queryString += fmt.Sprintf("%d, ", ids[i])
		}
	}

	queryString = queryString + fmt.Sprintf("%d )", ids[len(ids)-1])

	return fmt.Sprintf("DELETE FROM vehicles WHERE id IN %s", queryString)
}
