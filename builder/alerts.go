package builder

import (
	"fmt"
	"strings"

	"rest-csv/models"
)

type AlertBuilder interface {
	CreateAlert(alert models.Alert) string
	UpdateAlert(alert models.Notification) string
	ModifyAlert(alert models.Alert) string
	DeleteAlerts(ids []int64) string
	GetAlerts() string
	GetNotifications(alerts []models.Alert) string
}

type alertBuilder struct{}

func NewAlertBuilder() AlertBuilder {
	return &alertBuilder{}
}

func (a *alertBuilder) CreateAlert(alert models.Alert) string {
	return fmt.Sprintf(`INSERT INTO alerts(name, ba_number, alert_field, last_value, next_value, remarks)
			VALUES('%s', '%s', '%s', '%s', '%s', '%s')`, alert.Name, alert.BaNo, alert.AlertField, alert.LastValue, alert.NextValue, alert.Remarks)
}

func (a *alertBuilder) UpdateAlert(alert models.Notification) string {
	return fmt.Sprintf(`UPDATE alerts SET last_value = '%s', next_value = '%s',
						remarks = '%s' WHERE id = %d`, alert.LastValue, alert.NextValue, alert.AlertRemarks, alert.AlertId)
}

func (a *alertBuilder) ModifyAlert(alert models.Alert) string {
	return fmt.Sprintf(`UPDATE alerts SET name = '%s', ba_number = '%s', alert_field = '%s', last_value = '%s',
			next_value = '%s', remarks = '%s' WHERE id = %d`, alert.Name, alert.BaNo, alert.AlertField, alert.LastValue,
			alert.NextValue, alert.Remarks, alert.Id)
}

func (a *alertBuilder) GetAlerts() string {
	return fmt.Sprintf("SELECT * FROM alerts")
}

type fieldMap struct {
	operator string
	fieldName string
	filter string
}

var alertMap = map[string]fieldMap{
	"kilometers": {
		operator: "<=",
		fieldName: "kilometers",
	},
	"efc": {
		operator: "<=",
		fieldName: "efc",
	},
	"tm 1": {
		operator: "=",
		fieldName: "tm_1",
		filter: "AND julianday(a.next_value) <= julianday('now')",
	},
	"tm 2": {
		operator: "=",
		fieldName: "tm_2",
		filter: "AND julianday(a.next_value) <= julianday('now')",
	},
	"cms in": {
		operator: "=",
		fieldName: "cms_in",
		filter: "AND julianday(a.next_value) <= julianday('now')",
	},
	"cms out": {
		operator: "=",
		fieldName: "cms_out",
		filter: "AND julianday(a.next_value) <= julianday('now')",
	},
}

func (a *alertBuilder) isExcluded(value string) bool {
	if value == "tm_1" || value == "tm_2" || value == "efc" {
		return true
	}

	return false
}

func (a *alertBuilder) GetNotifications(alerts []models.Alert) string {
	var values []string
	for _, alert := range alerts {
		field := alertMap[strings.ToLower(alert.AlertField)]
		alertQuery := fmt.Sprintf(`SELECT DISTINCT a.id, v.id, a.name, v.ba_number, v.veh_type, a.alert_field,
			a.last_value, a.next_value, v.remarks, a.remarks
			FROM a_vehicles v INNER JOIN alerts a ON a.ba_number = v.ba_number
			WHERE a.next_value %s v.%s %s`, field.operator, field.fieldName, field.filter)
		if !a.isExcluded(field.fieldName) {
			alertQuery += fmt.Sprintf(` UNION SELECT DISTINCT a.id, v.id, a.name, v.ba_number, v.veh_type, a.alert_field,
			a.last_value, a.next_value, v.remarks, a.remarks
			FROM b_vehicles v INNER JOIN alerts a ON a.ba_number = v.ba_number
			WHERE a.next_value %s v.%s %s`, field.operator, field.fieldName, field.filter)
		}
		values = append(values, alertQuery)
	}

	return strings.Join(values, " UNION ")
}

func (a *alertBuilder) DeleteAlerts(ids []int64) string {
	queryString := "( "
	for i := range ids {
		if i != len(ids)-1 {
			queryString += fmt.Sprintf("%d, ", ids[i])
		}
	}

	queryString = queryString + fmt.Sprintf("%d )", ids[len(ids)-1])

	return fmt.Sprintf("DELETE FROM alerts WHERE id IN %s", queryString)
}