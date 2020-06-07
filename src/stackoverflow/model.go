package stackoverflow

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
