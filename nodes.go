package sitemap

import (
	"encoding/xml"
	"time"
)

// ParentNode is the primary node e.g. <urlset> or <sitemap_index>
type ParentNode interface {
	SetType(t *xml.Attr)
	AddEntry(sm ChildNode)
}

// Childnode is the nested node type e.g. <url> or <sitemap>
type ChildNode interface {
	SetLocation(l string)
	SetLastModified(t time.Time)
	SetNews(n *News)
	AddImage(i ...*Image)
	AddVideo(v ...*Video)
}
