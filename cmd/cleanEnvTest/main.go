package main

import (
	"github.com/alhaos/cleanEnvTest/internal/config"
	"github.com/alhaos/cleanEnvTest/internal/logger"
)

func main() {

	cfg := config.MustLoad()

	log := logger.New()

	log.Debug(cfg.App.Name)

}
