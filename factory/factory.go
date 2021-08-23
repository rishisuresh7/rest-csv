package factory

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"

	"rest-csv/auth"
	"rest-csv/builder"
	"rest-csv/category"
	"rest-csv/config"
	"rest-csv/demand"
	"rest-csv/middleware"
	"rest-csv/repository"
	"rest-csv/utility"
)

var fileWriter sync.Once
var sqlDriver sync.Once

type Factory interface {
	ReadWriter(file string) (*os.File, error)
	Category(name string) category.Category
	Demand() demand.Demand
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
		header: []string{},
	}

	return f
}

func (f *factory) initializeFiles(fileName string) (*os.File, error) {
	fileNameType := fmt.Sprintf("%s.csv", fileName)
	file, err := os.OpenFile(filepath.Join(f.config.DataLocation, fileNameType), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		return nil, err
	}

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if stat.Size() == 0 {
		csvWriter := csv.NewWriter(file)
		err = csvWriter.Write(f.header)
		csvWriter.Flush()
	}

	return file, nil
}

func (f *factory) initialize() (map[string]*os.File, error) {
	var scErr error
	fileWriter.Do(func() {
		f.files = make(map[string]*os.File)
		for _, name := range f.config.Categories {
			fileName := utility.SanitizeFileName(name)
			file, err := f.initializeFiles(fileName)
			if err != nil {
				scErr = err
				f.files = nil
				return
			}

			f.files[fileName] = file
		}
	})

	return f.files, scErr
}

func (f *factory) connect() *sql.DB {
	sqlDriver.Do(func() {
		conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/ems", f.config.SQLUsername, f.config.SQLPassword))
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

func (f *factory) ReadWriter(file string) (*os.File, error) {
	files, err := f.initialize()
	if err != nil {
		return nil, fmt.Errorf("ReadWriter: unable to get file(category) pointer: %s", err)
	}

	fileName := utility.SanitizeFileName(file)
	fl, ok := files[fileName]
	if !ok {
		return nil, fmt.Errorf("ReadWriter: invalid category: %s", err)
	}

	return fl, nil
}

func (f *factory) Category(name string) category.Category {
	file, _ := f.ReadWriter(name)
	return category.NewCategory(file, f.config.Categories, builder.NewCategories(), f.QueryExecutor())
}

func (f *factory) Demand() demand.Demand {
	return demand.NewDemand(builder.NewDemand(), f.QueryExecutor())
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
