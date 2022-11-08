package sitemap

// sitemap formatting with syntactic sugar. © Arthur Mingard 2022

import (
	"encoding/xml"
)

// parent is the parent node.
type parent struct {
	XMLName    xml.Name    `xml:""`
	Attributes []*xml.Attr `xml:",attr,omitempty"`
	Locations  []*Location `xml:",omitempty"`
}

func (p *parent) isSitemapIndex() {
	p.XMLName.Local = "sitemapindex"
}

// addXMLNS adds an XMLNS namespace type.
func (p *parent) addXMLNS(t *xml.Attr) {
	p.Attributes = append(p.Attributes, t)
}

// Add inserts a node into the XML's parent node.
func (p *parent) add(l *Location) {
	p.Locations = append(p.Locations, l)
}

// defaultParent creates a default location entity with required values.
func defaultParent() *parent {
	return &parent{
		XMLName:    xml.Name{Local: "urlset"},
		Attributes: []*xml.Attr{XMLNSDefault},
		Locations:  make([]*Location, 0),
	}
}

// NewIndex returns a new instance of the default SitemapIndex.
func newParent() *parent {
	return defaultParent()
}
