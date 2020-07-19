package services

import (
	//"strings"
	"github.com/jospradlin/bookstore_users_api/domain/users"
	"github.com/jospradlin/bookstore_users_api/utils/errors"
)

// CreateUser in the data service
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		 return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil

}

// GetUser in the data service
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ ID: userID }
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}