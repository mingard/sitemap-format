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
	ParentNodes  []ParentNode `xml:",omitempty"`
	pretty       bool
	outputPrefix string
	outputIndent string
}

// AddUrl inserts a URL node into the XML's UrlSet node.
func (s *SitemapXML) AddEntry(e ChildNode) {
	s.ParentNodes[0].AddEntry(e)
}

// SetType sets the urlset type and creates a default urlset if necessary.
func (s *SitemapXML) SetType(t *xml.Attr) {
	for _, parent := range s.ParentNodes {
		parent.SetType(t)
	}
}

// Output returns the output value as bytes
func (s *SitemapXML) Output() ([]byte, error) {
	out := []byte(s.headerString())
	var err error

	var marshalledXML []byte

	if s.pretty {
		marshalledXML, err = xml.MarshalIndent(s.ParentNodes, s.outputPrefix, s.outputIndent)
	} else {
		marshalledXML, err = xml.Marshal(s.ParentNodes)
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
		ParentNodes:  make([]ParentNode, 0),
		outputPrefix: defaultOutputPrefix,
		outputIndent: defaultOutputIndent,
	}
}

func NewSitemapIndex() *SitemapXML {
	sitemapIndex := NewIndex()
	out := defaultXML()
	out.ParentNodes = append(out.ParentNodes, sitemapIndex)
	return out
}

// New returns a new instance of the default XML.
func New() *SitemapXML {
	urlSet := NewUrlSet()
	out := defaultXML()
	out.ParentNodes = append(out.ParentNodes, urlSet)
	return out
}
