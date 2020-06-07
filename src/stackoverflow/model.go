package stackoverflow

import "encoding/xml"

type jobListingRSSModel struct {
	ID             string   `xml:"guid"`
	Link           string   `xml:"link"`
	AuthorName     string   `xml:"author>name"`
	Categories     []string `xml:"category"`
	Title          string   `xml:"title"`
	Description    string   `xml:"description"`
	PublishingDate string   `xml:"pubDate"`
	UpdateDate     string   `xml:"updated"`
}

type jobsRSSChannel struct {
	Title            string               `xml:"title"`
	Link             string               `xml:"link"`
	Description      string               `xml:"decription"`
	PositionListings []jobListingRSSModel `xml:"item"`
}

type jobsRSSFeed struct {
	XMLName xml.Name       `xml:"rss"`
	Channel jobsRSSChannel `xml:"channel"`
}
