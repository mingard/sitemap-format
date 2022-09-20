package sitemap

import "encoding/xml"

var TypeDefault = &xml.Attr{
	Name:  xml.Name{Local: "xmlns"},
	Value: "http://www.sitemaps.org/schemas/sitemap/0.9",
}

var TypeNews = &xml.Attr{
	Name:  xml.Name{Local: "xmlns:news"},
	Value: "http://www.google.com/schemas/sitemap-news/0.9",
}

var TypeXHTML = &xml.Attr{
	Name:  xml.Name{Local: "xmlns"},
	Value: "http://www.w3.org/1999/xhtml",
}

var TypeMobile = &xml.Attr{
	Name:  xml.Name{Local: "xmlns:mobile"},
	Value: "http://www.google.com/schemas/sitemap-mobile/1.0",
}

var TypeImage = &xml.Attr{
	Name:  xml.Name{Local: "xmlns:image"},
	Value: "http://www.google.com/schemas/sitemap-image/1.1",
}

var TypeVideo = &xml.Attr{
	Name:  xml.Name{Local: "xmlns:image"},
	Value: "http://www.google.com/schemas/sitemap-video/1.1",
}
