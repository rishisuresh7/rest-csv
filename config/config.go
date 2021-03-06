package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type Config struct {
	Token    string
	Port     int
	LogLevel uint32
	LogFile  *os.File
	Username string
	Password string
	Secret   string
	SQLdb    string
}

func GenerateConfig() (*Config, error) {
	c := &Config{
		Port:     9000,
		Token:    "foobarbaz",
		LogFile:  os.Stdout,
		LogLevel: 5,
		Secret:   "super-secure-secret",
		SQLdb:    "./data.db",
	}

	token := os.Getenv("TOKEN")
	if token != "" {
		c.Token = token
	}

	portString := os.Getenv("PORT")
	if portString != "" {
		port, err := strconv.Atoi(portString)
		if err != nil {
			return nil, fmt.Errorf("GenerateConfig: invalid port: %s", err)
		}

		c.Port = port
	}

	fileLocation := os.Getenv("LOG_LOCATION")
	if fileLocation != "" {
		file, err := os.OpenFile(filepath.Join(fileLocation, "server.log"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
		if err != nil {
			return nil, fmt.Errorf("GenerateConfig: unable to open log file: %s", err)
		}

		c.LogFile = file
	}

	userName := os.Getenv("USERNAME")
	if userName == "" {
		return nil, fmt.Errorf("GenerateConfig: missing username")
	}

	password := os.Getenv("PASSWORD")
	if password == "" {
		return nil, fmt.Errorf("GenerateConfig: missing password")
	}

	secret := os.Getenv("SECRET")
	if secret != "" {
		c.Secret = secret
	}

	sqlDb := os.Getenv("SQL_DB")
	if sqlDb != "" {
		c.SQLdb = sqlDb
	}

	c.Username = userName
	c.Password = password

	return c, nil
}
