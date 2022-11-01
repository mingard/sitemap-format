package sitemap

// See https://developers.google.com/search/docs/crawling-indexing/sitemaps/image-sitemaps

import (
	"encoding/xml"
)

// Image stores image entry data.
type Image struct {
	XMLName xml.Name `xml:"image:image"`
}
