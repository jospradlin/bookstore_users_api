package app

import (
	"github.com/gin-gonic/gin"
)

const (

  // PORT for REST API service
  PORT string = ":8080"

)

var (
	router = gin.Default()
)

// StartApplication start application
func StartApplication() {
	MapUrls()
	router.Run(PORT)
}
