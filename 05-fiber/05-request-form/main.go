package main

import "github.com/gofiber/fiber/v2"

func TestFormRequest(c *fiber.Ctx) error {
	name := c.FormValue("name")
	return c.SendString("Hello " + name)
}

func main() {
	app := fiber.New()

	app.Post("/hello", TestFormRequest)

	err := app.Listen("localhost:8080")
	if err != nil {
		panic(err)
	}
}
