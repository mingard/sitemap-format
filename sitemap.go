package sitemap

import (
	"encoding/xml"
)

const (
	defaultVersion      string = "1.0"
	defaultEncoding     string = "UTF-8"
	defaultOutputPrefix string = ""
	defaultOutputIndent string = "   "
)

// XML is the top level XML node
type XML struct {
	XMLName      xml.Name `xml:"xml"`
	UrlSet       []*Url   `xml:"urlset>url,omitempty"`
	Version      string   `xml:"version,attr,omitempty"`
	Encoding     string   `xml:"encoding,attr,omitempty"`
	pretty       bool
	outputPrefix string
	outputIndent string
}

// AddUrl inserts a URL node into the XML's UrlSet node.
func (x *XML) AddUrl(u *Url) {
	x.UrlSet = append(x.UrlSet, u)
}

// Output returns the output value as bytes
func (x *XML) Output() ([]byte, error) {
	if x.pretty {
		return xml.MarshalIndent(x, x.outputPrefix, x.outputIndent)
	}
	return xml.Marshal(x)
}

// Pretty sets the output to use a prettified format.
func (x *XML) PrettyFormat(prefix, indent string) {
	x.pretty = true
	x.outputPrefix = prefix
	x.outputIndent = indent
}

// OutputString returns the output as a string. Empty if there's an error.
func (x *XML) OutputString() (string, error) {
	out, err := x.Output()
	return string(out), err
}

// OutputPrettyString returns the output as a string with prettified rules. Empty if there's an error.
func (x *XML) OutputPrettyString(prefix, indent string) (string, error) {
	x.PrettyFormat(prefix, indent)
	out, err := x.Output()
	return string(out), err
}

// defaultXML creates a default xml entity with required values.
func defaultXML() *XML {
	return &XML{Version: defaultVersion, Encoding: defaultEncoding, outputPrefix: defaultOutputPrefix, outputIndent: defaultOutputIndent}
}

// New returns a new instance of the default XML.
func New() *XML {
	return defaultXML()
}
