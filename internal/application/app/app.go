package app

import (
	"context"
	"golang-project-template/internal/application/config"
	"golang-project-template/internal/infrastructure/storage"
)

type Application struct {
	Config *config.Config
	Db     *storage.SqlxStorage
}

func Create(ctx context.Context, config *config.Config) *Application {
	db := storage.NewStorage(config)

	return &Application{
		Config: config,
		Db:     db,
	}
}

func (app *Application) Start() {
	// Start app or server
}

func (app *Application) Stop(ctx context.Context) error {
	return nil
}
