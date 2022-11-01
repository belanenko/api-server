package info

import "github.com/gin-gonic/gin"

type UseCase interface {
	Ping(c *gin.Context)
	Uptime(c *gin.Context)
}
