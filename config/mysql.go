package config

import (
	"database/sql"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
	"time"
)

func GetConnection() *sql.DB{
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func close(db sql.DB){
	defer func() {
		if err := db.Close(); err != nil {
			log.Print(err)
		}
	}()
}