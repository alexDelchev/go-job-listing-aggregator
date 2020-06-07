package weworkremotely

// Module contains a Scraper, Scheduler, and a controller,
// with the Scraper and Scheduler being exported.
type Module struct {
	Scraper    Scraper
	Scheduler  Scheduler
	controller controller
}
