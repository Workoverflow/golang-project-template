package storage

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"golang-project-template/internal/application/config"
)

type Storage interface {
	Get(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
}

type SqlxStorage struct {
	Connection *sqlx.DB
}

func NewStorage(config *config.Config) *SqlxStorage {
	db, err := NewDatabase(config).Connect()

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Connected to database")

	return &SqlxStorage{Connection: db}
}

func (s *SqlxStorage) Get(dest interface{}, query string, args ...interface{}) error {
	return s.Connection.Get(dest, query, args...)
}

func (s *SqlxStorage) Exec(query string, args ...interface{}) (sql.Result, error) {
	return s.Connection.Exec(query, args...)
}

func (s *SqlxStorage) NamedQuery(query string, args interface{}) (*sqlx.Rows, error) {
	return s.Connection.NamedQuery(query, args)
}
