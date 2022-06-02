package get

import "github.com/golang-http-server/entities/xml"

type Response struct {
	NfeXmlDocument xml.XmlDocument `json:"nfe_xml_document"`
}
