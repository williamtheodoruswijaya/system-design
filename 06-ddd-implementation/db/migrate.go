package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Migrate(db *sql.DB, direction string) {
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
		"postgres://%s:%s@%s:%s/%s?sslmode=require",
		config.User, config.Password, config.Host, config.Port, config.DBName,
	)

	// Koneksi ke database-nya
	migrateDB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// Cek koneksi ke database-nya
	err = migrateDB.Ping()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Baca file migrasi-nya menggunakan for loop
	files, err := ioutil.ReadDir("./db/migrations")
	if err != nil {
		log.Fatalf("failed to read migration files: %v", err)
	}

	// Urutkan file berdasarkan nama
	sort.Slice(files, func(i, j int) bool {
		fmt.Print(files[i].Name())
		return strings.Compare(files[i].Name(), files[j].Name()) < 0
	})

	// Jalankan migrasi-nya
	for _, file := range files {
		// Pastikan hanya file yang berekstensi .sql yang diambil
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
			// Tentukan mau "up" atau "down"
			if direction == "up" && strings.Contains(file.Name(), ".up.sql") {
				// Baca isi file-nya
				filePath := fmt.Sprintf("./db/migrations/%s", file.Name())
				sqlContent, err := ioutil.ReadFile(filePath)
				if err != nil {
					log.Fatalf("failed to read migration file: %v", err)
				}

				// Eksekusi SQL-nya
				_, err = migrateDB.Exec(string(sqlContent))
				if err != nil {
					log.Fatalf("failed to execute migration: %v", err)
				}
				fmt.Printf("Migrated up: %s\n", file.Name())
			}
			if direction == "down" && strings.Contains(file.Name(), ".down.sql") {
				// Baca isi file-nya
				filePath := fmt.Sprintf("./db/migrations/%s", file.Name())
				sqlContent, err := ioutil.ReadFile(filePath)
				if err != nil {
					log.Fatalf("failed to read migration file: %v", err)
				}

				// Eksekusi SQL-nya
				_, err = migrateDB.Exec(string(sqlContent))
				if err != nil {
					log.Fatalf("failed to execute migration: %v", err)
				}
				fmt.Printf("Migrated down: %s\n", file.Name())
			}
		}
	}
}
