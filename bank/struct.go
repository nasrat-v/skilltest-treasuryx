package bank

import "encoding/xml"

type Status struct {
	Id     string
	Status string
}

type Id struct {
	IBAN string `xml:"IBAN"`
}

type CdtrAcct struct {
	Id Id `xml:"Id"`
}

type Dbtr struct {
	Nm       string   `xml:"Nm"`
	CdtrAcct CdtrAcct `xml:"CdtrAcct"`
}

type Cdtr struct {
	Nm       string   `xml:"Nm"`
	CdtrAcct CdtrAcct `xml:"CdtrAcct"`
}

type GrpHdr struct {
	MsgId   string `xml:"MsgId"`
	CreDtTm string `xml:"CreDtTm"`
}

type Document struct {
	XMLName xml.Name `xml:"Document"`
	GrpHdr  GrpHdr   `xml:"GrpHdr"`
	Cdtr    Cdtr     `xml:"Cdtr"`
	Dbtr    Dbtr     `xml:"Dbtr"`
	Amt     float64  `xml:"Amt"`
}
