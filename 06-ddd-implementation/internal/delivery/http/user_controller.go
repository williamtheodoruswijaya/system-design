package http

import (
	"06-ddd-implementation/internal/model"
	"06-ddd-implementation/internal/model/request"
	"06-ddd-implementation/internal/model/response"
	"06-ddd-implementation/internal/usecase"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	Register(c *fiber.Ctx) error
}

type UserControllerImpl struct {
	UserUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) UserController {
	return &UserControllerImpl{
		UserUsecase: userUsecase,
	}
}

func (u *UserControllerImpl) Register(c *fiber.Ctx) error {
	// 1. ambil request dari body
	var req request.CreateUserRequest
	err := c.BodyParser(&req)
	if err != nil {
		return fiber.ErrBadRequest
	}

	// 2. initialize context timeout signal
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 3. call use-case layer
	result, err := u.UserUsecase.Register(ctx, &req)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	// 4. kalau berhasil return success, kalau gagal return failed
	return c.JSON(model.WebResponse[*response.CreateUserResponse]{Data: result})
}
