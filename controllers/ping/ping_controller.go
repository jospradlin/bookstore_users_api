package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping is a healthcheck for the Service
func Ping (c *gin.Context) {
	c.String(http.StatusOK, "pong")
}