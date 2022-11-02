package sitemap

// See https://developers.google.com/search/docs/crawling-indexing/sitemaps/video-sitemaps

import (
	"encoding/xml"
	"strings"
	"time"
)

// Restriction stores country restriction details
type Restriction struct {
	XMLName      xml.Name  `xml:"video:restriction,omitempty"`
	Relationship *xml.Attr `xml:",attr,omitempty"`
	Value        string    `xml:",chardata"`
}

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
	FamilyFriendly BoolStr        `xml:"video:family_friendly,omitempty"`
	Restrictions   []*Restriction `xml:"video:restriction,omitempty"`
	Platforms      []*Platform    `xml:"video:platform,omitempty"`
	// Live indicates whether the video is a live stream
	Live BoolStr `xml:"video:live,omitempty"`
}

// addPlatform creates a platform restruction block.
func (v *Video) addPlatform(r *xml.Attr, p ...PlatformName) {
	// Convert platforms to string
	platforms := make([]string, 0)

	for _, a := range p {
		platforms = append(platforms, string(a))
	}

	v.Platforms = append(v.Platforms, &Platform{
		Relationship: r,
		Value:        strings.Join(platforms, " "),
	})
}

// AllowPlatforms creates a list of allowed platforms.
func (v *Video) AllowPlatforms(p ...PlatformName) *Video {
	v.addPlatform(Allow, p...)
	return v
}

// DenyPlatforms creates a list of denied platforms.
func (v *Video) DenyPlatforms(p ...PlatformName) *Video {
	v.addPlatform(Deny, p...)
	return v
}

// setRestrictions creates a country restriction block.
func (v *Video) setRestrictions(r *xml.Attr, c string) {
	v.Restrictions = append(v.Restrictions, &Restriction{
		Relationship: r,
		Value:        c,
	})
}

// AllowCountries creates a list of allowed countries.
func (v *Video) AllowCountries(c string) *Video {
	v.setRestrictions(Allow, c)
	return v
}

// DenyCountries creates a list of denied countries.
func (v *Video) DenyCountries(c string) *Video {
	v.setRestrictions(Deny, c)
	return v
}

// IsLive sets the live option to 'yes'
func (v *Video) IsLive() *Video {
	v.Live = Yes
	return v
}

// NotLive sets the live option to 'no'
func (v *Video) NotLive() *Video {
	v.Live = No
	return v
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

func defaultVideo() *Video {
	return &Video{
		Restrictions: make([]*Restriction, 0),
		Platforms:    make([]*Platform, 0),
	}
}

// NewVideo returns a new instance of the default Video extension.
func NewVideo() *Video {
	return defaultVideo()
}
