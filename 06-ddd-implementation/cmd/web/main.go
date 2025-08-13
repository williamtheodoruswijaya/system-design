package main

import "06-ddd-implementation/internal/config"

func main() {
	db := config.NewDatabase()
	defer db.Close()
}
