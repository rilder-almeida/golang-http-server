package internal

import "github.com/golang-http-server/entities/xml"

type GetHTTPResponse struct {
	Body struct {
		NFeXmlDocument xml.XmlDocument
	}
}
