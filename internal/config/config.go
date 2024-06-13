package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type App struct {
	Name string `yaml:"name" env-default:"cleanEnvTest"`
}

type Config struct {
	App App `yaml:"app"`
}

func MustLoad() *Config {
	const configPath = "config/config.yml"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("unable to load config file", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("unable to unmarshal config file: %s", err)
	}

	return &cfg
}
