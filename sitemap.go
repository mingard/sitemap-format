package sitemap

import (
	"encoding/xml"
	"time"
)

// Sitemap is a sitemap block to be nested under SitemapIndex.
type Sitemap struct {
	XMLName          xml.Name  `xml:"sitemap,omitempty"`
	Location         string    `xml:"loc,omitempty"`
	LastModifiedDate time.Time `xml:"lastmod,omitempty"`
}

// SetLocation sets the sitemap's location parameter
func (s *Sitemap) SetLocation(l string) {
	s.Location = l
}

// SetLastModified sets the value of the modified date field.
func (s *Sitemap) SetLastModified(t time.Time) {
	s.LastModifiedDate = t.UTC()
}

// Defaults
func (s *Sitemap) SetNews(n *News)      {}
func (s *Sitemap) AddImage(i ...*Image) {}
func (s *Sitemap) AddVideo(v ...*Video) {}

// defaultSitemap creates a default sitemap entity with required values.
func defaultSitemap() ChildNode {
	now := time.Now()
	sitemap := new(Sitemap)
	sitemap.SetLastModified(now)
	return sitemap
}

// NewSitemap returns a new instance of the default sitemap.
func NewSitemap() ChildNode {
	return defaultSitemap()
}
