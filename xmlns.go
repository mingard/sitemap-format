package sitemap

// sitemap formatting with syntactic sugar. © Arthur Mingard 2022

import "encoding/xml"

var (
	// XMLNSDefault is the default XML Namespace type.
	XMLNSDefault = &xml.Attr{
		Name:  xml.Name{Local: "xmlns"},
		Value: "http://www.sitemaps.org/schemas/sitemap/0.9",
	}
	// XMLNSNews is the XML Namespace type for news.
	XMLNSNews = &xml.Attr{
		Name:  xml.Name{Local: "xmlns:news"},
		Value: "http://www.google.com/schemas/sitemap-news/0.9",
	}
	// XMLNSXHTML is the XML Namespace type for XHTML.
	XMLNSXHTML = &xml.Attr{
		Name:  xml.Name{Local: "xmlns"},
		Value: "http://www.w3.org/1999/xhtml",
	}
	// XMLNSMobile is the XML Namespace type for mobile-only content.
	XMLNSMobile = &xml.Attr{
		Name:  xml.Name{Local: "xmlns:mobile"},
		Value: "http://www.google.com/schemas/sitemap-mobile/1.0",
	}
	// XMLNSImage is the XML Namespace type for images.
	XMLNSImage = &xml.Attr{
		Name:  xml.Name{Local: "xmlns:image"},
		Value: "http://www.google.com/schemas/sitemap-image/1.1",
	}
	// XMLNSVideo is the XML Namespace type for videos.
	XMLNSVideo = &xml.Attr{
		Name:  xml.Name{Local: "xmlns:image"},
		Value: "http://www.google.com/schemas/sitemap-video/1.1",
	}
)
