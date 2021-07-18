package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	ft "github.com/x-cray/logrus-prefixed-formatter"

	"rest-csv/config"
	"rest-csv/factory"
	"rest-csv/middleware"
	"rest-csv/router"
)

var (
	Version = "0.0.0"
)

func main() {
	publicRoutes := []string{"/health"}
	logger := logrus.New()
	logger.Formatter = &ft.TextFormatter{
		ForceFormatting:  true,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
	}

	conf, err := config.GenerateConfig()
	if err != nil {
		logger.Fatalf("Unable to generate config: %s", err)
	}

	logger.Out = conf.LogFile
	logger.Infof("Starting server version: %s", Version)
	logger.Level = logrus.Level(conf.LogLevel)
	f := factory.NewFactory(conf, logger)
	r := router.NewRouter(f, logger)
	n := negroni.New()

	n.Use(middleware.NewRequestLogger(logger))
	n.Use(middleware.NewAuthenticationMiddleware(logger, conf.Token, publicRoutes))
	n.UseHandler(r)
	n.Run(fmt.Sprintf("127.0.0.1:%d", conf.Port))
}
