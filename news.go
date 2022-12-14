package sitemap

// sitemap formatting with syntactic sugar. © Arthur Mingard 2022
// See https://developers.google.com/search/docs/crawling-indexing/sitemaps/news-sitemap

import (
	"encoding/xml"
	"time"
)

// News stores news entry data.
type News struct {
	XMLName         xml.Name     `xml:"news:news"`
	Publication     *Publication `xml:"news:publication"`
	PublicationDate customDate   `xml:"news:publication_date"`
	Title           string       `xml:"news:title"`
}

// SetTitle sets the news extensions title parameter.
func (n *News) SetTitle(t string) *News {
	n.Title = t
	return n
}

// SetName sets the news extensions publication name parameter.
func (n *News) SetName(na string) *News {
	n.Publication.Name = na
	return n
}

// SetPublicationDate sets the news extensions PublicationDate parameter.
func (n *News) SetPublicationDate(c customDate) *News {
	n.PublicationDate = c
	return n
}

// SetLanguage sets the news extensions Publication.NewsLanguage parameter.
// Must be Lang.
func (n *News) SetLanguage(l Lang) *News {
	n.Publication.NewsLanguage = l
	return n
}

// Publication is the news publication sub-component to News.
type Publication struct {
	XMLName      xml.Name `xml:"news:publication"`
	Name         string   `xml:"news:name,omitempty"`
	NewsLanguage Lang     `xml:"news:language,omitempty"`
}

// defaultNews creates a default news extension entity with required values.
func defaultNews() *News {
	now := time.Now()
	n := &News{
		Publication: &Publication{
			NewsLanguage: LangDefault,
		},
	}
	return n.SetPublicationDate(Date(now))
}

// NewNews returns a new instance of the default News extension.
func NewNews() *News {
	return defaultNews()
}
