package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"golang-project-template/internal/application/app"
	"golang-project-template/internal/application/config"
	"golang-project-template/internal/infrastructure/storage"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}
}

func main() {
	var runChan = make(chan os.Signal, 1)
	cfg := config.GetConfig()
	log := logrus.New()
	st := storage.NewStorage(cfg)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(cfg.Application.Server.Timeout),
	)
	defer cancel()

	// Create app instance
	app := app.Create(
		&ctx,
		cfg,
		st,
		log,
	)

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
