package sitemap

// See https://developers.google.com/search/docs/crawling-indexing/sitemaps/video-sitemaps

import (
	"encoding/xml"
	"strings"
	"time"
)

const (
	// TagLimit is the number of permitted video tags.
	TagLimit = 32
	// MaxDescriptionLength is the maximum length of the video:description.
	MaxDescriptionLength = 2048
	// DurationMin is the minimum video duration of 1 second.
	DurationMin int = 1
	// DurationMax is the maximum video duration length in seconds (8 hours).
	DurationMax int = 28800
	// RatingLow is the lowest posible rating.
	RatingLow float32 = 0.0
	// RatingHigh is the highest posible rating.
	RatingHigh float32 = 5.0
)

// Restriction stores country restriction details
type Restriction struct {
	XMLName      xml.Name  `xml:"video:restriction,omitempty"`
	Relationship *xml.Attr `xml:",attr,omitempty"`
	Value        string    `xml:",chardata"`
}

// Uploader stores the uploader of the video.
type Uploader struct {
	XMLName xml.Name `xml:"video:uploader,omitempty"`
	// Info (optional) specifies the URL of a webpage with additional information about this uploader.
	Info string `xml:"info,attr,omitempty"`
	// Value is the video uploader's name, a string with a maximum of 255 characters.
	Value string `xml:",chardata"`
}

// Video stores video entry data.
type Video struct {
	XMLName xml.Name `xml:"video:video"`

	// Required
	
	// Title HTML entities must be escaped or wrapped in a CDATA block.
	Title string `xml:"video:title"`
	// Description Maximum {{MaxDescriptionLength}} characters. All HTML entities must be escaped or wrapped in a CDATA block. It must match the description displayed on the web page (it doesn't need to be a word-for-word match).
	Description  string `xml:"video:description,omitempty"`
	ThumbnailLoc string `xml:"video:thumbnail_loc,omitempty"`
	// ContentLoc HTML and Flash aren't supported formats.
	ContentLoc string `xml:"video:content_loc,omitempty"`
	// Can be used instead of or alongside ContentLoc.
	PlayerLoc string `xml:"video:player_loc,omitempty"`
	
	// Recommended
	
	// Duration value must be from {{DurationMin}} to {{DurationMax}} inclusive.
	Duration int `xml:"video:duration,omitempty"`
	// ExpirationDate format either YYYY-MM-DD or YYYY-MM-DDThh:mm:ss+TZD
	ExpirationDate time.Time `xml:"video:expiration_date,omitempty"`
	
	// Optional
	
	// Rating values are float numbers in the range {{RatingLow}} (low) to {{RatingHigh}} (high), inclusive.
	Rating          float32   `xml:"video:rating,omitempty"`
	ViewCount       int       `xml:"video:view_count,omitempty"`
	PublicationDate time.Time `xml:"video:publication_date,omitempty"`
	// FamilyFriendly whether the video is available with SafeSearch.
	FamilyFriendly BoolStr        `xml:"video:family_friendly,omitempty"`
	Restrictions   []*Restriction `xml:"video:restriction,omitempty"`
	Platforms      []*Platform    `xml:"video:platform,omitempty"`
	// RequiresSubscription indicates whether a subscription is required to view the video.
	RequiresSubscription BoolStr   `xml:"video:requires_subscription,omitempty"`
	Uploader             *Uploader `xml:"video:uploader,omitempty"`
	// Live indicates whether the video is a live stream
	Live BoolStr `xml:"video:live,omitempty"`
	// Tags are limited to a max of {{TagLimit}}.
	Tags []string `xml:"video:tag,omitempty"`
}

// SetTitle sets the video extensions title parameter.
func (v *Video) SetTitle(t string) *Video {
	v.Title = t
	return v
}

// SetDescription sets the video extensions description parameter.
func (v *Video) SetDescription(d string) *Video {
	// Limit to the maximum length of a description.
	v.Description = d[:MaxDescriptionLength]
	return v
}

// SetThumbnailLoc sets the video thumbnail location parameter.
func (v *Video) SetThumbnailLoc(t string) *Video {
	v.ThumbnailLoc = t
	return v
}

// SetContentLoc sets the video content location parameter.
func (v *Video) SetContentLoc(c string) *Video {
	v.ContentLoc = c
	return v
}

// SetPlayerLoc sets the video player location parameter.
func (v *Video) SetPlayerLoc(p string) *Video {
	v.PlayerLoc = p
	return v
}

// SetDuration sets the video duration parameter.
func (v *Video) SetDuration(d int) *Video {
	// Must be no less than min and no more than max.
	if d >= DurationMin && d <= DurationMax {
		v.Duration = d
	}
	return v
}

// SetExpirationDate sets the video ExpirationDate parameter.
func (v *Video) SetExpirationDate(t time.Time) *Video {
	v.ExpirationDate = t.UTC()
	return v
}

// SetRating sets the video rating.
func (v *Video) SetRating(r float32) *Video {
	if r >= RatingLow && r <= RatingHigh {
		v.Rating = r
	}
	return v
}

// SetViewCount sets the video view_count.
func (v *Video) SetViewCount(vc int) *Video {
	v.ViewCount = vc
	return v
}

// SetPublicationDate sets the video extensions PublicationDate parameter.
func (v *Video) SetPublicationDate(t time.Time) *Video {
	v.PublicationDate = t.UTC()
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

// SubRequired sets the requires_subscription option to 'yes'
func (v *Video) SubRequired() *Video {
	v.RequiresSubscription = Yes
	return v
}

// SubNotRequired sets the requires_subscription option to 'no'
func (v *Video) SubNotRequired() *Video {
	v.RequiresSubscription = No
	return v
}

// SetUploader sets the uploader
func (v *Video) SetUploader(u *Uploader) *Video {
	v.FamilyFriendly = No
	return v
}

// SetUploaderInfo sets the uploader
func (v *Video) SetUploaderInfo(i string) *Video {
	if v.Uploader == nil {
		return v.SetUploader(&Uploader{
			Info: i,
		})
	}
	v.Uploader.Info = i
	return v
}

// SetUploaderVal sets the uploader value
func (v *Video) SetUploaderVal(val string) *Video {
	if v.Uploader == nil {
		return v.SetUploader(&Uploader{
			Value: val,
		})
	}
	v.Uploader.Value = val
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

// SetTags sets tag values
func (v *Video) SetTags(t ...string) *Video {
	// Only push tags if we're below the limit.
	if len(v.Tags)+len(t) <= cap(v.Tags) {
		v.Tags = append(v.Tags, t...)
	}
	return v
}

func defaultVideo() *Video {
	return &Video{
		Restrictions: make([]*Restriction, 0),
		Platforms:    make([]*Platform, 0),
		Tags:         make([]string, 0, TagLimit),
	}
}

// NewVideo returns a new instance of the default Video extension.
func NewVideo() *Video {
	return defaultVideo()
}
