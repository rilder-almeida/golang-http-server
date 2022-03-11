package main

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

func urlParser(url string, method string) string {
	switch method {
	case "POST":
		return strings.TrimPrefix(url, "/nfe/v1")

	case "GET":
		return strings.TrimPrefix(url, "/nfe/v1")
	}
	panic(fmt.Sprintf("Bad method: %s", method))
}

func requestBodyParser(bodyRequest io.ReadCloser) ([]byte, error) {
	body64, err := ioutil.ReadAll(bodyRequest)
	if err != nil {
		return nil, err
	}
	return body64, nil
}

func jsonRequestParser(body []byte, v interface{}) error {
	err := json.Unmarshal(body, &v)
	if err != nil {
		return err
	}
	return nil
}

func xmlParser(body64 []byte) (XmlDocument, error) {
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
