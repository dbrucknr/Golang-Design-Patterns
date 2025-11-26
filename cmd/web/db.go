package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql" // Prefix with _ since this import is not used directly
)

// https://github.com/go-sql-driver/mysql

const (
	maxOpenDbConn = 5
	maxIdleDbConn = 5
	maxDbLifetime = 5 * time.Minute
)

func initMySqlDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// Test the connection
	if err = db.Ping(); err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDbConn)
	db.SetConnMaxLifetime(maxDbLifetime)
	return db, nil
}
