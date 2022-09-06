package sitemap

import (
	"encoding/xml"
	"fmt"
)

const (
	defaultVersion  string = "1.0"
	defaultEncoding string = "UTF-8"
)

type XML struct {
	XMLName  xml.Name `xml:"xml"`
	UrlSet   []*Url   `xml:"urlset,omitempty"`
	Version  string   `xml:"version,attr,omitempty"`
	Encoding string   `xml:"encoding,attr,omitempty"`
}

func (x *XML) AddUrl(u *Url) {
	x.UrlSet = append(x.UrlSet, u)
}

func (x *XML) Output() ([]byte, error) {
	fmt.Println(len(x.UrlSet))
	return xml.MarshalIndent(x, "", "   ")
}

func defaultXML() *XML {
	return &XML{Version: defaultVersion, Encoding: defaultEncoding}
}

func New() *XML {
	return defaultXML()
}
