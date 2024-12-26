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

	"github.com/labstack/echo/v4"
	"github.com/yowger/pet-day-care-api-2/config"
	"github.com/yowger/pet-day-care-api-2/database"
	sqlc "github.com/yowger/pet-day-care-api-2/database/sqlc"
)

func main() {
	e := echo.New()
	cfg := config.LoadAppConfig()
	db := database.NewDatabase(cfg)
	queries := sqlc.New(db)
	ctx := context.Background()

	waitGrp := &sync.WaitGroup{}

	waitGrp.Add(1)
	go func() {
		defer waitGrp.Done()

		if err := e.Start(cfg.PORT); err != nil && err != http.ErrServerClosed {
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

	if err := e.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Error during server shutdown: %v", err)
	}

	waitGrp.Wait()
	db.Close()

	log.Println("Shutdown complete...")
}
