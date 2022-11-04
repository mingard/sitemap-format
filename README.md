# Sitemap Format
![Coverage](https://img.shields.io/badge/Coverage-33.6%25-yellow)

[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/mingard/sitemap-format)
[![Go Report Card](https://goreportcard.com/badge/mingard/sitemap-format)](https://goreportcard.com/report/mingard/sitemap-format)
[![Coverage](https://gocover.io/_badge/github.com/mingard/sitemap-format)](https://gocover.io/github.com/mingard/sitemap-format)

Simple sitemap formatting, with a little syntactic sugar.

## Installation

```sh
go get github.com/mingard/sitemap-format
```

## Usage

### Creating a basic Sitemap

```go
package main

import sitemap "github.com/mingard/sitemap-format"

func main() {
  xml := sitemap.New()

  url := sitemap.NewUrl()
  url.SetLocation("https://domain.com")
  xml.AddEntry(url)

  out, _ := xml.OutputString()
  fmt.Println(out)
}

// Output: <?xml version="1.0" encoding="UTF-8"?><urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"><url><loc>https://domain.com</loc><lastmod>2022-11-03T11:56:00.26065Z</lastmod></url></urlset>
```
