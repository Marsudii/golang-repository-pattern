// config/config.go

package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetMySQLConnection() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	database := os.Getenv("MYSQL_DATABASE")

	return username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database
}
