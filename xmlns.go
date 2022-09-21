package sitemap

import "encoding/xml"

var (
	// TypeDefault is the default XML Namespace type.
	TypeDefault = xml.Attr{
		Name:  xml.Name{Local: "xmlns"},
		Value: "http://www.sitemaps.org/schemas/sitemap/0.9",
	}
	// TypeNews is the XML Namespace type for news.
	TypeNews = xml.Attr{
		Name:  xml.Name{Local: "xmlns:news"},
		Value: "http://www.google.com/schemas/sitemap-news/0.9",
	}
	// TypeXHTML is the XML Namespace type for XHTML.
	TypeXHTML = xml.Attr{
		Name:  xml.Name{Local: "xmlns"},
		Value: "http://www.w3.org/1999/xhtml",
	}
	// TypeMobile is the XML Namespace type for mobile-only content.
	TypeMobile = xml.Attr{
		Name:  xml.Name{Local: "xmlns:mobile"},
		Value: "http://www.google.com/schemas/sitemap-mobile/1.0",
	}
	// TypeImage is the XML Namespace type for images.
	TypeImage = xml.Attr{
		Name:  xml.Name{Local: "xmlns:image"},
		Value: "http://www.google.com/schemas/sitemap-image/1.1",
	}
	// TypeVideo is the XML Namespace type for videos.
	TypeVideo = xml.Attr{
		Name:  xml.Name{Local: "xmlns:image"},
		Value: "http://www.google.com/schemas/sitemap-video/1.1",
	}
)
