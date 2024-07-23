package cmd

import (
	"os"

	"github.com/opencardsonline/oco-web/internal/database"
)

func RunMigrations() {
	database.RunMigrations()
	os.Exit(1)
}
