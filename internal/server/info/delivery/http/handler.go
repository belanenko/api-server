package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func (h *Handler) Uptime(c *gin.Context) {
	response := map[string]interface{}{
		"uptime": time.Since(startTime).String(),
	}
	j, _ := json.Marshal(response)

	c.String(http.StatusOK, string(j))
}
