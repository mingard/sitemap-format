package sitemap

import "encoding/xml"

type News struct {
	XMLName         xml.Name         `xml:"news:news,omitempty"`
	NewsTitle       string           `xml:"news:title,omitempty"`
	PublicationDate string           `xml:"news:publication_date"`
	Publication     *NewsPublication `xml:"news:publication"`
}

type NewsPublication struct {
	XMLName      xml.Name `xml:"news:publication"`
	NewsName     string   `xml:"news:name,omitempty"`
	NewsLanguage string   `xml:"news:language,omitempty"`
}
