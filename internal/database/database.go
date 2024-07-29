package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	logger "github.com/opencardsonline/oco-web/logging"
)

type AppDBConn struct {
	db *pgx.Conn
}

func (_appDBConn *AppDBConn) New(connectionString string) {
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
	_appDBConn.db = conn
}

func (_appDBConn *AppDBConn) ExecuteQueryWithNoReturn(sql string, arguments ...any) error {
	_, err := _appDBConn.db.Exec(context.Background(), sql, arguments)
	if err != nil {
		logger.Log.Error("unable to execute sql", "AppDBConn.ExecuteQueryWithNoReturn", err)
		return err
	}
	return nil
}

func (_appDBConn *AppDBConn) InsertRowAndReturnLastID(sql string, arguments ...any) (lastInsertId *int, err error) {
	err = _appDBConn.db.QueryRow(context.Background(), sql, arguments).Scan(&lastInsertId)
	if err != nil {
		logger.Log.Error("unable to insert row", "AppDBConn.InsertRowAndReturnLastID", err)
		return nil, err
	}
	return lastInsertId, nil
}

func (_appDBConn *AppDBConn) QueryRows(sql string, arguments ...any) (results pgx.Rows, err error) {
	results, err = _appDBConn.db.Query(context.Background(), sql, arguments)
	if err != nil {
		logger.Log.Error("unable to query data", "AppDBConn.QueryRows", err)
		return nil, err
	}
	return results, nil
}

func (_appDBConn *AppDBConn) QueryRow(sql string, arguments ...any) (result pgx.Row) {
	result = _appDBConn.db.QueryRow(context.Background(), sql, arguments)
	return result
}
