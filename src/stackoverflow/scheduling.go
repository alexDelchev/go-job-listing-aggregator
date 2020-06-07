package stackoverflow

// Scheduler calls Scraper.RunForActiveQueries at a
// given interval.
type Scheduler struct {
	Scraper     *Scraper
	stopChannel chan<- bool
}
