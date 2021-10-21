package builder

import (
	"fmt"

	"rest-csv/models"
)

type AlertBuilder interface {
	CreateAlert(alert models.Alert) string
	UpdateAlert(alert models.Notification) string
	GetAlerts() string
	GetNotifications() string
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

func (a *alertBuilder) GetAlerts() string {
	return fmt.Sprintf("SELECT * FROM alerts")
}

func (a *alertBuilder) GetNotifications() string {
	return fmt.Sprintf(`SELECT a.id, v.id, a.name, v.ba_number, v.veh_type,
			a.alert_field, a.last_value, a.next_value, v.remarks, a.remarks FROM vehicles v INNER JOIN alerts a
			ON a.ba_number = v.ba_number AND (a.next_value = v.kilometers OR a.next_value = v.efc)`)
}
