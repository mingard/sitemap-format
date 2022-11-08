package sitemap

// sitemap formatting with syntactic sugar. © Arthur Mingard 2022
import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/antchfx/xmlquery"
	"github.com/stretchr/testify/assert"
)

const (
	testDefaultXML             string = `<?xml version="1.0" encoding="UTF-8"?><urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"></urlset>`
	testDefaultSitemapIndexXML string = `<?xml version="1.0" encoding="UTF-8"?><sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"></sitemapindex>`
	testDefaultPrettyXML       string = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\">\nXXXXX<url>\nXXXXXXXXXX<lastmod>2021-08-15T14:30:45.0000001Z</lastmod>\nXXXXX</url>\n</urlset>"
	testFullNewsXML            string = `<?xml version="1.0" encoding="UTF-8"?><urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:news="http://www.google.com/schemas/sitemap-news/0.9"><url><lastmod>2021-08-15T14:30:45.0000001Z</lastmod><news:news><news:name>SOMENAME</news:name><news:publication><news:language>en</news:language></news:publication><news:publication_date>2021-08-15T14:30:45.0000001Z</news:publication_date><news:title>TITLE</news:title></news:news></url></urlset>`
	testFullImageXML           string = `<?xml version="1.0" encoding="UTF-8"?><urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:image="http://www.google.com/schemas/sitemap-image/1.1"><url><loc>https://somedomain.com/article</loc><lastmod>2021-08-15T14:30:45.0000001Z</lastmod><image:image><image:loc>https://somedomain.com/image.jpg</image:loc></image:image></url></urlset>`
)

var fixedTime = time.Date(2021, 8, 15, 14, 30, 45, 100, time.UTC)

func getDoc(str string) (*xmlquery.Node, error) {
	return xmlquery.Parse(strings.NewReader(str))
}

func findOne(str, tags string) (*xmlquery.Node, error) {
	doc, err := getDoc(str)
	if err != nil {
		return nil, err
	}

	tag := xmlquery.FindOne(doc, tags)
	if tag != nil {
		return tag, nil
	}

	return nil, fmt.Errorf(`Failed to find %s`, tags)
}

func TestDefaults(t *testing.T) {
	xml := New()

	out, _ := xml.OutputString()

	assert.Equal(t, testDefaultXML, out, "Should output default empty XML")
}

func TestNewSitemapIndex(t *testing.T) {
	xml := NewSitemapIndex()

	out, _ := xml.OutputString()

	assert.Equal(t, testDefaultSitemapIndexXML, out, "Should output default empty XML")
}

func TestNewSitemapIndexAdd(t *testing.T) {
	xml := NewSitemapIndex()
	loc := NewLoc()
	xml.Add(loc)

	assert.Equal(t, "sitemap", xml.parent.Locations[0].XMLName.Local, "Should have an XML node local of 'sitemap")
}

func TestPrettyOutput(t *testing.T) {
	xml := New()

	loc := NewLoc()
	loc.SetLastModified(fixedTime)
	xml.Add(loc)

	out, _ := xml.OutputPrettyString("", "XXXXX")

	assert.Equal(t, testDefaultPrettyXML, out, "Should output prettified XML")
}

func TestXMLAddDefaultUrl(t *testing.T) {
	xml := New()
	loc := NewLoc()
	xml.Add(loc)
	out, _ := xml.OutputString()

	tag, _ := findOne(out, "//urlset/url/loc")
	assert.True(t, tag == nil, "Should have urlset, nested url nodes but without location")
	lastModTag, _ := findOne(out, "//urlset/url/lastmod")
	assert.True(t, lastModTag != nil, "Should have default last modified")
}

func TestXMLAddEntryWithLocation(t *testing.T) {
	xml := New()
	loc := NewLoc()

	loc.SetLocation("https://domain.com")
	xml.Add(loc)

	out, _ := xml.OutputString()
	tag, _ := findOne(out, "//urlset/url/loc")
	assert.True(t, tag != nil, "Should have urlset,nested url nodes, and a location field")
}

