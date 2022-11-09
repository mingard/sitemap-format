package sitemap

import (
	"encoding/xml"
	"time"
)

const (
	shortFormat string = "2006-01-02"
	fullFormat  string = "2006-01-02T15:04:05-07:00"
)

// customDate is a time.Time value with custom XML unmarshal method.
type customDate struct {
	time.Time
	format string
}

// MarshalXML is a customer marshal method for customDate.
func (c *customDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	t := time.Time(c.Time)
	v := t.Format(c.format)
	return e.EncodeElement(v, start)
}

// Date returns a new short format date.
func Date(t time.Time) customDate {
	return customDate{t, shortFormat}
}

// DateFull returns a long format date.
func DateFull(t time.Time) customDate {
	return customDate{t, fullFormat}
}
