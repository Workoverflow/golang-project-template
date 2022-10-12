package app

import (
	"context"
	"github.com/sirupsen/logrus"
	"golang-project-template/internal/application/config"
	"golang-project-template/internal/infrastructure/storage"
)

type Application struct {
	Ctx     *context.Context
	Config  *config.Config
	Storage *storage.SqlxStorage
	Logger  *logrus.Logger
}

func Create(
	ctx *context.Context,
	config *config.Config,
	storage *storage.SqlxStorage,
	log *logrus.Logger,
) *Application {

	return &Application{
		Ctx:     ctx,
		Config:  config,
		Storage: storage,
		Logger:  log,
	}
}

func (app *Application) Start() {
	// Start app or server
}

func (app *Application) Stop(ctx context.Context) error {
	return nil
}
