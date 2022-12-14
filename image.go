package sitemap

// sitemap formatting with syntactic sugar. © Arthur Mingard 2022
// See https://developers.google.com/search/docs/crawling-indexing/sitemaps/image-sitemaps

import (
	"encoding/xml"
)

// Image stores image entry data.
type Image struct {
	XMLName xml.Name `xml:"image:image"`
	// Loc is the location of the image
	Loc string `xml:"image:loc"`
}

// SetLocation sets the image location.
func (i *Image) SetLocation(l string) *Image {
	i.Loc = l
	return i
}

// NewImage returns a new instance of the default Image extension.
func NewImage() *Image {
	return new(Image)
}
