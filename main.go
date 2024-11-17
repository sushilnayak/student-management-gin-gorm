package main

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"io"
	"log"
	"student-service/internal/api"
	"student-service/internal/config"
	"student-service/internal/db"
	"student-service/internal/middleware"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger
	logger := middleware.NewLogger(cfg)
	defer logger.Sync()

	// Initialize tracer
	var tracer opentracing.Tracer
	var closer io.Closer
	tracer, closer, err = middleware.InitTracer(cfg)
	if err != nil {
		logger.Fatal("Failed to initialize tracer")
	}
	defer closer.Close()

	// Initialize database
	database, err := db.NewDatabase(cfg)
	if err != nil {
		logger.Fatal("Failed to connect to database")
	}

	// Initialize server
	//server := api.NewServer(cfg, database, logger)
	server := api.NewServer(cfg, database, logger, tracer)
	if err := server.Start(context.Background()); err != nil {
		logger.Fatal("Failed to start server")
	}
}
