package github

// Scheduler contains a pointer to a Scraper struct
// and a stop channel which is used to stop the
// scheduled execution of Scraper.Scrape.
type Scheduler struct {
	Scraper     *Scraper
	stopChannel chan<- bool
}
