package main

import (
	"context"
	"fmt"
	"github.com/SunFlowers04/SmartWayTT/config"
	"github.com/SunFlowers04/SmartWayTT/internal/ports"
	"github.com/SunFlowers04/SmartWayTT/internal/ports/handler"
	repo "github.com/SunFlowers04/SmartWayTT/internal/repository"
	"github.com/SunFlowers04/SmartWayTT/migrations"
	"github.com/SunFlowers04/SmartWayTT/pkg/httpserver"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Repository
	dsn := cfg.PGConnectionURLString()
	pg, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Postgres connection error: %s", err)
		return
	}

	defer pg.Close(ctx)

	// Init Migrations
	err = migrations.InitMigrate(pg)
	if err != nil {
		log.Fatalf("Migrations error: %s", err)
		return
	}

	// repositories
	flightRepo := repo.NewStorageFlight(pg)
	passengerRepo := repo.NewStoragePass(pg)
	documentRepo := repo.NewStorageDocument(pg)

	// handlers
	flightHandler := handler.NewFlightHandler(flightRepo)
	passengerHandler := handler.NewPassengerHandler(passengerRepo)
	documentHandler := handler.NewDocumentHandler(documentRepo)

	// HTTP Server
	engine := gin.New()
	ports.Router(engine, flightHandler, passengerHandler, documentHandler)
	httpServer := httpserver.New(engine, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		fmt.Println("app - Run - signal: " + s.String())
		// Shutdown
		err = httpServer.Shutdown()
		if err != nil {
			fmt.Printf("app - Run - httpServer.Shutdown: %s\r\n", err.Error())
		}
		fmt.Println("app - Run - Shutdown")
	case err = <-httpServer.Notify():

		// Shutdown
		err = httpServer.Shutdown()
		if err != nil {
			fmt.Printf("app - Run - httpServer.Shutdown: %s\r\n", err.Error())
		}
	}
}
