package sitemap

// sitemap formatting with syntactic sugar. © Arthur Mingard 2022

import "encoding/xml"

var (
	// Allow is the allow attribute.
	Allow = &xml.Attr{
		Name:  xml.Name{Local: "relationship"},
		Value: "allow",
	}
	// Deny is the deny attribute.
	Deny = &xml.Attr{
		Name:  xml.Name{Local: "relationship"},
		Value: "deny",
	}
)
