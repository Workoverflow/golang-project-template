package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path/filepath"
)

const (
	path = "../../internal/application/config/sources/config.%s.yml"
)

type Config struct {
	Application struct {
		Environment string      `yaml:"environment"`
		Params      interface{} `yaml:"params"`
		Server      struct {
			Host    string `yaml:"host"`
			Port    int64  `yaml:"port"`
			Timeout int64  `yaml:"timeout"`
		} `yaml:"server"`
		Database struct {
			Name     string `yaml:"name"`
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
		} `yaml:"database"`
	} `yaml:"application"`
}

func GetConfig() *Config {
	config := &Config{}

	env, ok := os.LookupEnv("ENVIRONMENT")
	if ok != true {
		logrus.Fatal("Undefined environment. Set ENVIRONMENT in .env file")
	}

	configPath, _ := filepath.Abs(fmt.Sprintf(path, env))
	logrus.Infof("Use config %s", fmt.Sprintf(path, env))

	file, err := os.Open(configPath)
	if err != nil {
		logrus.Fatal("Can't open configuration file")
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	content = []byte(os.ExpandEnv(string(content)))
	if err != nil {
		logrus.Fatal("Can't read configuration file")
	}

	if err := yaml.Unmarshal(content, config); err != nil {
		logrus.Fatal("Can't parse YAML config")
	}

	return config
}
