package server

import (
	"context"
	"fmt"
	"net/http"
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

func (a *App) Run(port int) error {
	if port <= 0 {
		return fmt.Errorf("invalid port value: %d", port)
	}
	listenPort := fmt.Sprintf("%d", port)

	router := gin.Default()
	api := router.Group("/api")

	httpinfo.RegisterHTTPEndpoints(api, a.info)

	a.httpServer = &http.Server{
		Addr:    ":" + listenPort,
		Handler: router,
	}

	go a.listen(listenPort)

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
