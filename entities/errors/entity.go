package errors

import (
	"fmt"
)

type Error struct {
	ErrorCode        string `json:"errorCode"`
	Message          string `json:"message"`
	ApplicationError error  `json:"applicationError"`
}

func (e Error) Error() string {
	return fmt.Sprintf(`
		{
			"errorCode": "%s",
			"message": "%s",
			"applicationError": "%s"
		}`, e.ErrorCode, e.Message, e.ApplicationError)
}

func New(errorCode string, message string, applicationError error) Error {
	return Error{
		ErrorCode:        errorCode,
		Message:          message,
		ApplicationError: applicationError,
	}
}
