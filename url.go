package sitemap

import (
	"encoding/xml"
	"time"
)

// Url is a url block to be nested under UrlSet.
type Url struct {
	XMLName          xml.Name  `xml:"url,omitempty"`
	Location         string    `xml:"loc,omitempty"`
	ChangeFrequency  string    `xml:"changefreq,omitempty"`
	Priority         string    `xml:"priority,omitempty"`
	LastModifiedDate time.Time `xml:"lastmod,omitempty"`
	News             *News     `xml:"news:news,omitempty"`
	Image            []*Image  `xml:"image:image,omitempty"`
	Video            []*Video  `xml:"video:video,omitempty"`
}

// SetLocation sets the url's location parameter
func (u *Url) SetLocation(l string) *Url {
	u.Location = l
	return u
}

// SetLastModified sets the value of the modified date field.
func (u *Url) SetLastModified(t time.Time) *Url {
	u.LastModifiedDate = t.UTC()
	return u
}

// defaultUrl creates a default url entity with required values.
func defaultUrl() *Url {
	now := time.Now()
	url := &Url{
		Image: make([]*Image, 0),
		Video: make([]*Video, 0),
	}
	return url.SetLastModified(now)
}

// NewUrl returns a new instance of the default URL.
func NewUrl() *Url {
	return defaultUrl()
}
