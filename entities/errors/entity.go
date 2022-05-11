package errors

import (
	"encoding/json"
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

func StatusCode(err error) int {
	var parsedError Error
	json.Unmarshal([]byte(err.Error()), &parsedError)

	switch parsedError.ErrorCode {
	case "FAILED_INSERT_XML":
		return 500
	case "FAILED_GET_NFE":
		return 500
	case "ID_IS_EMPTY":
		return 400
	case "XML_IS_EMPTY":
		return 400
	case "INVALID_REQUEST":
		return 400
	default:
		return 500
	}
}