func TestDefaultType(t *testing.T) {
	xml := New()

	out, _ := xml.OutputString()
	doc, _ := findOne(out, "//urlset/@xmlns")
	assert.Equal(t, doc.InnerText(), XMLNSDefault.Value, "Should contain correct default type")
}

func TestSetNews(t *testing.T) {
	xml := New()
	loc := NewLoc()
	news := NewNews()
	loc.SetNews(news)
	xml.Add(loc)

	out, _ := xml.OutputString()
	doc, _ := findOne(out, "//urlset/@xmlns:news")
	assert.Equal(t, doc.InnerText(), XMLNSNews.Value, "Should contain correct XMLNS attr")
}

func TestSetNewsValues(t *testing.T) {

	xml := New()
	loc := NewLoc()
	loc.SetLastModified(fixedTime)
	news := NewNews()
	news.SetTitle("TITLE")
	news.SetName("SOMENAME")
	news.SetPublicationDate(fixedTime)
	news.SetLanguage(LanguageEN)
	loc.SetNews(news)
	xml.Add(loc)

	out, _ := xml.OutputString()
	assert.Equal(t, testFullNewsXML, out, "Should output news XML")
}

func TestAddImage(t *testing.T) {
	xml := New()
	loc := NewLoc()
	image := NewImage()
	loc.AddImage(image)
	xml.Add(loc)

	out, _ := xml.OutputString()
	doc, _ := findOne(out, "//urlset/@xmlns:image")
	assert.Equal(t, doc.InnerText(), XMLNSImage.Value, "Should contain correct XMLNS attr")
}

func TestAddImageValues(t *testing.T) {

	xml := New()
	loc := NewLoc()
	loc.SetLocation("https://somedomain.com/article")
	loc.SetLastModified(fixedTime)

	image := NewImage()
	image.SetLocation("https://somedomain.com/image.jpg")
	loc.AddImage(image)
	xml.Add(loc)

	out, _ := xml.OutputString()
	assert.Equal(t, testFullImageXML, out, "Should output image XML")
}

func TestAddVideo(t *testing.T) {
	xml := New()
	loc := NewLoc()
	video := NewVideo()
	loc.AddVideo(video)
	xml.Add(loc)

	out, _ := xml.OutputString()
	doc, _ := findOne(out, "//urlset/@xmlns:video")
	assert.Equal(t, doc.InnerText(), XMLNSVideo.Value, "Should contain correct XMLNS attr")
}

func TestAddVideoValues(t *testing.T) {

	xml := New()
	loc := NewLoc()
	loc.SetLocation("https://somedomain.com/article")
	loc.SetLastModified(fixedTime)

	video := NewVideo()
	video.SetTitle("Some title")
	video.SetThumbnailLocation("https://somedomain.com/thumb.jpg")
	video.SetContentLocation("https://somedomain.com/video.mp4")
	video.SetPlayerLocation("https://somedomain.com/player")
	video.SetDuration(100)
	video.SetExpirationDate(fixedTime)
	video.SetRating(2)
	video.SetViewCount(100)
	video.SetPublicationDate(fixedTime)
	loc.AddVideo(video)
	xml.Add(loc)

	out, _ := xml.OutputString()
	assert.Equal(t, testFullImageXML, out, "Should output video XML")
}

func TestXMLAddEntryWithLastModified(t *testing.T) {
	xml := New()
	loc := NewLoc()

	now := time.Now()
	later := now.Add(3 * time.Hour)
	loc.SetLastModified(later)
	xml.Add(loc)

	out, _ := xml.OutputString()

	tag, _ := findOne(out, "//urlset/url/lastmod")
	assert.True(t, tag != nil, "Should have lastmod set")

	nodeTime, _ := time.Parse(time.RFC3339, tag.InnerText())
	assert.Equal(t, nodeTime, later.UTC(), "Should match input timestamp")
}
