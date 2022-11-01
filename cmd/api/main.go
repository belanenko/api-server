package main

import (
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/belanenko/api-server/config"
	"github.com/belanenko/api-server/internal/logger"
	"github.com/belanenko/api-server/internal/server"
	httpinfo "github.com/belanenko/api-server/internal/server/info/delivery/http"
)

func main() {
	logger := logger.NewLogger()

	if err := config.Init(); err != nil {
		logger.Fatal(err)
	}

	if viper.GetString("app.environment") == "production" {
		logger.SetProductionFormatter()
	}
	if err := logger.SetLogLevel(viper.GetString("log_level")); err != nil {
		logger.Fatal(err)
	}

	app := server.NewApp(
		logger,
		httpinfo.NewHandler(),
	)

	go func() {
		if err := app.Run(viper.GetString("app.port")); err != nil {
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
