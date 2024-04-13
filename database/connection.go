package database

import (
	"database/sql"
	"time"
	"wallet-cli/lib/config"
)

// dbConnect -> connect to database func
func dbConnect() *sql.DB {
	connectionStr := config.GetSQLDatabaseConfig()
	db, err := sql.Open("mysql", connectionStr)
	if err != nil {
		return nil
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(0)
	// db.SetMaxIdleConns(40)
	return db
}
