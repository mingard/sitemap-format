package sitemap

import (
	"encoding/xml"
	"time"
)

type Node struct {
	nodes []*Node
}

// Url
type Url struct {
	XMLName          xml.Name   `xml:"url,omitempty"`
	Location         string     `xml:"loc,omitempty"`
	ChangeFrequency  string     `xml:"changefreq,omitempty"`
	Priority         string     `xml:"priority,omitempty"`
	LastModifiedDate *time.Time `xml:"lastmod,omitempty"`
}

func (u *Url) SetLocation(l string) *Url {
	u.Location = l
	return u
}

func (u *Url) SetLastModified(t *time.Time) *Url {
	u.LastModifiedDate = t
	return u
}

func defaultUrl() *Url {
	now := time.Now().UTC()
	return new(Url).SetLastModified(&now)
}

func NewUrl() *Url {
	return defaultUrl()
}
