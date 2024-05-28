package main

import (
	"clean-code-structure/config"
	"clean-code-structure/delivery/httpserver"
	"clean-code-structure/logger"
	"clean-code-structure/service/healthservice"
	"clean-code-structure/validator/healthvalidator"
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	done := make(chan bool)

	cfg := config.Load("config.yml")
	fmt.Printf("cfg: %+v\n", cfg)

	logger.Start(cfg.Logger)

	rSvcs, rVal := setupServices(cfg, wg, done)

	server := httpserver.New(cfg, rSvcs, rVal)
	go func() {
		server.Serve()
	}()

	if cfg.Application.EnableProfiling {
		profiling(cfg, wg, done)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("Quit signal received")
	close(done)

	fmt.Println("Wait for done all processes")
	wg.Wait()

	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx,
		time.Duration(cfg.Application.GracefulShutdownTimeoutInSecond)*time.Second)
	defer cancel()

	fmt.Println("Shutting down server router")
	if err := server.Router.Shutdown(ctxWithTimeout); err != nil {
		fmt.Println("http server shutdown error", err)
	}

	fmt.Println("received interrupt signal, shutting down gracefully..")
	<-ctxWithTimeout.Done()
}

func profiling(cfg config.Config, wg *sync.WaitGroup, done <-chan bool) {
	fmt.Printf("Profiling enabled on port %d", cfg.Application.ProfilingPort)
	srv := &http.Server{Addr: fmt.Sprintf(":%d", cfg.Application.ProfilingPort)}
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// unexpected error. port in use?
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	go func() {
		for {
			select {
			case <-done:
				ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
				defer cancel()
				if err := srv.Shutdown(ctx); err != nil {
					panic(err)
				}
			}
		}
	}()
}
func setupServices(cfg config.Config, wg *sync.WaitGroup, done chan bool) (requiredServices httpserver.RequiredServices, requiredValidators httpserver.RequiredValidators) {

	requiredValidators.HealthValidator = healthvalidator.New()

	requiredServices.HealthService = healthservice.New()

	return
}
