package storage

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"golang-project-template/internal/application/config"
	"time"
)

const (
	MaxOpenConns = 10
	MaxIdleConns = 5
	MaxLifetime  = 10 * time.Minute
)

type Database struct {
	config *config.Config
}

func NewDatabase(config *config.Config) *Database {
	return &Database{config: config}
}

func (database *Database) getDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		database.config.Application.Database.User,
		database.config.Application.Database.Password,
		database.config.Application.Database.Host,
		database.config.Application.Database.Port,
		database.config.Application.Database.Name,
	)
}

func (database *Database) Connect() (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", database.getDSN())

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(MaxOpenConns)
	db.SetMaxIdleConns(MaxIdleConns)
	db.SetConnMaxLifetime(MaxLifetime)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
