package entities

import "encoding/xml"

type XmlDocument struct {
	XMLName xml.Name `xml:"nfeProc" json:"-"`
	NFe     NFeField `xml:"NFe"`
}

type NFeField struct {
	XMLName xml.Name    `xml:"NFe" json:"-"`
	InfNFe  InfNFeField `xml:"infNFe"`
}

type InfNFeField struct {
	XMLName xml.Name   `xml:"infNFe" json:"-"`
	Id      string     `xml:"Id,attr"`
	Emit    EmitField  `xml:"emit"`
	Total   TotalField `xml:"total"`
}

type EmitField struct {
	XMLName xml.Name `xml:"emit" json:"-"`
	CNPJ    string   `xml:"CNPJ"`
}

type TotalField struct {
	XMLName xml.Name     `xml:"total" json:"-"`
	ICMSTot ICMSTotField `xml:"ICMSTot"`
}

type ICMSTotField struct {
	XMLName xml.Name `xml:"ICMSTot" json:"-"`
	VNF     string   `xml:"vNF"`
}
