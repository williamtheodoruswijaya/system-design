package http

import (
	"06-ddd-implementation/internal/model"
	"06-ddd-implementation/internal/model/request"
	"06-ddd-implementation/internal/model/response"
	"06-ddd-implementation/internal/usecase"
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	// Auth
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error

	// Find User Info
	FindByUserID(c *fiber.Ctx) error
	FindByUsername(c *fiber.Ctx) error
	FindByEmail(c *fiber.Ctx) error
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

func (u *UserControllerImpl) Login(c *fiber.Ctx) error {
	// step 1: ambil request dari body + ubah jadi struct
	var req *request.ValidateUserRequest
	err := c.BodyParser(&req)
	if err != nil {
		return fiber.ErrBadRequest
	}

	// step 2: buat context timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// step 3: call usecase layer
	result, err := u.UserUsecase.Login(ctx, req)
	if err != nil {
		return fiber.ErrInternalServerError
	} else {
		// step 4: kalau berhasil return hasil-nya
		return c.JSON(model.WebResponse[*response.ValidateUserResponse]{Data: result})
	}
}

func (u *UserControllerImpl) FindByUserID(c *fiber.Ctx) error {
	// step 1: ambil userID dari path URL (convert to int)
	userIDString := c.Params("userID")
	if userIDString == "" {
		return fiber.ErrBadRequest
	}
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		return fiber.ErrBadRequest
	}

	// step 2: buat context timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// step 3: go to usecase layer
	result, err := u.UserUsecase.FindByUserID(ctx, userID)
	if err != nil {
		return fiber.ErrInternalServerError
	} else {
		// step 4: return response
		return c.JSON(model.WebResponse[*response.CreateUserResponse]{Data: result})
	}
}

func (u *UserControllerImpl) FindByUsername(c *fiber.Ctx) error {
	// step 1: ambil username dari path URL
	username := c.Params("username")
	if username == "" {
		return fiber.ErrBadRequest
	}

	// step 2: create context timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// step 3: go to usecase layer
	result, err := u.UserUsecase.FindByUsername(ctx, username)
	if err != nil {
		return fiber.ErrInternalServerError
	} else {
		// step 4: return response if succeeded
		return c.JSON(model.WebResponse[*response.CreateUserResponse]{Data: result})
	}
}

func (u *UserControllerImpl) FindByEmail(c *fiber.Ctx) error {
	// step 1: ambil email dari params
	email := c.Params("email")
	if email == "" {
		return fiber.ErrBadRequest
	}

	// step 2: create context timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// step 3: go to usecase layer
	result, err := u.UserUsecase.FindByEmail(ctx, email)
	if err != nil {
		return fiber.ErrInternalServerError
	} else {
		// return response if succeeded
		return c.JSON(model.WebResponse[*response.CreateUserResponse]{Data: result})
	}
}
