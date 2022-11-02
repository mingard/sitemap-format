package sitemap

// See https://developers.google.com/search/docs/crawling-indexing/sitemaps/image-sitemaps

import (
	"encoding/xml"
)

// Image stores image entry data.
type Image struct {
	XMLName xml.Name `xml:"image:image"`
	Loc     string   `xml:"image:loc"`
}

// NewImage returns a new instance of the default Image extension.
func NewImage() *Image {
	return new(Image)
}
