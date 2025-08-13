package config

import (
	"06-ddd-implementation/db"
	"database/sql"
)

func NewDatabase() *sql.DB {
	// step 1: inisialisasi koneksi ke database PostgreSQL
	database := db.NewDbConnection()

	// step 2: migrate tabelnya
	db.Migrate(database, "up")

	return database
}
