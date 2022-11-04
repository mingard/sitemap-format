package sitemap

// sitemap formatting with syntactic sugar. © Arthur Mingard 2022

import (
	"encoding/xml"
	"time"
)

// URL is a url block to be nested under URLSet.
type URL struct {
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
func (u *URL) SetLocation(l string) {
	u.Location = l
}

// SetLastModified sets the value of the modified date field.
func (u *URL) SetLastModified(t time.Time) {
	u.LastModifiedDate = t.UTC()
}

// SetNews sets the single news sub-node.
func (u *URL) SetNews(n *News) {
	u.News = n
}

// AddImage adds an image sub-node.
func (u *URL) AddImage(i ...*Image) {
	u.Images = append(u.Images, i...)
}

// AddVideo adds an video sub-node.
func (u *URL) AddVideo(v ...*Video) {
	u.Videos = append(u.Videos, v...)
}

// defaultURL creates a default url entity with required values.
func defaultURL() ChildNode {
	now := time.Now()
	url := &URL{
		Images: make([]*Image, 0),
		Videos: make([]*Video, 0),
	}
	url.SetLastModified(now)
	return url
}

// NewUrl returns a new instance of the default URL.
func NewUrl() ChildNode {
	return defaultURL()
}
