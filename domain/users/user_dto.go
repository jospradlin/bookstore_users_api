package users

import (
	"strings"
	"github.com/jospradlin/bookstore_users_api/utils/errors"
)

// User data object structure
type User struct {
	ID 				int64 	`json:"id"`
	FirstName 		string	`json:"first_name"`
	LastName 		string	`json:"last_name"`
	Email 			string	`json:"email"`
	DateCreated 	string	`json:"date_created"`
}

// Validate the User data object
// func Validate(user *User) *errors.RestErr {
// 	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
// 	if user.Email == "" {
// 		return errors.NewBadRequestError("Invalid email address")
// 	}
// 	return nil
// }

// Validate a method that a User can validate itself
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}
	return nil
}