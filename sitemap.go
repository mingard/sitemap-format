package sitemap

import (
	"encoding/xml"
	"fmt"
)

const (
	defaultVersion  string = "1.0"
	defaultEncoding string = "UTF-8"
)

// XML is the top level XML node
type XML struct {
	XMLName  xml.Name `xml:"xml"`
	UrlSet   []*Url   `xml:"urlset>url,omitempty"`
	Version  string   `xml:"version,attr,omitempty"`
	Encoding string   `xml:"encoding,attr,omitempty"`
}

// AddUrl inserts a URL node into the XML's UrlSet node.
func (x *XML) AddUrl(u *Url) {
	x.UrlSet = append(x.UrlSet, u)
}

// Output returns the output value as bytes
func (x *XML) Output() ([]byte, error) {
	fmt.Println(len(x.UrlSet))
	return xml.MarshalIndent(x, "", "   ")
}

// OutputString returns the output as a string. Empty if there's an error.
func (x *XML) OutputString() string {
	out, err := x.Output()
	if err != nil {
		return ""
	}
	return string(out)
}

// defaultXML creates a default xml entity with required values.
func defaultXML() *XML {
	return &XML{Version: defaultVersion, Encoding: defaultEncoding}
}

// New returns a new instance of the default XML.
func New() *XML {
	return defaultXML()
}
