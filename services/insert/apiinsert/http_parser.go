package apiinsert

import (
	"encoding/json"

	"github.com/golang-http-server/entities/httpmessage"
	"github.com/golang-http-server/services/insert"
	"github.com/golang-http-server/shared"
)

func HttpMessageToRequest(httpMessage httpmessage.HttpMessage) (insert.Request, error) {
	var request insert.Request
	err := json.Unmarshal(httpMessage.BodyData, &request)
	if err != nil {
		return insert.Request{}, err
	}

	data, err := shared.FromBase64ToBase32([]byte(request.XML))
	if err != nil {
		return request, err
	}
	request.XML = string(data)

	return request, nil
}

func ResponseToHttpMessage(response insert.Response) (httpmessage.HttpMessage, error) {
	bodyData, err := json.Marshal(response)
	if err != nil {
		return httpmessage.HttpMessage{}, err
	}
	return httpmessage.HttpMessage{
		BodyData: bodyData,
	}, nil
}
