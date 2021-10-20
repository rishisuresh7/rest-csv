package alerts

import (
	"fmt"
	"strconv"

	"rest-csv/builder"
	"rest-csv/models"
	"rest-csv/repository"
)

type Alerts interface {
	CreateAlert(alert models.Alert) error
	UpdateAlert(alert models.Alert) error
	GetAlerts() ([]models.Alert, error)
	GetNotifications() ([]models.Notification, error)
}

type alert struct {
	alertBuilder builder.AlertBuilder
	exec repository.QueryExecutor
}

func NewAlerts(b builder.AlertBuilder, e repository.QueryExecutor) Alerts {
	return &alert{alertBuilder: b, exec: e}
}

func (a *alert) CreateAlert(alert models.Alert) error {
	query := a.alertBuilder.CreateAlert(alert)
	_, err := a.exec.Exec(query)
	if err != nil {
		return fmt.Errorf("CreateAlert: unable to create alert: %s", err)
	}

	return nil
}

func (a *alert) UpdateAlert(alert models.Alert) error {
	query := a.alertBuilder.UpdateAlert(alert)
	_, err := a.exec.Exec(query)
	if err != nil {
		return fmt.Errorf("UpdateAlert: unable to update alert: %s", err)
	}

	return nil
}

func (a *alert) GetAlerts() ([]models.Alert, error) {
	query := a.alertBuilder.GetAlerts()
	sqlRows, err := a.exec.Query(query)
	if err != nil {
		return nil, fmt.Errorf("GetAlerts: unable to query alerts: %s", err)
	}

	rows, err := a.exec.ParseRows(sqlRows)
	if err != nil {
		return nil, fmt.Errorf("GetAlerts: unable to parse rows: %s", err)
	}

	var res []models.Alert
	for _, row := range rows {
		alert := models.Alert{
			Id:         stringToInteger(row[0]),
			Name:       row[1],
			BaNo:       row[2],
			AlertField: row[3],
			LastValue:  row[4],
			NextValue:  row[5],
		}
		res = append(res, alert)
	}

	return res, nil
}

func (a *alert) GetNotifications() ([]models.Notification, error) {
	query := a.alertBuilder.GetNotifications()
	sqlRows, err := a.exec.Query(query)
	if err != nil {
		return nil, fmt.Errorf("GetNotifications: unable to query db: %s", err)
	}

	rows, err := a.exec.ParseRows(sqlRows)
	if err != nil {
		return nil, fmt.Errorf("GetNotifications: unable to parse rows: %s", err)
	}

	res := make([]models.Notification, 0)
	for _, row := range rows {
		notification := models.Notification{
			AlertId:     stringToInteger(row[0]),
			VehicleId:   stringToInteger(row[1]),
			AlertName:   row[2],
			BaNo:        row[3],
			VehicleType: row[4],
			AlertField:  row[5],
			LastValue:   row[6],
			NextValue:   row[7],
		}
		res = append(res, notification)
	}

	return res, nil
}

func stringToInteger(value string) int64 {
	num, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return -1
	}

	return num
}