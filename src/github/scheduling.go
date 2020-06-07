package github

import (
	"log"
	"time"
)

// Scheduler contains a pointer to a Scraper struct
// and a stop channel which is used to stop the
// scheduled execution of Scraper.Scrape.
type Scheduler struct {
	Scraper     *Scraper
	stopChannel chan<- bool
}

// NewScheduler creates a new instance and starts the
// scheduler.
func NewScheduler(Scraper *Scraper) Scheduler {
	scheduler := Scheduler{Scraper: Scraper}
	scheduler.Start()
	return scheduler
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

// Start is used to start the scheduler. If currently
// another schedule is running, it is stopped before
// starting the new one.
func (s *Scheduler) Start() {
	log.Println("Starting GitHub Scheduler")
	action := s.Scraper.RunForActiveQueries
	interval := 1 * time.Minute

	channel := schedule(action, interval)

	if s.stopChannel != nil {
		s.stopChannel <- true
		close(s.stopChannel)
	}
	s.stopChannel = channel
}

// Stop is used to stop the scheduler.
func (s *Scheduler) Stop() {
	if s.stopChannel != nil {
		log.Println("Stopping GitHub Scheduler")
		s.stopChannel <- true
		close(s.stopChannel)
		s.stopChannel = nil
	}
}
