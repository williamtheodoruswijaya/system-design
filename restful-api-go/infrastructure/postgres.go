package infrastructure

import (
	"database/sql"

	"github.com/joho/godotenv"
)

type Config struct {
	Host 		string
	Port 		string
	User 		string
	Password 	string
	DbName 		string
}



func NewDbConnection() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}