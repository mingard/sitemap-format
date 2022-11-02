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
func (s *Sitemap) SetLocation(l string) *Sitemap {
	s.Location = l
	return s
}

// SetLastModified sets the value of the modified date field.
func (s *Sitemap) SetLastModified(t time.Time) *Sitemap {
	s.LastModifiedDate = t.UTC()
	return s
}

// defaultSitemap creates a default sitemap entity with required values.
func defaultSitemap() *Sitemap {
	now := time.Now()
	sitemap := new(Sitemap)
	return sitemap.SetLastModified(now)
}

// NewSitemap returns a new instance of the default sitemap.
func NewSitemap() *Sitemap {
	return defaultSitemap()
}
