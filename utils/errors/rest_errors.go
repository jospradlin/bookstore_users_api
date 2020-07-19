package errors

import (
	"net/http"
)

// RestErr is a response messaging object
type RestErr struct {
	Message string	`json:"message"`
	Status int		`json:"status"`
	Error string	`json:"error"`
}

// NewBadRequestError is a standardized bad request response data object
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status: http.StatusBadRequest,
		Error: "Bad Request",
	}
}

// NewBadRequestError is a standardized bad request response data object
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status: http.StatusNotFound,
		Error: "Not Found",
	}
}