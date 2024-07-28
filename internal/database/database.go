package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	logger "github.com/opencardsonline/oco-web/logging"
)

func InitializeDBConnection(connectionString string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		logger.Log.Error("unable to connect to database", "InitializeDBConnection", err)
		os.Exit(1)
	}
	err = conn.Ping(context.Background())
	if err != nil {
		logger.Log.Error("unable to ping database", "InitializeDBConnection", err)
		os.Exit(1)
	}
	logger.Log.Info(fmt.Sprintf("Database is connected! [%s]", conn.Config().Database))
	return conn
}
