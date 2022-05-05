package apiget

import (
	"encoding/json"

	"github.com/golang-http-server/entities/httpmessage"
	"github.com/golang-http-server/services/get"
)

func HttpMessageToRequest(httpMessage httpmessage.HttpMessage) (get.Request, error) {
	var request get.Request
	err := json.Unmarshal(httpMessage.BodyData, &request)
	if err != nil {
		return get.Request{}, err
	}
	return request, nil
}

func ResponseToHttpMessage(response get.Response) (httpmessage.HttpMessage, error) {
	bodyData, err := json.Marshal(response)
	if err != nil {
		return httpmessage.HttpMessage{}, err
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