package main

import (
	"quick-start-go/infrastructure/db"
	"quick-start-go/internal/api"
)

func main() {
	// Pertama-tama kita akan inisialisasi koneksi ke database PostgreSQL
	database := db.NewDbConnection()
	defer database.Close()

	// Nah, kita sudah punya koneksi ke database, sekarang kita bisa migrasi database-nya
	db.Migrate(database)

	// Disini kita akan manggil SetupRoutes dari routes.go
	router := api.SetupRoutes(database)

	// Sekarang, kita bisa jalanin App Server-nya
	router.Run(":8080")
}