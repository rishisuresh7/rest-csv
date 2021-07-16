package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	ft "github.com/x-cray/logrus-prefixed-formatter"

	"rest-csv/middleware"
	"rest-csv/router"
)

var (
	Version = "0.0.0"
)

func main() {
	port := 9000
	level := 5
	token := "foobarbaz"
	publicRoutes := []string{"/health"}
	logger := logrus.New()
	logger.Level = logrus.Level(level)
	logger.Formatter = &ft.TextFormatter{
		ForceFormatting:  true,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
	}

	logger.Infof("Starting server version: %s", Version)
	r := router.NewRouter()

	n := negroni.New()
	n.Use(middleware.NewRequestLogger(logger))
	n.Use(middleware.NewAuthenticationMiddleware(logger, token, publicRoutes))
	n.UseHandler(r)
	n.Run(fmt.Sprintf("127.0.0.1:%d", port))
}
