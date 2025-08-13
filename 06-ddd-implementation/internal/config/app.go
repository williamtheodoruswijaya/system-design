package config

import (
	"06-ddd-implementation/internal/delivery/http"
	"06-ddd-implementation/internal/delivery/http/route"
	"06-ddd-implementation/internal/gateway/messaging"
	"06-ddd-implementation/internal/repository"
	"06-ddd-implementation/internal/usecase"
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
	userRepository := repository.NewUserRepository()

	// setup producer
	var userProducer messaging.UserProducer
	if config.Kafka != nil {
		userProducer = messaging.NewUserProducer(config.Kafka)
	}

	// setup usecase
	userUsecase := usecase.NewUserUsecase(config.DB, config.Validate, userRepository, userProducer)

	// setup controller
	userController := http.NewUserController(userUsecase)

	// setup middleware

	// return route config
	routeConfig := route.RouteConfig{
		App:            config.App,
		UserController: userController,
		// other controller
	}

	routeConfig.Setup()
}
