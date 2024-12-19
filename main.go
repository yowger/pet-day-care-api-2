package main

import (
	"context"
	"fmt"
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
)

func main() {
	e := echo.New()
	waitGroup := &sync.WaitGroup{}
	notifyCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	healthCheckInterval := 30 * time.Second

	configPath := "."
	configName := ".env"
	conf := config.LoadAppConfig(configPath, configName)

	pgxPool, err := database.ConnectDatabase(conf.DATABASE_URL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()

		port := fmt.Sprintf(":%s", conf.PORT)
		if err := e.Start(port); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()

		for {
			select {
			case <-notifyCtx.Done():
				return
			case <-time.After(healthCheckInterval):
				if err := pgxPool.Ping(context.Background()); err != nil {
					log.Fatalf("Error connecting to database: %v", err)
				}
			}
		}
	}()

	<-notifyCtx.Done()

}
