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
	testDefaultXML       string = `<xml version="1.0" encoding="UTF-8"><urlset></urlset></xml>`
	testDefaultPrettyXML string = "<xml version=\"1.0\" encoding=\"UTF-8\">\nXXXXX<urlset></urlset>\n</xml>"
)

func getTag(str string, tags string) (*xmlquery.Node, error) {
	doc, err := xmlquery.Parse(strings.NewReader(str))
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

	tag, _ := getTag(out, "//urlset/url/loc")
	assert.True(t, tag == nil, "Should have urlset, nested url nodes but without location")
	lastModTag, _ := getTag(out, "//urlset/url/lastmod")
	assert.True(t, lastModTag != nil, "Should have default last modified")
}

func TestXMLAddUrlWithLocation(t *testing.T) {
	xml := New()
	url := NewUrl()

	url.SetLocation("https://domain.com")
	xml.AddUrl(url)

	out, _ := xml.OutputString()
	tag, _ := getTag(out, "//urlset/url/loc")
	assert.True(t, tag != nil, "Should have urlset,nested url nodes, and a location field")
}

func TestXMLAddUrlWithLastModified(t *testing.T) {
	xml := New()
	url := NewUrl()

	now := time.Now()
	later := now.Add(3 * time.Hour)
	url.SetLastModified(later)
	xml.AddUrl(url)

	out, _ := xml.OutputString()

	tag, _ := getTag(out, "//urlset/url/lastmod")
	assert.True(t, tag != nil, "Should have lastmod set")

	nodeTime, _ := time.Parse(time.RFC3339, tag.InnerText())
	assert.Equal(t, nodeTime, later.UTC(), "Should match input timestamp")
}
