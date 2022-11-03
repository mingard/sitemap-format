package sitemap

import "encoding/xml"

//SitemapIndex is the primary sitemap index node.
type SitemapIndex struct {
	XMLName    xml.Name    `xml:"sitemapindex"`
	Attributes []*xml.Attr `xml:",attr,omitempty"`
	Sitemap    []ChildNode `xml:"sitemap,omitempty"`
}

// SetType sets the sitemapindex xml namespace type.
func (s *SitemapIndex) SetType(t *xml.Attr) {
	s.Attributes = append(s.Attributes, t)
}

// AddEntry inserts a sitemap node into the XML's Sitemap node.
func (s *SitemapIndex) AddEntry(sm ChildNode) {
	s.Sitemap = append(s.Sitemap, sm)
}

// defaultSitemapIndex creates a default sitemapindex entity with required values.
func defaultSitemapIndex() ParentNode {
	return &SitemapIndex{
		Attributes: []*xml.Attr{XMLNSDefault},
		Sitemap:    make([]ChildNode, 0),
	}
}

// NewIndex returns a new instance of the default SitemapIndex.
func NewIndex() ParentNode {
	return defaultSitemapIndex()
}
