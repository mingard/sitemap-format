package sitemap

// See https://developers.google.com/search/docs/crawling-indexing/sitemaps/news-sitemap

import (
	"encoding/xml"
	"time"
)

// News stores news entry data.
type News struct {
	XMLName         xml.Name         `xml:"news:news"`
	Publication     *NewsPublication `xml:"news:publication"`
	PublicationDate time.Time        `xml:"news:publication_date"`
	NewsTitle       string           `xml:"news:title"`
}

// SetTitle sets the news extensions title parameter.
func (n *News) SetTitle(t string) *News {
	n.NewsTitle = t
	return n
}

// SetPublicationDate sets the news extensions PublicationDate parameter.
func (n *News) SetPublicationDate(t time.Time) *News {
	n.PublicationDate = t.UTC()
	return n
}

// SetLanguage sets the news extensions Publication.NewsLanguage parameter.
// Must be Lang.
func (n *News) SetLanguage(l Lang) *News {
	n.Publication.NewsLanguage = l
	return n
}

// NewsPublication is the publication sub-component to News.
type NewsPublication struct {
	XMLName      xml.Name `xml:"news:publication"`
	NewsName     string   `xml:"news:name,omitempty"`
	NewsLanguage Lang     `xml:"news:language,omitempty"`
}

// defaultNewsExtension creates a default news extension entity with required values.
func defaultNewsExtension() *News {
	now := time.Now()
	n := &News{
		Publication: &NewsPublication{
			NewsLanguage: LanguageDefault,
		},
	}
	return n.SetPublicationDate(now)
}

// NewNewsExtension returns a new instance of the default NewsExtension.
func NewNewsExtension() *News {
	return defaultNewsExtension()
}
