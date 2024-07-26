package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	logger "github.com/opencardsonline/oco-web/internal/logging"
)

func InitializeDBConnection(connectionString string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	logger.Log.Info("Database is connected!")
	return conn
}
