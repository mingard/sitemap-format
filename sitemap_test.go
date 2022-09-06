package sitemap

import (
	"fmt"
	"testing"
)

func TestXML(t *testing.T) {
	xml := New()
	url := NewUrl()
	url.SetLocation("https://domain.com")
	xml.AddUrl(url)
	out, err := xml.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
	t.Errorf("Failed")
}
