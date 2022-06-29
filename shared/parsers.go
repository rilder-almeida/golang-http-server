package shared

import (
	"encoding/base64"
	"encoding/json"
	encodingXml "encoding/xml"
	"net/http"

	"github.com/golang-http-server/entities/nfe"
	"github.com/golang-http-server/entities/xml"
)

// func ToNFeDocument(data []byte) (nfe.NFeDocument, error) {
// 	var nfeDocument nfe.NFeDocument
// 	err := json.Unmarshal(data, &nfeDocument)
// 	if err != nil {
// 		return nfe.NFeDocument{}, err
// 	}
// 	return nfeDocument, nil
// }

func ToNFeDocuments(data []byte) (nfe.NFeDocuments, error) {
	var NFeDocuments nfe.NFeDocuments
	err := json.Unmarshal(data, &NFeDocuments)
	if err != nil {
		return nfe.NFeDocuments{}, err
	}
	return NFeDocuments, nil
}

func FromNFeDocuments(NFeDocuments nfe.NFeDocuments) ([]byte, error) {
	data, err := json.Marshal(NFeDocuments)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ToXmlDocument(data []byte) (xml.XmlDocument, error) {
	var xmlDocument xml.XmlDocument

	data32, err := FromBase64ToBase32(data)
	if err != nil {
		return xml.XmlDocument{}, err
	}

	err = encodingXml.Unmarshal(data32, &xmlDocument)
	if err != nil {
		return xml.XmlDocument{}, err
	}

	return xmlDocument, nil
}

func FromBase64ToBase32(data []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(data))
}

// func FromBase32ToBase64(data []byte) ([]byte, error) {
// 	return []byte(base64.StdEncoding.EncodeToString(data)), nil
// }

func EncodeJSONResponse(w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
