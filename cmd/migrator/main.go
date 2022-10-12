package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"golang-project-template/internal/application/config"
	"golang-project-template/internal/application/migration"
	"golang-project-template/internal/infrastructure/storage"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}
}

func main() {
	cfg := config.GetConfig()
	log := logrus.New()
	db := storage.NewStorage(cfg)

	// Run migrator
	mgr := migration.NewMigrator(cfg, db, log)
	mgr.Run()
}
