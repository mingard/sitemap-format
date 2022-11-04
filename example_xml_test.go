package sitemap

// sitemap formatting with syntactic sugar. © Arthur Mingard 2022

import "fmt"

func ExampleNew() {
	xml := New()

	out, _ := xml.Output()
	fmt.Println(string(out))
	// Output: <?xml version="1.0" encoding="UTF-8"?><urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"></urlset>
}

func ExampleXML_OutputString() {
	xml := New()

	out, _ := xml.OutputString()
	fmt.Println(out)
	// Output: <?xml version="1.0" encoding="UTF-8"?><urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"></urlset>
}

func ExampleXML_OutputPrettyString() {
	xml := New()

	out, _ := xml.OutputPrettyString("", "  ")
	fmt.Println(out)
	// Output: <?xml version="1.0" encoding="UTF-8"?>
	//<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"></urlset>
}
