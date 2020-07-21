package headers


import (
	"strings"
	"github.com/gin-gonic/gin"
)

// PassthroughIstioXHeaders passes all X-<headers> through to the Response for Tracing
func PassthroughIstioXHeaders(c *gin.Context) {
	for key, element := range c.Request.Header {
		if strings.HasPrefix(key, "x-") || strings.HasPrefix(key, "X-") {
			c.Header(key, element[0])
		}
	}
}