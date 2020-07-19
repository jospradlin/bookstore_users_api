package users

import (
	"fmt"
	"github.com/jospradlin/bookstore_users_api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)



// Get User from Datasource
func (user *User) Get() *errors.RestErr {
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.ID))
	}
	
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	
	return nil
}


// Save User into Datasource
func (user *User) Save() *errors.RestErr {
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprint("Email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprint("User %d already exists", user.ID))
	}

	usersDB[user.ID] = user
	return nil
}