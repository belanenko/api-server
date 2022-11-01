package http

import (
	"github.com/gin-gonic/gin"

	"github.com/belanenko/api-server/internal/server/info"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc info.UseCase) {
	h := NewHandler()

	{
		router.GET("ping", h.Ping)
		router.GET("uptime", h.Uptime)
	}
}
