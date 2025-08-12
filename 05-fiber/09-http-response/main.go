package main

import "github.com/gofiber/fiber/v2"

func ReturnResponseExample(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"username": "William",
		"name":     "William as well",
	})
}

func main() {
	app := fiber.New()

	app.Get("/", ReturnResponseExample)

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
