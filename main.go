package main

import (
	"github.com/opencardsonline/oco-web/cmd"
	logger "github.com/opencardsonline/oco-web/logging"
)

func main() {
	log := logger.New()
	log.Info("logger is initialized")
	cmd.CommandHandler()
}
