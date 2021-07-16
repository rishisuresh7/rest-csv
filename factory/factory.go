package factory

import "rest-csv/config"

type Factory interface {
}

type factory struct {
	config *config.Config
}

func NewFactory(c *config.Config) Factory {
	return &factory{config: c}
}
