package config

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/twmb/franz-go/pkg/kgo"
)

type BootstrapConfig struct {
	DB       *sql.DB
	App      *fiber.App
	Kafka    *kgo.Client
	Validate *validator.Validate
	Redis    *redis.Client
}

// function ini baru bisa dibuat setelah semua configuration selesai sampai ke layer-layernya. (mainly, main.go ga kita sentuh lagi)
func Bootstrap(config *BootstrapConfig) {
	// setup repositories

	// setup producer

	// setup usecase

	// setup controller

	// setup middleware
}
