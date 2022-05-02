package nfe

import "github.com/golang-http-server/entities/xml"

type NfeDocument struct {
	RawXml         string          `json:"raw_xml"`
	NfeXmlDocument xml.XmlDocument `json:"nfe_xml_document"`
}
