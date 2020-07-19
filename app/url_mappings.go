package app

import (
	"github.com/jospradlin/bookstore_users_api/controllers/ping"
	"github.com/jospradlin/bookstore_users_api/controllers/users"
)

// MapUrls provides the REST API routing
func MapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	// FIX LATER router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", users.CreateUser)


}

