package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/jospradlin/bookstore_users_api/utils/headers"
)

// Ping is a healthcheck for the Service
func Ping (c *gin.Context) {
	headers.PassthroughIstioXHeaders(c)
	c.String(http.StatusOK, "pong")
}