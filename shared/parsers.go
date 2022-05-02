package shared

import (
	"encoding/base64"
	"encoding/json"

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

func FromNfeDocument(nfeDocument nfe.NfeDocument) ([]byte, error) {
	data, err := json.Marshal(nfeDocument)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ToNfeDocumentList(data []byte) (nfe.NfeDocumentList, error) {
	var nfeDocumentList nfe.NfeDocumentList
	err := json.Unmarshal(data, &nfeDocumentList)
	if err != nil {
		return nfe.NfeDocumentList{}, err
	}
	return nfeDocumentList, nil
}

func FromNfeDocumentList(nfeDocumentList nfe.NfeDocumentList) ([]byte, error) {
	data, err := json.Marshal(nfeDocumentList)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ToXmlDocument(data []byte) (xml.XmlDocument, error) {
	var xmlDocument xml.XmlDocument
	err := json.Unmarshal(data, &xmlDocument)
	if err != nil {
		return xml.XmlDocument{}, err
	}
	return xmlDocument, nil
}

func FromXmlDocument(xmlDocument xml.XmlDocument) ([]byte, error) {
	data, err := json.Marshal(xmlDocument)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func FromBase64ToBase32(data []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(data))
}

func FromBase32ToBase64(data []byte) ([]byte, error) {
	return []byte(base64.StdEncoding.EncodeToString(data)), nil
}
