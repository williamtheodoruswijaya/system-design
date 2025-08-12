package main

import "github.com/gofiber/fiber/v2"

type RegisterRequest struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
	Name     string `json:"name" xml:"name" form:"name"`
}

func RegisterController(c *fiber.Ctx) error {
	// 1. initialize request struct as a variable
	var request RegisterRequest

	// 2. parse body from json to our defined struct
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	// 3. call use-case process, etc...

	// 4. return response
	return c.JSON(request)
}

func main() {
	app := fiber.New()

	app.Post("/register", RegisterController)

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
