# Sitemap Format

![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)
[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/mingard/sitemap-format)
[![Go Report Card](https://goreportcard.com/badge/mingard/sitemap-format)](https://goreportcard.com/report/mingard/sitemap-format)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/e0330046dd004b99a155ad56031acd9f)](https://www.codacy.com/gh/mingard/sitemap-format/dashboard?utm_source=github.com&utm_medium=referral&utm_content=mingard/sitemap-format&utm_campaign=Badge_Grade)

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

  loc := sitemap.NewLoc()
  loc.SetLocation("https://domain.com")
  xml.Add(loc)

  out, _ := xml.OutputString()
  fmt.Println(out)
}

// Output: <?xml version="1.0" encoding="UTF-8"?><urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"><url><loc>https://domain.com</loc><lastmod>2022-11-03T11:56:00.26065Z</lastmod></url></urlset>
```
