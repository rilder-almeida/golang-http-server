package nfe

import "github.com/golang-http-server/entities/xml"

type NFeDocument struct {
	RawXml         string          `json:"raw_xml"`
	NFeXmlDocument xml.XmlDocument `json:"nfe_xml_document"`
}
