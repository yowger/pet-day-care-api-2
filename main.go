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
	"github.com/yowger/pet-day-care-api-2/router"
)

func main() {
	cfg := config.LoadAppConfig()
	e := echo.New()
	db := database.NewDatabase(cfg)
	queries := sqlc.New(db)
	ctx := context.Background()

	router.NewRouter(e, queries, ctx)

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

	waitGrp.Add(1)
	go func() {
		defer waitGrp.Done()

		for {
			select {
			case <-notifyCtx.Done():
				return
			case <-time.After(30 * time.Second):
				if err := db.Ping(context.Background()); err != nil {
					log.Fatalf("Error connecting to database: %v", err)
				}
			}
		}
	}()

	<-notifyCtx.Done()

	log.Println("Shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Error during server shutdown: %v", err)
	}

	waitGrp.Wait()
	db.Close()

	log.Println("Shutdown complete...")
}
