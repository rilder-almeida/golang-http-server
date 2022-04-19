package get

import (
	"github.com/golang-http-server/entities"
)

type Response struct {
	RawXml         string               `json:"raw_xml"`
	NfeXmlDocument entities.XmlDocument `json:"nfe_xml_document"`
}
