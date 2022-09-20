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

// SitemapXML is the top level XML node for the sitemap
type SitemapXML struct {
	XMLName      xml.Name    `xml:"xml"`
	Attributes   []*xml.Attr `xml:",attr,omitempty"`
	UrlSet       []*Url      `xml:"urlset>url,omitempty"`
	Version      string      `xml:"version,attr,omitempty"`
	Encoding     string      `xml:"encoding,attr,omitempty"`
	pretty       bool
	outputPrefix string
	outputIndent string
}

// AddUrl inserts a URL node into the XML's UrlSet node.
func (s *SitemapXML) AddUrl(u *Url) {
	s.UrlSet = append(s.UrlSet, u)
}

// Output returns the output value as bytes
func (s *SitemapXML) Output() ([]byte, error) {
	if s.pretty {
		return xml.MarshalIndent(s, s.outputPrefix, s.outputIndent)
	}
	return xml.Marshal(s)
}

// Pretty sets the output to use a prettified format.
func (s *SitemapXML) PrettyFormat(prefix, indent string) {
	s.pretty = true
	s.outputPrefix = prefix
	s.outputIndent = indent
}

// OutputString returns the output as a string. Empty if there's an error.
func (s *SitemapXML) OutputString() (string, error) {
	out, err := s.Output()
	return string(out), err
}

func (s *SitemapXML) SetType(t *xml.Attr) {
	s.Attributes = append(s.Attributes, t)
}

// OutputPrettyString returns the output as a string with prettified rules. Empty if there's an error.
func (s *SitemapXML) OutputPrettyString(prefix, indent string) (string, error) {
	s.PrettyFormat(prefix, indent)
	out, err := s.Output()
	return string(out), err
}

// defaultXML creates a default xml entity with required values.
func defaultXML() *SitemapXML {
	return &SitemapXML{
		Version:      defaultVersion,
		Encoding:     defaultEncoding,
		outputPrefix: defaultOutputPrefix,
		outputIndent: defaultOutputIndent,
		Attributes:   []*xml.Attr{TypeDefault},
	}
}

// New returns a new instance of the default XML.
func New() *SitemapXML {
	return defaultXML()
}
