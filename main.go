package main

import (
	"github.com/opencardsonline/oco-web/cmd"
	logger "github.com/opencardsonline/oco-web/internal/logging"
)

func main() {
	logger.InitializeLogger()
	logger.Log.Info("Logger is initialized!")
	cmd.CommandHandler()
}
