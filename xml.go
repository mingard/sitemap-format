package sitemap

// sitemap formatting with syntactic sugar. © Arthur Mingard 2022

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

// XML is the top level XML node for the sitemap
type XML struct {
	ParentNodes  []ParentNode `xml:",omitempty"`
	pretty       bool
	outputPrefix string
	outputIndent string
}

// AddEntry inserts an entry into the XML's first parent node.
func (x *XML) AddEntry(e ChildNode) {
	x.ParentNodes[0].AddEntry(e)
}

// SetType sets the parent node's xmlns type attribute.
func (x *XML) SetType(t *xml.Attr) {
	for _, parent := range x.ParentNodes {
		parent.SetType(t)
	}
}

// Output returns the output value as bytes
func (x *XML) Output() ([]byte, error) {
	out := []byte(x.headerString())
	var err error

	var marshalledXML []byte

	if x.pretty {
		marshalledXML, err = xml.MarshalIndent(x.ParentNodes, x.outputPrefix, x.outputIndent)
	} else {
		marshalledXML, err = xml.Marshal(x.ParentNodes)
	}

	if err == nil {
		out = append(out, marshalledXML...)
	}

	return out, err
}

// PrettyFormat sets preferences for the prettified output.
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
	return x.OutputString()
}

// headerString outputs the XML header string
func (x *XML) headerString() string {
	xmlHeader := fmt.Sprintf(`<?xml version="%s" encoding="%s"?>`, defaultVersion, defaultEncoding)
	if x.pretty {
		return xmlHeader + "\n"
	}
	return xmlHeader
}

// defaultXML creates a default xml entity with required values.
func defaultXML() *XML {
	return &XML{
		ParentNodes:  make([]ParentNode, 0),
		outputPrefix: defaultOutputPrefix,
		outputIndent: defaultOutputIndent,
	}
}

// NewSitemapIndex creates a new sitemap with a sitemap_index child node.
func NewSitemapIndex() *XML {
	sitemapIndex := NewIndex()
	out := defaultXML()
	out.ParentNodes = append(out.ParentNodes, sitemapIndex)
	return out
}

// New returns a new instance of the default XML.
func New() *XML {
	urlSet := NewURLSet()
	out := defaultXML()
	out.ParentNodes = append(out.ParentNodes, urlSet)
	return out
}
