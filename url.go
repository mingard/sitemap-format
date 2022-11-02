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
	Images           []*Image  `xml:"image:image,omitempty"`
	Videos           []*Video  `xml:"video:video,omitempty"`
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

// SetNews sets the news value.
func (u *Url) SetNews(n *News) *Url {
	u.News = n
	return u
}

// AddImage adds one or more image block value.
func (u *Url) AddImage(i ...*Image) *Url {
	u.Images = append(u.Images, i...)
	return u
}

// AddVideo adds one or more video block value.
func (u *Url) AddVideo(v ...*Video) *Url {
	u.Videos = append(u.Videos, v...)
	return u
}

// defaultUrl creates a default url entity with required values.
func defaultUrl() *Url {
	now := time.Now()
	url := &Url{
		Images: make([]*Image, 0),
		Videos: make([]*Video, 0),
	}
	return url.SetLastModified(now)
}

// NewUrl returns a new instance of the default URL.
func NewUrl() *Url {
	return defaultUrl()
}
