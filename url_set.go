package sitemap

// sitemap formatting with syntactic sugar. © Arthur Mingard 2022

import "encoding/xml"

// URLSet is a parent node for <url>'s.
type URLSet struct {
	XMLName    xml.Name    `xml:"urlset"`
	Attributes []*xml.Attr `xml:",attr,omitempty"`
	URL        []ChildNode `xml:"url,omitempty"`
}

// SetType sets the urlset xml namespace type.
func (u *URLSet) SetType(t *xml.Attr) {
	u.Attributes = append(u.Attributes, t)
}

// AddEntry inserts a URL node into the XML's URLSet node.
func (u *URLSet) AddEntry(url ChildNode) {
	u.URL = append(u.URL, url)
}

// defaultURLSet creates a default urlset entity with required values.
func defaultURLSet() ParentNode {
	return &URLSet{
		Attributes: []*xml.Attr{XMLNSDefault},
		URL:        make([]ChildNode, 0),
	}
}

// NewURLSet returns a new instance of the default URLSet.
func NewURLSet() ParentNode {
	return defaultURLSet()
}
