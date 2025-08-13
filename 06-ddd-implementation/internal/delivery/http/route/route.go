package route

import (
	"06-ddd-implementation/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	UserController http.UserController
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/users", c.UserController.Register)
}

func (c *RouteConfig) SetupAuthRoute() {}
