package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"wallet-cli/config"

	_ "github.com/go-sql-driver/mysql"
)

// dbConnect -> connect to database func
func dbConnect() *sql.DB {
	connectionStr := config.GetSQLDatabaseConfig()
	db, err := sql.Open("mysql", connectionStr)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(40)
	// db.SetMaxIdleConns(40)

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	} else {
		fmt.Println("Connected!")
	}

	return db
}
