package app

import (
	"encoding/xml"
	"io"
)

// Root - GnuCash xml root element
type Root struct {
	XMLName xml.Name    `xml:"gnc-v2"`
	Book    GnuCashBook `xml:"book"`
}

// GnuCashBook - subset of GnuCash xml book element
type GnuCashBook struct {
	Accounts     []GnuCashAccount     `xml:"account"`
	Transactions []GnuCashTransaction `xml:"transaction"`
}

// GnuCashAccount - subset of GnuCash xml account element
type GnuCashAccount struct {
	Name   string `xml:"name"`
	Id     string `xml:"id"` // nolint: golint
	Type   string `xml:"type"`
	Parent string `xml:"parent"`
}

// GnuCashTransaction - subset of GnuCash xml transaction element
type GnuCashTransaction struct {
	Id          string         `xml:"id"` // nolint: golint
	Description string         `xml:"description"`
	DatePosted  string         `xml:"date-posted>date"`
	DateEntered string         `xml:"date-entered>date"`
	Splits      []GnuCashSplit `xml:"splits>split"`
}

// GnuCashSplit - subset of GnuCash xml split element
type GnuCashSplit struct {
	Id      string `xml:"id"` // nolint: golint
	Value   string `xml:"value"`
	Account string `xml:"account"`
}

func Deserialize(reader io.Reader) *Root { // nolint: golint
	var root Root
	decoder := xml.NewDecoder(reader)
	err := decoder.Decode(&root)
	if err != nil {
		panic(err)
	}
	return &root
}
