package jobsbg

// Scheduler executes Service.RunForActiveQueries at a given time period.Scheduler
// This scheduled execution can be also be started/stopped.
type Scheduler struct {
	Scraper     *Scraper
	stopChannel chan<- bool
}
