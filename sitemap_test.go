package sitemap

import (
	"encoding/xml"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func checkTagExists(src string, tags ...string) (bool, error) {
	decoder := xml.NewDecoder(strings.NewReader(src))

	for {
		t, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				return false, nil
			}
			return false, err
		}
		if se, ok := t.(xml.StartElement); ok {
			if se.Name.Local == tags[0] {
				// Keep checking nested tags
				if len(tags) > 1 {
					return checkTagExists(src, tags[1:]...)
				}

				return true, nil
			}
		}
	}
}

// func getTagCharData(src string) (string, error) {
// 	decoder := xml.NewDecoder(strings.NewReader(src))
// 	for {
// 		t, err := decoder.Token()
// 		if err != nil {
// 			if err == io.EOF {
// 				return "", nil
// 			}
// 			return "", err

// 		}

// 		if se, ok := t.(xml.CharData); ok {
// 			// println(src)
// 			// fmt.Printf(`value %v+`, se)

// 			return string(se), nil
// 		}
// 	}
// }

func TestXMLErrorHandling(t *testing.T) {
	// TODO: This test should assess whether an error is returned properly.
	// xml := New()
	// TODO: Test .Output return of error
	// TODO: Test .OutputString returns empty string.
}

func TestDefaults(t *testing.T) {
	xml := New()

	out := xml.OutputString()

	assert.Equal(t, `<xml version="1.0" encoding="UTF-8">
   <urlset></urlset>
</xml>`, out, "Should output default empty XML")
}

func TestXMLAddDefaultUrl(t *testing.T) {
	xml := New()
	url := NewUrl()

	xml.AddUrl(url)

	out := xml.OutputString()

	locExists, _ := checkTagExists(out, "urlset", "url", "loc")
	assert.False(t, locExists, "Should have urlset, nested url nodes but without location")
	lastModExists, _ := checkTagExists(out, "urlset", "url", "lastmod")
	assert.True(t, lastModExists, "Should have default last modified")
}

func TestXMLAddUrlWithLocation(t *testing.T) {
	xml := New()
	url := NewUrl()

	url.SetLocation("https://domain.com")
	xml.AddUrl(url)

	out := xml.OutputString()
	exists, _ := checkTagExists(out, "urlset", "url", "loc")
	assert.True(t, exists, "Should have urlset,nested url nodes, and a location field")
}

func TestXMLAddUrlWithLastModified(t *testing.T) {
	xml := New()
	url := NewUrl()

	now := time.Now()
	later := now.Add(3 * time.Hour)
	url.SetLastModified(later)
	xml.AddUrl(url)

	out := xml.OutputString()
	exists, _ := checkTagExists(out, "urlset", "url", "lastmod")

	assert.True(t, exists, "Should have lastmod set")

	// value, _ := getTagCharData(src)
	// fmt.Println(value)
	// fmt.Println(out)
	// t.Error("Failed")
}
