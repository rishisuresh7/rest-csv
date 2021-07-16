package config

import (
	"os"
)

type Config struct {
	Token      string
	Port       int
	LogLevel   uint32
	Categories []string
	LogFile    *os.File
}

func GenerateConfig() (*Config, error) {
	c := &Config{
		Port: 9000,
		Token: "foobarbaz",
		LogFile: os.Stdout,
		LogLevel: 5,
	}

	return c, nil
}
