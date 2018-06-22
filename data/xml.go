package data

import "encoding/xml"

type Book struct {
	XMLName          xml.Name `xml:"book"`
	ID               string   `xml:"id,attr"`
	Title            string   `xml:"title"`
	Authors          []string `xml:"authors>author"`
	URL              string   `xml:"url"`
	ISBN             string   `xml:"isbn"`
	ImageURLTemplate string   `xml:"image-url-template"`
}

type Chapter struct {
	XMLName   xml.Name `xml:"chapter"`
	ID        string   `xml:"id,attr"`
	Number    string   `xml:"number"`
	Title     string   `xml:"title"`
	Subtitle  string   `xml:"subtitle"`
	FirstPage int      `xml:"firstPage"`
	LastPage  int      `xml:"lastPage"`
}

type Page struct {
	XMLName       xml.Name `xml:"page"`
	ID            string   `xml:"id,attr"`
	VisibleNumber string   `xml:"visibleNumber"`
	ActualNumber  int      `xml:"actualNumber"`
}
