package api

import (
	"06-ddd-implementation/internal/config"
	"net/http"

	"github.com/gofiber/adaptor/v2"
)

var fiberHandler http.Handler

func init() {
	db := config.NewDatabase()
	app := config.NewFiber()
	kafka := config.NewProducer()
	redis := config.NewRedisClient()
	validate := config.NewValidator()

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Kafka:    kafka,
		Redis:    redis,
		Validate: validate,
	})

	// Konversi Fiber ke http.Handler untuk Vercel
	fiberHandler = adaptor.FiberApp(app)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fiberHandler.ServeHTTP(w, r)
}
