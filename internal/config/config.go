package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Endpoint   string        `yaml:"endpoint" env-required:"true"`
	Port       string        `yaml:"port" env-default:"8080"`
	Limit      float64       `yaml:"limit" env-default:"2"`
	MaxRequest int           `yaml:"max-request" env-default:"5"`
	ResetLimit time.Duration `yaml:"reset-Limit" env-default:"60s"`
}

func MustLoadConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatal("error loading config:", err)
	}

	return &cfg
}
