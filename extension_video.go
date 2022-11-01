package sitemap

// See https://developers.google.com/search/docs/crawling-indexing/sitemaps/video-sitemaps

import (
	"encoding/xml"
)

// Video stores video entry data.
type Video struct {
	XMLName xml.Name `xml:"video:video"`
}
