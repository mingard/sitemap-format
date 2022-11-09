package sitemap

// sitemap formatting with syntactic sugar. © Arthur Mingard 2022

import (
	"encoding/xml"
	"time"
)

// Location is a location block to be nested under the parent node.
type Location struct {
	XMLName          xml.Name   `xml:"url,omitempty"`
	Location         string     `xml:"loc,omitempty"`
	ChangeFrequency  string     `xml:"changefreq,omitempty"`
	Priority         string     `xml:"priority,omitempty"`
	LastModifiedDate customDate `xml:"lastmod,omitempty"`
	News             *News      `xml:"news:news,omitempty"`
	Images           []*Image   `xml:"image:image,omitempty"`
	Videos           []*Video   `xml:"video:video,omitempty"`
}

// isSitemapIndex sets the XML node name to 'sitemap'
func (l *Location) isSitemapIndex() {
	l.XMLName.Local = "sitemap"
}

// SetLocation sets the url's location parameter.
func (l *Location) SetLocation(loc string) {
	l.Location = loc
}

// SetLastModified sets the value of the modified date field.
func (l *Location) SetLastModified(t time.Time) {
	l.LastModifiedDate.Time = t.UTC()
}

// SetNews sets the single news sub-node.
func (l *Location) SetNews(n *News) {
	l.News = n
}

// AddImage adds an image sub-node.
func (l *Location) AddImage(i ...*Image) {
	l.Images = append(l.Images, i...)
}

// AddVideo adds an video sub-node.
func (l *Location) AddVideo(v ...*Video) {
	l.Videos = append(l.Videos, v...)
}

// defaultLoc creates a default entity with required values.
func defaultLoc() *Location {
	now := time.Now()
	url := &Location{
		XMLName: xml.Name{Local: "url"},
		Images:  make([]*Image, 0),
		Videos:  make([]*Video, 0),
	}
	url.SetLastModified(now)
	return url
}

// NewLoc returns a new instance of the default loc.
func NewLoc() *Location {
	return defaultLoc()
}
