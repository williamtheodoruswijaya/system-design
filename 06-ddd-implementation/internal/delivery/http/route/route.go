package route

import (
	"06-ddd-implementation/internal/delivery/http"
	"06-ddd-implementation/internal/delivery/http/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type RouteConfig struct {
	App            *fiber.App
	UserController http.UserController
}

func (c *RouteConfig) Setup() {
	// 1. pakai middleware untuk Recover(), jadi kalau ada error bisa di handle sama dia biar ga crash total
	c.App.Use(recover.New())

	// 2. pakai middleware untuk CORS, biar ga sembarang website bisa akses endpoint kita.
	c.App.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000", // Ganti dengan URL frontend Anda
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// 3. pakai middleware logger untuk mencatat detail setiap request yang masuk ke console.
	c.App.Use(logger.New())

	// 4. pakai Rate Limiter biar bisa nge-limit hit API (mencegah DDOS, dsbnya).
	c.App.Use(limiter.New(limiter.Config{
		Max:        100,             // Maksimal 100 request...
		Expiration: 1 * time.Minute, // ...dalam rentang waktu 1 menit.
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // Batasi berdasarkan alamat IP
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests",
			})
		},
	}))

	// 5. initialize route (ada 2 jenis, Guest Route itu ga perlu lewat auth_middleware.go, kalau Auth Route perlu lewat auth_middleware.go buat verify JWT Token.
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	api := c.App.Group("/api")

	users := api.Group("/users")
	{
		users.Post("/register", c.UserController.Register)
		users.Post("/login", c.UserController.Login)
	}
}

func (c *RouteConfig) SetupAuthRoute() {
	// lanjut routing api dibawah sini
	api := c.App.Group("/api")

	// pakai auth-middleware disini
	api.Use(middleware.Authenticate())

	users := api.Group("/users")
	{
		users.Get("/by-id/:userID", c.UserController.FindByUserID)
		users.Get("/by-name/:username", c.UserController.FindByUsername)
		users.Get("/by-email/:email", c.UserController.FindByEmail)
	}
}
