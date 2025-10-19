package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"wallet/wallet/config"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseService struct {
	db *sql.DB
}

func InitDbService() *DatabaseService {
	db := dbConnect()
	return &DatabaseService{db}
}

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
	}

	fmt.Println("mysql was connected!")

	return db
}
