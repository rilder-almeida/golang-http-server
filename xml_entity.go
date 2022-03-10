package main

import "encoding/xml"

// type XmlDocument struct {
// 	Id    string `xml:"nfeProc>NFe>infNFe->Id"`
// 	Cnpj  string `xml:"nfeProc>NFe>infNFe>emit>CNPJ"`
// 	Total string `xml:"nfeProc>NFe>infNFe>total>ICMSTot>vNF"`
// }

type XmlDocument struct {
	XMLName xml.Name `xml:"nfeProc"`
	NFe     NFeField `xml:"NFe"`
}

type NFeField struct {
	XMLName xml.Name    `xml:"NFe"`
	InfNFe  InfNFeField `xml:"infNFe"`
}

type InfNFeField struct {
	XMLName xml.Name   `xml:"infNFe"`
	Id      string     `xml:"Id,attr"`
	Emit    EmitField  `xml:"emit"`
	Total   TotalField `xml:"total"`
}

type EmitField struct {
	XMLName xml.Name `xml:"emit"`
	CNPJ    string   `xml:"CNPJ"`
}

type TotalField struct {
	XMLName xml.Name     `xml:"total"`
	ICMSTot ICMSTotField `xml:"ICMSTot"`
}

type ICMSTotField struct {
	XMLName xml.Name `xml:"ICMSTot"`
	VNF     string   `xml:"vNF"`
}
