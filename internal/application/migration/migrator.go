package migration

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"golang-project-template/internal/application/config"
	"golang-project-template/internal/infrastructure/storage"
)

const (
	path = "../../deploy/migrations/%s"
)

type Migrator struct {
	Config     *config.Config
	Storage    *storage.SqlxStorage
	Logger     *logrus.Logger
	SourcePath string
}

func NewMigrator(
	config *config.Config,
	storage *storage.SqlxStorage,
	logger *logrus.Logger,
) *Migrator {
	return &Migrator{
		Config:     config,
		Storage:    storage,
		Logger:     logger,
		SourcePath: fmt.Sprintf(path, config.Application.Database.Name),
	}
}

func (m *Migrator) Run() {
	files, err := os.ReadDir(m.SourcePath)
	if err != nil {
		m.Logger.Fatalf("can't scan dir %s", m.SourcePath)
		return
	}

	for _, file := range files {
		m.Logger.Infof("processing %s", file.Name())
		command := m.readSqlCommand(file.Name())
		_, err := m.Storage.Exec(command)
		if err != nil {
			logrus.Fatalf("can't applie migration %s", file.Name())
		}

		logrus.Infof("applied successful: %s", file.Name())
	}
}

func (m *Migrator) readSqlCommand(fileName string) string {
	sqlFile, err := os.Open(fmt.Sprintf("%s/%s", m.SourcePath, fileName))
	if err != nil {
		logrus.Fatalf("can't open migration %s", sqlFile.Name())
	}
	defer sqlFile.Close()

	command, err := io.ReadAll(sqlFile)
	if err != nil {
		logrus.Fatal("can't read file")
	}

	return string(command)
}
