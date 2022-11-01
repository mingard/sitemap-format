package sitemap

import (
	"encoding/xml"
	"fmt"
)

const (
	defaultVersion      string = "1.0"
	defaultEncoding     string = "UTF-8"
	defaultOutputPrefix string = ""
	defaultOutputIndent string = "   "
)

// SitemapXML is the top level XML node for the sitemap
type SitemapXML struct {
	UrlSet       []*UrlSet `xml:"urlset,omitempty"`
	pretty       bool
	outputPrefix string
	outputIndent string
}

// AddUrl inserts a URL node into the XML's UrlSet node.
func (s *SitemapXML) AddUrl(u *Url) {
	s.addDefaultUrlSet()
	s.UrlSet[0].AddUrl(u)
}

// SetType sets the urlset type and creates a default urlset if necessary.
func (s *SitemapXML) SetType(t *xml.Attr) {
	s.addDefaultUrlSet()
	s.UrlSet[0].SetType(t)
}

// Output returns the output value as bytes
func (s *SitemapXML) Output() ([]byte, error) {
	s.addDefaultUrlSet()

	out := []byte(s.headerString())
	var err error
	var marshalledXML []byte

	if s.pretty {
		marshalledXML, err = xml.MarshalIndent(s.UrlSet, s.outputPrefix, s.outputIndent)
	} else {
		marshalledXML, err = xml.Marshal(s.UrlSet)
	}

	if err == nil {
		out = append(out, marshalledXML...)
	}

	return out, err
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

// OutputPrettyString returns the output as a string with prettified rules. Empty if there's an error.
func (s *SitemapXML) OutputPrettyString(prefix, indent string) (string, error) {
	s.PrettyFormat(prefix, indent)
	return s.OutputString()
}

func (s *SitemapXML) addDefaultUrlSet() {
	if len(s.UrlSet) == 0 {
		urlSet := NewUrlSet()
		s.UrlSet = append(s.UrlSet, urlSet)
	}
}

// headerString outputs the XML header string
func (s *SitemapXML) headerString() string {
	xmlHeader := fmt.Sprintf(`<?xml version="%s" encoding="%s"?>`, defaultVersion, defaultEncoding)
	if s.pretty {
		return xmlHeader + "\n"
	}
	return xmlHeader
}

// defaultXML creates a default xml entity with required values.
func defaultXML() *SitemapXML {
	return &SitemapXML{
		UrlSet:       make([]*UrlSet, 0),
		outputPrefix: defaultOutputPrefix,
		outputIndent: defaultOutputIndent,
	}
}

// New returns a new instance of the default XML.
func New() *SitemapXML {
	return defaultXML()
}
