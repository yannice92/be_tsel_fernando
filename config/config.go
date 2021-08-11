package config

import (
	"github.com/joho/godotenv"
	"log"
)

func GetEnv(){
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("error load .env")
	}
}