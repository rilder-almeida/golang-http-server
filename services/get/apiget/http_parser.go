package apiget

import (
	"encoding/json"

	customerrors "github.com/golang-http-server/entities/errors"
	"github.com/golang-http-server/entities/httpmessage"
	"github.com/golang-http-server/services/get"
)

func HttpMessageToRequest(httpMessage httpmessage.HttpMessage) (get.Request, error) {
	var request get.Request
	err := json.Unmarshal(httpMessage.BodyData, &request)
	if err != nil {
		return get.Request{}, customerrors.New("INVALID_REQUEST", "Can not unmarshal the request body", err)
	}
	return request, nil
}

func ResponseToHttpMessage(response get.Response) (httpmessage.HttpMessage, error) {
	bodyData, err := json.Marshal(response)
	if err != nil {
		return httpmessage.HttpMessage{}, customerrors.New("INVALID_RESPONSE", "Can not marshal the response body", err)
	}
	return httpmessage.HttpMessage{
		BodyData: bodyData,
	}, nil
}

// func JSONMarshal(t interface{}) ([]byte, error) {
// 	buffer := &bytes.Buffer{}
// 	encoder := json.NewEncoder(buffer)
// 	encoder.SetEscapeHTML(false)
// 	encoder.SetIndent("", "  ")
// 	err := encoder.Encode(t)
// 	return buffer.Bytes(), err
// }
