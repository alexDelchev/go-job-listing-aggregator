package github

// Module contains a scraper, scheduler and controller,
// with the Scraper and Scheduler being exported.
type Module struct {
	Scraper    Scraper
	Scheduler  Scheduler
	controller controller
}
