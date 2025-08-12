package main

import (
	"github.com/gofiber/fiber/v2"
)

func HelloWorldController(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func main() {
	app := fiber.New()

	/*
		Parameter:
			'path': string
			'handler function': func using *fiber.Ctx and returns an error
	*/
	app.Get("/", HelloWorldController)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
