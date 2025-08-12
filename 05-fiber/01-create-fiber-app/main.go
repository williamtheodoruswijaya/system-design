package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}) // config2 Fiber nanti akan dibahas

	err := app.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}
