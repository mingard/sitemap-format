package sitemap

// sitemap formatting with syntactic sugar. © Arthur Mingard 2022

import (
	"encoding/xml"
	"time"
)

// ParentNode is the primary node e.g. <urlset> or <sitemap_index>
type ParentNode interface {
	// SetType sets the xmlns attribute.
	SetType(t *xml.Attr)
	// AddEntry adds an entry to the parent node.
	AddEntry(sm ChildNode)
}

// ChildNode is the nested node type e.g. <url> or <sitemap>
type ChildNode interface {
	// SetLocation sets the value of the nodes location block.
	SetLocation(l string)
	// SetLastModified sets the nodes last modified date.
	SetLastModified(t time.Time)
	// SetNews sets the single news sub-node to the child node.
	SetNews(n *News)
	// AddImage adds an image sub-node to the child node.
	AddImage(i ...*Image)
	// AddVideo adds an video sub-node to the child node.
	AddVideo(v ...*Video)
}
