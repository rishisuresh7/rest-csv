package factory

import (
	"database/sql"
	"os"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"

	"rest-csv/acsfp"
	"rest-csv/alerts"
	"rest-csv/auth"
	"rest-csv/builder"
	"rest-csv/config"
	"rest-csv/constant"
	"rest-csv/demand"
	"rest-csv/export"
	"rest-csv/importer"
	"rest-csv/middleware"
	"rest-csv/repository"
	"rest-csv/vehicle"
)

var sqlDriver sync.Once

type Factory interface {
	Vehicles(vehicleType string) vehicle.Vehicle
	Demand() demand.Demand
	ACSFP() acsfp.ACSFP
	Exporter() export.Exporter
	Importer(viewName string) importer.Importer
	Alerts() alerts.Alerts
	Auth() auth.Auth
	NewJWTAuth() *middleware.JWTAuthenticator
}

type factory struct {
	logger *logrus.Logger
	config *config.Config
	db     *sql.DB
	header []string
	files  map[string]*os.File
}

func NewFactory(c *config.Config, l *logrus.Logger) Factory {
	f := &factory{
		config: c,
		logger: l,
	}

	return f
}

func (f *factory) connect() *sql.DB {
	sqlDriver.Do(func() {
		conn, err := sql.Open("sqlite3", f.config.SQLdb)
		if err != nil {
			f.logger.Fatalf("Unable to connect to DB: %s....quitting....\n", err)
		}

		conn.SetConnMaxLifetime(time.Minute * 3)
		conn.SetMaxOpenConns(10)
		conn.SetMaxIdleConns(10)
		f.db = conn
	})

	return f.db
}

func (f *factory) Vehicles(vehicleType string) vehicle.Vehicle {
	return vehicle.NewVehicle(builder.NewVehicles(vehicleType), f.QueryExecutor())
}

func (f *factory) Demand() demand.Demand {
	return demand.NewDemand(builder.NewDemand(), f.QueryExecutor())
}

func (f *factory) ACSFP() acsfp.ACSFP {
	return acsfp.NewACSFP(builder.NewACSFPBuilder(), f.QueryExecutor())
}

func (f *factory) Alerts() alerts.Alerts {
	return alerts.NewAlerts(builder.NewAlertBuilder(), f.QueryExecutor())
}

func (f *factory) Exporter() export.Exporter {
	return export.NewExporter(builder.NewExportBuilder(), f.QueryExecutor())
}

func (f *factory) Importer(viewName string) importer.Importer {
	var inserter interface{}
	switch viewName {
	case "a_vehicles":
		inserter = f.Vehicles(constant.AVehicle)
	case "b_vehicles":
		inserter = f.Vehicles(constant.BVehicle)
	case "demands":
		inserter = f.Demand()
	case "acsfp":
		inserter = f.ACSFP()
	}

	return importer.NewImporter(inserter, f.QueryExecutor())
}

func (f *factory) QueryExecutor() repository.QueryExecutor {
	return repository.NewExecutor(f.connect())
}

func (f *factory) Auth() auth.Auth {
	return auth.NewAuth(f.config.Username, f.config.Password, f.config.Secret)
}

func (f *factory) NewJWTAuth() *middleware.JWTAuthenticator {
	return middleware.NewJWTAuthenticator(f.logger, f.config.Secret)
}
