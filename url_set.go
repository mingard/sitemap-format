package sitemap

import "encoding/xml"

type UrlSet struct {
	XMLName    xml.Name    `xml:"urlset"`
	Attributes []*xml.Attr `xml:",attr,omitempty"`
	Url        []*Url      `xml:"url,omitempty"`
}

// SetType sets the urlset xml namespace type.
func (u *UrlSet) SetType(t *xml.Attr) {
	u.Attributes = append(u.Attributes, t)
}

// AddUrl inserts a URL node into the XML's UrlSet node.
func (u *UrlSet) AddUrl(url *Url) {
	u.Url = append(u.Url, url)
}

// defaultUrlSet creates a default urlset entity with required values.
func defaultUrlSet() *UrlSet {
	return &UrlSet{
		Attributes: []*xml.Attr{XMLNSDefault},
		Url:        make([]*Url, 0),
	}
}

// NewUrlSet returns a new instance of the default UrlSet.
func NewUrlSet() *UrlSet {
	return defaultUrlSet()
}
