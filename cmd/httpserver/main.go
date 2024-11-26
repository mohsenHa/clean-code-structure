package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"clean-code-structure/config"
	"clean-code-structure/delivery/httpserver"
	"clean-code-structure/logger"
	"clean-code-structure/service/healthservice"
	"clean-code-structure/validator/healthvalidator"
)

func main() {
	wg := &sync.WaitGroup{}
	done := make(chan bool)

	cfg, err := config.Load("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("cfg: %+v\n", cfg)

	logger.Start(cfg.Logger)

	rSvcs := setupServices(cfg, wg, done)

	server := httpserver.New(cfg, rSvcs)
	go func() {
		server.Serve()
	}()

	if cfg.Application.EnableProfiling {
		profiling(cfg, wg, done)
	}

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	<-exit
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
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Application.ProfilingPort),
		ReadHeaderTimeout: time.Duration(cfg.Application.ReadHeaderTimeout) * time.Second,
	}
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// unexpected error. port in use?
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	go func() {
		<-done
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Application.GracefulShutdownTimeoutInSecond))
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			panic(err)
		}
	}()
}

func setupServices(_ config.Config, _ *sync.WaitGroup, _ chan bool) (requiredServices httpserver.RequiredServices) {
	healthValidator := healthvalidator.New()

	requiredServices.HealthService = healthservice.New(healthValidator)

	return requiredServices
}
