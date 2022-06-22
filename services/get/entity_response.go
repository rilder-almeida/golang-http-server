package get

import "github.com/golang-http-server/entities/xml"

type Response struct {
	NFeXmlDocument xml.XmlDocument `json:"nfe_xml_document"`
}
