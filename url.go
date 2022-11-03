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
func (u *Url) SetLocation(l string) {
	u.Location = l
}

// SetLastModified sets the value of the modified date field.
func (u *Url) SetLastModified(t time.Time) {
	u.LastModifiedDate = t.UTC()
}

// SetNews sets the news value.
func (u *Url) SetNews(n *News) {
	u.News = n
}

// AddImage adds one or more image block value.
func (u *Url) AddImage(i ...*Image) {
	u.Images = append(u.Images, i...)
}

// AddVideo adds one or more video block value.
func (u *Url) AddVideo(v ...*Video) {
	u.Videos = append(u.Videos, v...)
}

// defaultUrl creates a default url entity with required values.
func defaultUrl() ChildNode {
	now := time.Now()
	url := &Url{
		Images: make([]*Image, 0),
		Videos: make([]*Video, 0),
	}
	url.SetLastModified(now)
	return url
}

// NewUrl returns a new instance of the default URL.
func NewUrl() ChildNode {
	return defaultUrl()
}
