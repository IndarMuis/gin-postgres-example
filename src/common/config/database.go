package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

func NewDatabase() *sql.DB {
	pgHost := "localhost"
	pgPort := "5432"
	pgDbName := "gin_db"
	pgUser := "postgres"
	pgPassword := "al0homora"

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", pgHost, pgPort, pgUser, pgPassword, pgDbName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("ERROR CONNECT TO DATABASE")
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	fmt.Println("SUCCESS CONNECT TO DATABASE")

	return db
}
