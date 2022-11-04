package sitemap

// sitemap formatting with syntactic sugar. © Arthur Mingard 2022

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
	// Web is the string key for the web platform.
	Web PlatformName = "web"
	// Mobile is the string key for the mobile platform.
	Mobile PlatformName = "mobile"
	// TV is the string key for the tv platform.
	TV PlatformName = "tv"
)
