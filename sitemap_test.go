package sitemap

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/antchfx/xmlquery"
	"github.com/stretchr/testify/assert"
)

const (
	testDefaultXML       string = `<xml xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" version="1.0" encoding="UTF-8"><urlset></urlset></xml>`
	testDefaultPrettyXML string = "<xml xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\" version=\"1.0\" encoding=\"UTF-8\">\nXXXXX<urlset></urlset>\n</xml>"
)

func getDoc(str string) (*xmlquery.Node, error) {
	return xmlquery.Parse(strings.NewReader(str))
}

func find(str, search string) ([]*xmlquery.Node, error) {
	doc, err := getDoc(str)
	if err != nil {
		return nil, err
	}

	tags := xmlquery.Find(doc, search)
	if len(tags) == 0 {
		return tags, nil
	}

	return nil, fmt.Errorf(`Failed to find %s`, search)
}

func findOne(str, tags string) (*xmlquery.Node, error) {
	doc, err := getDoc(str)
	if err != nil {
		return nil, err
	}

	tag := xmlquery.FindOne(doc, tags)
	if tag != nil {
		return tag, nil
	}

	return nil, fmt.Errorf(`Failed to find %s`, tags)
}

func TestDefaults(t *testing.T) {
	xml := New()

	out, _ := xml.OutputString()

	assert.Equal(t, testDefaultXML, out, "Should output default empty XML")
}

func TestPrettyOutput(t *testing.T) {
	xml := New()

	out, _ := xml.OutputPrettyString("", "XXXXX")

	assert.Equal(t, testDefaultPrettyXML, out, "Should output prettified XML")
}

func TestXMLAddDefaultUrl(t *testing.T) {
	xml := New()
	url := NewUrl()

	xml.AddUrl(url)

	out, _ := xml.OutputString()

	tag, _ := findOne(out, "//urlset/url/loc")
	assert.True(t, tag == nil, "Should have urlset, nested url nodes but without location")
	lastModTag, _ := findOne(out, "//urlset/url/lastmod")
	assert.True(t, lastModTag != nil, "Should have default last modified")
}

func TestXMLAddUrlWithLocation(t *testing.T) {
	xml := New()
	url := NewUrl()

	url.SetLocation("https://domain.com")
	xml.AddUrl(url)

	out, _ := xml.OutputString()
	tag, _ := findOne(out, "//urlset/url/loc")
	assert.True(t, tag != nil, "Should have urlset,nested url nodes, and a location field")
}

func TestDefaultType(t *testing.T) {
	xml := New()

	out, _ := xml.OutputString()
	doc, _ := findOne(out, "//.")
	assert.Equal(t, doc.Attr[0].Name.Local, "xmlns", "Should contain correct default type")
}

func TestSetType(t *testing.T) {
	xml := New()

	xml.SetType(TypeNews)
	out, _ := xml.OutputString()
	doc, _ := findOne(out, "//.")
	assert.Equal(t, doc.Attr[1].Name.Local, "news", "Should contain correct type")
	// t.Error("False neg")
}

func TestXMLAddUrlWithLastModified(t *testing.T) {
	xml := New()
	url := NewUrl()

	now := time.Now()
	later := now.Add(3 * time.Hour)
	url.SetLastModified(later)
	xml.AddUrl(url)

	out, _ := xml.OutputString()

	tag, _ := findOne(out, "//urlset/url/lastmod")
	assert.True(t, tag != nil, "Should have lastmod set")

	nodeTime, _ := time.Parse(time.RFC3339, tag.InnerText())
	assert.Equal(t, nodeTime, later.UTC(), "Should match input timestamp")
}
