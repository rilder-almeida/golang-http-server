package shared

import (
	"encoding/base64"
	"encoding/json"
	encodingXml "encoding/xml"

	"github.com/golang-http-server/entities/nfe"
	"github.com/golang-http-server/entities/xml"
)

func ToNfeDocument(data []byte) (nfe.NfeDocument, error) {
	var nfeDocument nfe.NfeDocument
	err := json.Unmarshal(data, &nfeDocument)
	if err != nil {
		return nfe.NfeDocument{}, err
	}
	return nfeDocument, nil
}

// func FromNfeDocument(nfeDocument nfe.NfeDocument) ([]byte, error) {
// 	data, err := json.Marshal(nfeDocument)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return data, nil
// }

func ToNfeDocuments(data []byte) (nfe.NfeDocuments, error) {
	var NfeDocuments nfe.NfeDocuments
	err := json.Unmarshal(data, &NfeDocuments)
	if err != nil {
		return nfe.NfeDocuments{}, err
	}
	return NfeDocuments, nil
}

func FromNfeDocuments(NfeDocuments nfe.NfeDocuments) ([]byte, error) {
	data, err := json.Marshal(NfeDocuments)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ToXmlDocument(data []byte) (xml.XmlDocument, error) {
	var xmlDocument xml.XmlDocument
	err := encodingXml.Unmarshal(data, &xmlDocument)
	if err != nil {
		return xml.XmlDocument{}, err
	}
	return xmlDocument, nil
}

// func FromXmlDocument(xmlDocument xml.XmlDocument) ([]byte, error) {
// 	data, err := json.Marshal(xmlDocument)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return data, nil
// }

func FromBase64ToBase32(data []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(data))
}

func FromBase32ToBase64(data []byte) ([]byte, error) {
	return []byte(base64.StdEncoding.EncodeToString(data)), nil
}
