package github

import (
	"time"
)

// Scheduler contains a pointer to a Scraper struct
// and a stop channel which is used to stop the
// scheduled execution of Scraper.Scrape.
type Scheduler struct {
	Scraper     *Scraper
	stopChannel chan<- bool
}

type task func()

func schedule(action task, interval time.Duration) chan<- bool {
	ticker := time.NewTicker(interval)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				action()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	return quit
}
