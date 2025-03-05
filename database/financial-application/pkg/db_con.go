package pkg

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

const (
	dbUser = "db_user"
	dbPass = "password"
	dbHost = "db_host"
	dbPort = 5432
	dbName = "db_name"
)

func DbCon() (*sql.DB, error) {
	db_string := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := sql.Open("postgres", db_string)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(time.Hour)

	return db, nil
}
