package sitemap

import "encoding/xml"

//SitemapIndex is the primary sitemap index node.
type SitemapIndex struct {
	XMLName    xml.Name    `xml:"sitemapindex"`
	Attributes []*xml.Attr `xml:",attr,omitempty"`
	Sitemap    []*Sitemap  `xml:"sitemap,omitempty"`
}

// SetType sets the sitemapindex xml namespace type.
func (s *SitemapIndex) SetType(t *xml.Attr) {
	s.Attributes = append(s.Attributes, t)
}

// AddSitemap inserts a sitemap node into the XML's Sitemap node.
func (s *SitemapIndex) AddSitemap(sm *Sitemap) {
	s.Sitemap = append(s.Sitemap, sm)
}

// defaultSitemapIndex creates a default sitemapindex entity with required values.
func defaultSitemapIndex() *SitemapIndex {
	return &SitemapIndex{
		Attributes: []*xml.Attr{XMLNSDefault},
		Sitemap:    make([]*Sitemap, 0),
	}
}

// NewSitemapIndex returns a new instance of the default SitemapIndex.
func NewSitemapIndex() *SitemapIndex {
	return defaultSitemapIndex()
}
