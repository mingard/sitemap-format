package sitemap

import (
	"encoding/xml"
	"fmt"
	"time"
)

const (
	shortFormat   string = "YYYY-MM-DD"
	defaultFormat string = "YYYY-MM-DDThh:mmTZD"
)

// customDate is a time.Time value with custom XML unmarshal method.
type customDate struct {
	time.Time
	shortFormat bool
}

// useShortFormat sets the date format to short.
func (c *customDate) useShortFormat() {
	c.shortFormat = true
}

// UnmarshalXML is a customer unmarshal method for customDate.
func (c *customDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string

	format := defaultFormat
	if c.shortFormat {
		format = shortFormat
	}

	fmt.Println("ASD", format)

	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	parse, err := time.Parse(format, v)
	if err != nil {
		return err
	}
	*c = customDate{parse, c.shortFormat}
	return nil
}
