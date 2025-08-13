package config

import "github.com/gofiber/fiber/v2"

func NewFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:      true,
		AppName:      "Fiber",
		ErrorHandler: NewErrorHandler(),
	})

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return ctx.Status(code).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}
