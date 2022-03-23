package main

const (
	POST_RESPONSE_CONTENT_TYPE = "application/json"
	JSON_FILENAME              = "store.json"
)

type NfeStore interface {
	PostRequestReceiver(JsonPostRequest) (JsonResponse, error)
	GetRequestReceiver(JsonGetRequest) (JsonResponse, error)
	RequestResponder(string, []byte, int) JsonResponse
	//utils
	AssertIdIsNew(string) bool
	StoreNfe(JsonPostRequest) ([]byte, error)
	GetNfeById(string) ([]byte, error)
	MakeJsonNfeIsNew(bool) []byte
}

type JsonResponse struct {
	contentType string
	bodyData    []byte
	httpStatus  int
}

type JsonPostRequest struct {
	XML string `json:"XML"`
}

type JsonGetRequest struct {
	Id string `json:"id"`
}

type NfeDocument struct {
	RawXml         string      `json:"raw_xml"`
	NfeXmlDocument XmlDocument `json:"nfe_xml_document"`
}

type JsonNfeIsNew struct {
	IsNewNfe bool `json:"IsNewNfe"`
}
