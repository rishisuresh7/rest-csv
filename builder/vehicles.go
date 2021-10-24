package builder

import (
	"fmt"

	"rest-csv/constant"
	"rest-csv/models"
)

type Vehicles interface {
	GetVehicles(filters map[string]string) string
	AddVehicles(items []models.Vehicle) string
	UpdateVehicles(items []models.Vehicle) string
	DeleteVehicles(ids []int64) string
}

type aVehicles struct{}

func NewVehicles(vehicleType string) Vehicles {
	if vehicleType == constant.AVehicle {
		return &aVehicles{}
	}

	return &bVehicles{}
}

func (c *aVehicles) GetVehicles(filters map[string]string) string {
	queryFilters := ""
	for key, value := range filters {
		if key != "search" {
			queryFilters += fmt.Sprintf(" AND %s = '%s'", key, value)
		} else {
			queryFilters += " AND (ba_number LIKE '%" + value + "%')"
		}
	}

	return `SELECT * FROM a_vehicles WHERE 1=1 ` + queryFilters
}

func (c *aVehicles) AddVehicles(items []models.Vehicle) string {
	item := items[0]
	return fmt.Sprintf(`INSERT INTO a_vehicles(id, squadron, veh_type, ba_number, type, kilometers, engine_hours,
			efc, tm_1, tm_2, cms_in, cms_out, workshop_in, workshop_out, mr_1, mr_2, fd_firing,
			series_inspection, trg_op, remarks)
			VALUES(NULL, '%s', '%s', '%s', '%s', '%d', %d, %d, '%s', '%s', '%s', '%s', '%s', '%s', '%s',
			'%s', '%s', '%s', '%s', '%s')`,
		item.Sqn, item.VehicleType, item.BaNo, item.Type, item.Kilometers, item.EngineHours, item.Efc, item.TM1, item.TM2,
		item.CMSIn, item.CMSOut, item.WorkshopIn, item.WorkshopOut, item.MR1, item.MR2, item.FDFiring,
		item.SeriesInspection, item.Trg, item.Remarks)
}

func (c *aVehicles) UpdateVehicles(items []models.Vehicle) string {
	item := items[0]
	return fmt.Sprintf(`UPDATE a_vehicles
			SET squadron = '%s', veh_type = '%s', ba_number = '%s', type = '%s',
				kilometers = %d, engine_hours = %d, efc = %d, tm_1 = '%s', tm_2 = '%s',
				cms_in = '%s', cms_out = '%s', series_inspection = '%s', trg_op = '%s', remarks = '%s',
				workshop_in = '%s', workshop_out = '%s', mr_1 = '%s', mr_2 = '%s', fd_firing = '%s'
			WHERE id = %d;
			`, item.Sqn, item.VehicleType, item.BaNo, item.Type, item.Kilometers, item.EngineHours, item.Efc, item.TM1,
			item.TM2, item.CMSIn, item.CMSOut, item.SeriesInspection, item.Trg, item.Remarks, item.WorkshopIn,
			item.WorkshopOut, item.MR1, item.MR2, item.FDFiring, item.Id)
}

func (c *aVehicles) DeleteVehicles(ids []int64) string {
	queryString := "( "
	for i := range ids {
		if i != len(ids)-1 {
			queryString += fmt.Sprintf("%d, ", ids[i])
		}
	}

	queryString = queryString + fmt.Sprintf("%d )", ids[len(ids)-1])

	return fmt.Sprintf("DELETE FROM a_vehicles WHERE id IN %s", queryString)
}

type bVehicles struct{}

func (c *bVehicles) GetVehicles(filters map[string]string) string {
	queryFilters := ""
	for key, value := range filters {
		if key != "search" {
			queryFilters += fmt.Sprintf(" AND %s = '%s'", key, value)
		} else {
			queryFilters += " AND (ba_number LIKE '%" + value + "%')"
		}
	}

	return `SELECT * FROM b_vehicles WHERE 1=1 ` + queryFilters
}

func (c *bVehicles) AddVehicles(items []models.Vehicle) string {
	item := items[0]
	return fmt.Sprintf(`INSERT INTO b_vehicles(id, squadron, veh_type, ba_number, type, kilometers,
			cms_in, cms_out, workshop_in, workshop_out, remarks)
			VALUES(NULL, '%s', '%s', '%s', '%s', '%d', '%s', '%s', '%s', '%s', '%s')`,
		item.Sqn, item.VehicleType, item.BaNo, item.Type, item.Kilometers, item.CMSIn, item.CMSOut,
		item.WorkshopIn, item.WorkshopOut, item.Remarks)
}

func (c *bVehicles) UpdateVehicles(items []models.Vehicle) string {
	item := items[0]
	return fmt.Sprintf(`UPDATE b_vehicles
			SET squadron = '%s', veh_type = '%s', ba_number = '%s', type = '%s', kilometers = %d,
			cms_in = '%s', cms_out = '%s', remarks = '%s', workshop_in = '%s', workshop_out = '%s'
			WHERE id = %d;
			`, item.Sqn, item.VehicleType, item.BaNo, item.Type, item.Kilometers, item.CMSIn,
			item.CMSOut, item.Remarks, item.WorkshopIn, item.WorkshopOut, item.Id)
}

func (c *bVehicles) DeleteVehicles(ids []int64) string {
	queryString := "( "
	for i := range ids {
		if i != len(ids)-1 {
			queryString += fmt.Sprintf("%d, ", ids[i])
		}
	}

	queryString = queryString + fmt.Sprintf("%d )", ids[len(ids)-1])

	return fmt.Sprintf("DELETE FROM b_vehicles WHERE id IN %s", queryString)
}