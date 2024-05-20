package scheduler

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"sync"
	"time"
)

type Config struct {
	doNothingJobInterval int `koanf:"do_nothing_job_interval"`
}

type Scheduler struct {
	sch    *gocron.Scheduler
	config Config
}

func New(config Config) Scheduler {
	return Scheduler{
		config: config,
		sch:    gocron.NewScheduler(time.UTC)}
}

func (s Scheduler) Start(done <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	_, err := s.sch.Every(s.config.doNothingJobInterval).Second().Do(s.doNothing)
	if err != nil {
		panic(err)
	}

	s.sch.StartAsync()

	<-done
	// wait to finish job
	fmt.Println("stop scheduler..")
	s.sch.Stop()
}
