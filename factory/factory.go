package factory

import (
	"database/sql"
	"os"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"

	"rest-csv/alerts"
	"rest-csv/auth"
	"rest-csv/builder"
	"rest-csv/category"
	"rest-csv/config"
	"rest-csv/demand"
	"rest-csv/middleware"
	"rest-csv/repository"
)

var sqlDriver sync.Once

type Factory interface {
	Category() category.Category
	Demand() demand.Demand
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

func (f *factory) Category() category.Category {
	return category.NewCategory(builder.NewCategories(), f.QueryExecutor())
}

func (f *factory) Demand() demand.Demand {
	return demand.NewDemand(builder.NewDemand(), f.QueryExecutor())
}

func (f *factory) Alerts() alerts.Alerts {
	return alerts.NewAlerts(builder.NewAlertBuilder(), f.QueryExecutor())
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
