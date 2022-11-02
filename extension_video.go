package sitemap

// See https://developers.google.com/search/docs/crawling-indexing/sitemaps/video-sitemaps

import (
	"encoding/xml"
	"time"
)

type BoolStr string

const (
	Yes BoolStr = "yes"
	No  BoolStr = "no"
)

// Video stores video entry data.
type Video struct {
	XMLName xml.Name `xml:"video:video"`

	// Required
	// Title HTML entities must be escaped or wrapped in a CDATA block.
	Title string `xml:"video:title"`
	// Description Maximum 2048 characters. All HTML entities must be escaped or wrapped in a CDATA block. It must match the description displayed on the web page (it doesn't need to be a word-for-word match).
	Description  string `xml:"video:description,omitempty"`
	ThumbnailLoc string `xml:"video:thumbnail_loc,omitempty"`
	// ContentLoc HTML and Flash aren't supported formats.
	ContentLoc string `xml:"video:content_loc,omitempty"`
	// Can be used instead of or alongside ContentLoc
	PlayerLoc string `xml:"video:player_loc,omitempty"`

	// Recommended
	// Duration value must be from 1 to 28800 (8 hours) inclusive
	Duration int `xml:"video:duration,omitempty"`
	// ExpirationDate format either YYYY-MM-DD or YYYY-MM-DDThh:mm:ss+TZD
	ExpirationDate time.Time `xml:"video:expiration_date,omitempty"`

	// Optional
	// Rating values are float numbers in the range 0.0 (low) to 5.0 (high), inclusive
	Rating          float32   `xml:"video:rating,omitempty"`
	ViewCount       int       `xml:"video:view_count,omitempty"`
	PublicationDate time.Time `xml:"video:publication_date,omitempty"`
	// FamilyFriendly whether the video is available with SafeSearch.
	FamilyFriendly BoolStr `xml:"video:family_friendly,omitempty"`
}

// IsFamilyFriendly sets the family friendly option to 'yes'
func (v *Video) IsFamilyFriendly() *Video {
	v.FamilyFriendly = Yes
	return v
}

// NotFamilyFriendly sets the family friendly option to 'no'
func (v *Video) NotFamilyFriendly() *Video {
	v.FamilyFriendly = No
	return v
}

// NewVideo returns a new instance of the default Video extension.
func NewVideo() *Video {
	return new(Video)
}
