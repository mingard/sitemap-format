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
	XMLName      xml.Name  `xml:"xml"`
	UrlSet       []*UrlSet `xml:"urlset,omitempty"`
	Version      string    `xml:"version,attr,omitempty"`
	Encoding     string    `xml:"encoding,attr,omitempty"`
	pretty       bool
	outputPrefix string
	outputIndent string
}

// AddUrl inserts a URL node into the XML's UrlSet node.
func (s *SitemapXML) AddUrl(u *Url) {
	s.addDefaultUrlSet()
	s.UrlSet[0].AddUrl(u)
}

// Output returns the output value as bytes
func (s *SitemapXML) Output() ([]byte, error) {
	s.addDefaultUrlSet()
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

// SetType sets the urlset type and creates a default urlset if necessary.
func (s *SitemapXML) SetType(t *xml.Attr) {
	s.addDefaultUrlSet()
	s.UrlSet[0].SetType(t)
}

// OutputPrettyString returns the output as a string with prettified rules. Empty if there's an error.
func (s *SitemapXML) OutputPrettyString(prefix, indent string) (string, error) {
	s.PrettyFormat(prefix, indent)
	out, err := s.Output()
	return string(out), err
}

func (s *SitemapXML) addDefaultUrlSet() {
	if len(s.UrlSet) == 0 {
		urlSet := NewUrlSet()
		s.UrlSet = append(s.UrlSet, urlSet)
	}
}

// defaultXML creates a default xml entity with required values.
func defaultXML() *SitemapXML {
	return &SitemapXML{
		UrlSet:       make([]*UrlSet, 0),
		Version:      defaultVersion,
		Encoding:     defaultEncoding,
		outputPrefix: defaultOutputPrefix,
		outputIndent: defaultOutputIndent,
	}
}

// New returns a new instance of the default XML.
func New() *SitemapXML {
	return defaultXML()
}
