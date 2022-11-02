package sitemap

import "encoding/xml"

// Platform stores allowed and restricted platforms.
type Platform struct {
	XMLName      xml.Name  `xml:"video:restriction,omitempty"`
	Relationship *xml.Attr `xml:",attr,omitempty"`
	Value        string    `xml:",chardata"`
}

// PlatformName is a string name for a platform.
type PlatformName string

const (
	Web    PlatformName = "web"
	Mobile PlatformName = "mobile"
	TV     PlatformName = "tv"
)
