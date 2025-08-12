package main

import "github.com/gofiber/fiber/v2"

func TestRouteParameter(c *fiber.Ctx) error {
	userId := c.Params("userId")
	orderId := c.Params("orderId")
	return c.SendString("Get Order " + orderId + " From User " + userId)
}

func main() {
	app := fiber.New()

	app.Get("/users/:userId/orders/:orderId", TestRouteParameter)

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
