package main

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
)

func FromJsonRequestParser(body []byte, v interface{}) error {
	err := json.Unmarshal(body, &v)
	if err != nil {
		return err
	}
	return nil
}

func ToXmlParser(body64 []byte) (XmlDocument, error) {
	var xmlDoc XmlDocument

	body, err := base64.StdEncoding.DecodeString(string(body64))
	if err != nil {
		return xmlDoc, err
	}

	err = xml.Unmarshal(body, &xmlDoc)
	if err != nil {
		return xmlDoc, err
	}
	return xmlDoc, nil
}
