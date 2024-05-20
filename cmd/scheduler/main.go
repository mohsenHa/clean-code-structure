package main

import (
	"clean-code-structure/config"
	"clean-code-structure/scheduler"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	cfg := config.Load("config.yml")
	fmt.Printf("cfg: %+v\n", cfg)

	done := make(chan bool)
	wg := &sync.WaitGroup{}

	go func() {
		sch := scheduler.New(cfg.Scheduler)
		sch.Start(done, wg)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("received interrupt signal, shutting down gracefully..")
	done <- true
	time.Sleep(time.Second * time.Duration(cfg.Application.GracefulShutdownTimeoutInSecond))
}
