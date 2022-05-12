package apiinsert

import (
	"encoding/json"

	customErrors "github.com/golang-http-server/entities/errors"
	"github.com/golang-http-server/entities/httpmessage"
	"github.com/golang-http-server/services/insert"
	"github.com/golang-http-server/shared"
)

func HttpMessageToRequest(httpMessage httpmessage.HttpMessage) (insert.Request, error) {
	var request insert.Request
	err := json.Unmarshal(httpMessage.BodyData, &request)
	if err != nil {
		return insert.Request{}, customErrors.New("INVALID_REQUEST", "Can not unmarshal the request body", err)
	}

	data, err := shared.FromBase64ToBase32([]byte(request.XML))
	if err != nil {
		return request, customErrors.New("INVALID_REQUEST", "Can not parse the request body bytes from base64 to base32", err)
	}
	request.XML = string(data)

	return request, nil
}

func ResponseToHttpMessage(response insert.Response) (httpmessage.HttpMessage, error) {
	bodyData, err := json.Marshal(response)
	if err != nil {
		return httpmessage.HttpMessage{}, customErrors.New("INVALID_RESPONSE", "Can not marshal the response body", err)
	}
	return httpmessage.HttpMessage{
		BodyData: bodyData,
	}, nil
}
