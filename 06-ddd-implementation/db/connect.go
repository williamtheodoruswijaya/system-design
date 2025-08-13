package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // ini wajib ada jangan lupa ges
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewDbConnection() *sql.DB {
	// Load .env file-nya
	err := godotenv.Load()
	if err != nil {
		err = godotenv.Load("../.env")
		if err != nil {
			log.Fatal(err)
		}
	}

	// Ambil konfigurasi dari environment variable
	config := Config{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		DBName:   os.Getenv("DBNAME"),
	}

	// Buat connection string-nya (ini template sih)
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		config.Host, config.Port, config.User, config.Password, config.DBName,
	)

	// Koneksi ke database-nya
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// Coba lakuin ping ke database-nya
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Kalau semua berhasil, ywd kita return db-nya
	return db // ini literally pointer sql.DB
}
