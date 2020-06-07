package jobsbg

import (
	"log"
	"time"
)

// Scheduler executes Service.RunForActiveQueries at a given time period.Scheduler
// This scheduled execution can be also be started/stopped.
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

// Start starts the scheduled execution of Service.RunForActiveQueries,
// using a time interval ot 1 minute.
func (s *Scheduler) Start() {
	log.Println("Starting Jobs.bg Scheduler")
	action := s.Scraper.RunForActiveQueries
	interval := 1 * time.Minute

	channel := schedule(action, interval)

	if s.stopChannel != nil {
		s.stopChannel <- true
	}
	s.stopChannel = channel
}
