package server

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/belanenko/api-server/internal/logger"
	"github.com/belanenko/api-server/internal/server/info"
	httpinfo "github.com/belanenko/api-server/internal/server/info/delivery/http"
)

type App struct {
	httpServer *http.Server
	logger     *logger.Logger

	info info.UseCase
}

func NewApp(logger *logger.Logger, useCase info.UseCase) *App {
	return &App{
		logger: logger,
		info:   useCase,
	}
}

func (a *App) Run(port string) error {
	if err := a.validatePort(port); err != nil {
		return fmt.Errorf("port is invalid: %w", err)
	}

	router := gin.Default()
	api := router.Group("/api")

	httpinfo.RegisterHTTPEndpoints(api, a.info)

	a.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go a.listen(port)

	return nil
}

func (a *App) listen(port string) error {
	a.logger.WithFields(logrus.Fields{
		"port": port,
	}).Info("start listen connections")
	if err := a.httpServer.ListenAndServe(); err != nil {
		a.logger.WithFields(logrus.Fields{
			"port":  port,
			"error": err,
		}).Warn()
	}
	return nil
}

func (a *App) Shutdown() error {

	a.logger.Info("Shutdown app...")

	ctx, shutdown := context.WithTimeout(context.Background(), time.Second*5)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func (a *App) validatePort(port string) error {
	if port == "" {
		a.logger.WithFields(logrus.Fields{
			"port": port,
		}).Error("port is empty")
		return fmt.Errorf("port is empty")
	}

	if _, err := strconv.ParseInt(port, 10, 32); err != nil {
		a.logger.WithFields(logrus.Fields{
			"port": port,
		}).Error("can't convert port to int")
		return fmt.Errorf("can't convert port to int")
	}
	return nil
}
