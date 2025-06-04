package errhandler

import (
	"encoding/json"
	"net/http"
)

// CustomError is a struct that contains the error code and message

type CustomError struct {
	cause           error
	statusCode      int
	message         string
	isInternalError bool
}

// this is a global unknown error that can be used when error  is found is not of type CustomError
var (
	unknownError = NewCustomError(nil, http.StatusInternalServerError, "Unknown error", true)
)

// implements error interface
func (e *CustomError) Error() string {
	return e.cause.Error()
}

func NewCustomError(cause error, statusCode int, message string, isInternalError bool) *CustomError {
	return &CustomError{
		cause:           cause,
		statusCode:      statusCode,
		message:         message,
		isInternalError: isInternalError,
	}
}

func (e *CustomError) ErrorResponseBody() ([]byte, error) {
	payload := map[string]interface{}{}

	if e.isInternalError {
		payload["message"] = "Something went wrong"
		payload["description"] = "something went wrong internally"
	} else {
		payload["description"] = e.cause
		payload["message"] = e.message
		payload["status"] = e.statusCode
		payload["statusDescription"] = http.StatusText(e.statusCode)
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Setters
func (e *CustomError) SetMessage(message string) *CustomError {
	e.message = message
	return e
}

func (e *CustomError) SetStatusCode(statusCode int) *CustomError {
	e.statusCode = statusCode
	return e
}

func (e *CustomError) SetIsInternalError(isInternalError bool) *CustomError {
	e.isInternalError = isInternalError
	return e
}

func (e *CustomError) SetError(cause error) *CustomError {
	e.cause = cause
	return e
}
