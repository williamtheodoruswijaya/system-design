package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type ContohRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetRequestBody(c *fiber.Ctx) error {
	// 1. ambil data request body-nya yang dikirim (ini masih dalam bentuk binary)
	body := c.Body()

	// 2. buat variable request-nya
	var request ContohRequest

	// 3. ubah body ke dalam bentuk request-nya
	err := json.Unmarshal(body, &request)
	if err != nil {
		return err
	}

	// process lanjutan either hit usecase layer etc...

	// return response
	return c.SendString("Hello " + request.Username)
}

func main() {
	app := fiber.New()

	app.Post("/login", GetRequestBody)

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
