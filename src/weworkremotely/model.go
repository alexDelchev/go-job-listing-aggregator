package weworkremotely

type jobListingRSSModel struct {
	ID          string `xml:"guid"`
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"PubDate"`
}
