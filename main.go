package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/yowger/pet-day-care-api-2/config"
	"github.com/yowger/pet-day-care-api-2/database"
	"github.com/yowger/pet-day-care-api-2/server"
)

func main() {
	configPath := "."
	configName := ".env"
	conf := config.LoadAppConfig(configPath, configName)

	db, err := database.NewDatabase(conf.DATABASE_URL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	server := server.NewServer(conf.PORT)

	waitGrp := &sync.WaitGroup{}

	waitGrp.Add(1)
	go func() {
		defer waitGrp.Done()

		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	notifyCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	healthCheckInterval := 30 * time.Second
	waitGrp.Add(1)
	go func() {
		defer waitGrp.Done()

		for {
			select {
			case <-notifyCtx.Done():
				return
			case <-time.After(healthCheckInterval):
				if err := db.Ping(context.Background()); err != nil {
					log.Fatalf("Error connecting to database: %v", err)
				}
			}
		}
	}()

	<-notifyCtx.Done()

	log.Println("Shutting down server...")

	shutdownDelay := 10 * time.Second
	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownDelay)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Error during server shutdown: %v", err)
	}

	waitGrp.Wait()
	db.Close()

	log.Println("Shutdown complete...")
}
