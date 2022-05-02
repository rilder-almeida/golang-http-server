package insert

import (
	"encoding/json"

	"github.com/golang-http-server/entities/httpmessage"
	"github.com/golang-http-server/shared"
)

func HttpMessageToRequest(httpMessage httpmessage.HttpMessage) (Request, error) {
	var request Request
	data, err := shared.FromBase64ToBase32(httpMessage.BodyData)
	if err != nil {
		return request, err
	}
	err = json.Unmarshal(data, &request)
	if err != nil {
		return Request{}, err
	}
	return request, nil
}

func ResponseToHttpMessage(response Response) (httpmessage.HttpMessage, error) {
	bodyData, err := json.Marshal(response)
	if err != nil {
		return httpmessage.HttpMessage{}, err
	}
	return httpmessage.HttpMessage{
		BodyData: bodyData,
	}, nil
}
