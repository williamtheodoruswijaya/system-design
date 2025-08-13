package main

import "06-ddd-implementation/internal/config"

func main() {
	// initialize semua app yang ada di config
	db := config.NewDatabase()
	app := config.NewFiber()
	kafka := config.NewProducer() // khusus kafka, yang kita use di main.go itu producer-nya (consumer itu ada di worker/main.go)
	redis := config.NewRedisClient()
	validate := config.NewValidator()

	// initialize Bootstrap dari app.go di config
	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Kafka:    kafka,
		Redis:    redis,
		Validate: validate,
	})

	// run fiber App-nya
	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
