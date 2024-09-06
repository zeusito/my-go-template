package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/zeusito/my-go-template/pkg/config"

	"github.com/zeusito/my-go-template/internal/api"
	"github.com/zeusito/my-go-template/pkg/logger"
	"github.com/zeusito/my-go-template/pkg/router"
)

func main() {
	// Setup logger
	logger.MustConfigure()

	// Load config
	theConfs, err := config.LoadConfigurations()
	if err != nil {
		log.Fatalf("Error loading configurations: %s", err.Error())
	}

	// Init router
	theRouter := router.NewHTTPRouter(theConfs.Server)

	// Controllers
	_ = api.NewHealthController(theRouter.Mux)

	// Start server in background
	go theRouter.Start()

	// Graceful shutdown
	gracefulShutdown(theRouter)
}

func gracefulShutdown(myRouter *router.HTTPRouter) {
	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	// Signal acquired, starting to shut down all systems
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	myRouter.Shutdown(ctx)
}
