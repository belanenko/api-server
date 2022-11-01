package main

import (
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"

	"github.com/belanenko/api-server/config"
	"github.com/belanenko/api-server/internal/logger"
	"github.com/belanenko/api-server/internal/server"
	httpinfo "github.com/belanenko/api-server/internal/server/info/delivery/http"
)

func main() {
	logger := logger.NewLogger()

	config, err := config.Init()
	if err != nil {
		logger.Fatal(err)
	}

	if config.Environment == "production" {
		logger.SetProductionFormatter()
	}
	if err := logger.SetLogLevel(config.LogLevel); err != nil {
		logger.Fatal(err)
	}

	logrus.WithFields(logrus.Fields{
		"environment": config.Environment,
	}).Info()

	app := server.NewApp(
		logger,
		httpinfo.NewHandler(),
	)

	go func() {
		if err := app.Run(config.App.Port); err != nil {
			logger.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	if err := app.Shutdown(); err != nil {
		logrus.WithFields(
			logrus.Fields{
				"error": err,
			},
		).Error("app shutdown")
	}
}
