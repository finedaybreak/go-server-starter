package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"go-server-starter/internal/app"
	"go-server-starter/internal/config"
	"go-server-starter/pkg/logger"
)

func main() {
	serverMode, err := config.ParseMode()
	if err != nil {
		panic(err)
	}

	vc, err := config.NewViperConfig(serverMode)
	if err != nil {
		panic(err)
	}

	logger := logger.NewLogger(&vc.GetConfig().Logger, *serverMode)
	defer logger.Sync()

	logger.Info(fmt.Sprintf("Starting server in %s mode", serverMode.String()))

	app := app.NewApp(vc.GetConfig(), logger)
	if err := app.Start(); err != nil {
		logger.Fatal(err.Error())
	}

	// Setup signal handling for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	sig := <-quit
	logger.Info(fmt.Sprintf("Received signal: %v", sig))

	// Gracefully shutdown the server
	if err := app.Shutdown(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
