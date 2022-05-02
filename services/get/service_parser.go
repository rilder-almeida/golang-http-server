package get

import (
	"encoding/json"

	"github.com/golang-http-server/entities/httpmessage"
)

func HttpMessageToRequest(httpMessage httpmessage.HttpMessage) (Request, error) {
	var request Request
	err := json.Unmarshal(httpMessage.BodyData, &request)
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
