package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Config struct {
	Token        string
	Port         int
	LogLevel     uint32
	Categories   []string
	LogFile      *os.File
	DataLocation string
}

func GenerateConfig() (*Config, error) {
	c := &Config{
		Port:     9000,
		Token:    "foobarbaz",
		LogFile:  os.Stdout,
		LogLevel: 5,
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

	categoriesArray := os.Getenv("CATEGORIES")
	if categoriesArray != "" {
		categories := strings.Split(categoriesArray, ",")
		for _, value := range categories {
			c.Categories = append(c.Categories, strings.TrimSpace(value))
		}
	}

	dataLoc := os.Getenv("DATA_LOCATION")
	if dataLoc == "" {
		return nil, fmt.Errorf("GenerateConfig: missing data location")
	}

	if err := os.MkdirAll(dataLoc, 0755); err != nil {
		return nil, fmt.Errorf("GenerateConfig: unable to create directory: %s", err)
	}

	c.DataLocation = dataLoc

	return c, nil
}
