package weworkremotely

// Scheduler calls Service.RunForActiveQueries at a given
// interval.
type Scheduler struct {
	Scraper     *Scraper
	stopChannel chan<- bool
}
