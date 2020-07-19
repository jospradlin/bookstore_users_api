package users

import (
	"strconv"
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jospradlin/bookstore_users_api/domain/users"
	"github.com/jospradlin/bookstore_users_api/services"
	"github.com/jospradlin/bookstore_users_api/utils/errors"
)

// CreateUser creates a User
func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	// TODO:  Handle error
	// 	return
	// }

	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	// TODO:  Handle json error
	// 	fmt.Println(err.Error())
	// 	return
	// }

	result, saveErr := services.CreateUser( user )
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// GetUser gets a User
func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("Invalid User ID")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser( userID )
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	
	c.JSON(http.StatusOK, user)
}

// SearchUser searches for a User
// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement me!")
// }
