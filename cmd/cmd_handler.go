package cmd

import (
	"flag"
	"fmt"
)

var (
	cmdserver     = "server"
	cmdmigrations = "migrations"
)

func CommandHandler() {
	command := flag.String("command", "server", "which command to run when executing the CLI. Options: [server,migrations]")

	flag.Parse()

	if *command == cmdserver {
		StartServer()
	} else if *command == cmdmigrations {
		RunMigrations()
	} else {
		fmt.Printf("The command [%s] is not valid.", *command)
	}

}
