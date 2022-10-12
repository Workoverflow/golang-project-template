package main

import (
	"context"
	"github.com/joho/godotenv"
	"golang-project-template/internal/application/app"
	"golang-project-template/internal/application/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	var runChan = make(chan os.Signal, 1)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		config.GetConfig().Application.Server.Timeout,
	)
	defer cancel()

	// Create app instance
	app := app.Create(ctx, config.GetConfig())

	// Handle ctrl+c/ctrl+x interrupt
	signal.Notify(runChan, os.Interrupt, syscall.SIGTSTP)
	interrupt := <-runChan

	// Start app or server
	app.Start()

	log.Printf("Server is shutting down due to %+v\n", interrupt)
	if err := app.Stop(ctx); err != nil {
		log.Fatalf("Server was unable to gracefully shutdown due to err: %+v", err)
	}
}
