package usecase

import (
	"06-ddd-implementation/internal/entity"
	"06-ddd-implementation/internal/gateway/messaging"
	"06-ddd-implementation/internal/model/request"
	"06-ddd-implementation/internal/model/response"
	"06-ddd-implementation/internal/repository"
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type UserUsecase interface {
	Register(ctx context.Context, req *request.CreateUserRequest) (*response.CreateUserResponse, error)
}

type UserUsecaseImpl struct {
	DB             *sql.DB
	Validate       *validator.Validate
	UserRepository repository.UserRepository
	UserProducer   messaging.UserProducer
}

func NewUserUsecase(db *sql.DB, validate *validator.Validate, userRepository repository.UserRepository, userProducer messaging.UserProducer) UserUsecase {
	return &UserUsecaseImpl{
		DB:             db,
		Validate:       validate,
		UserRepository: userRepository,
		UserProducer:   userProducer,
	}
}

func (uc *UserUsecaseImpl) Register(ctx context.Context, req *request.CreateUserRequest) (*response.CreateUserResponse, error) {
	// 1. begin transaction
	tx, err := uc.DB.Begin()
	if err != nil {
		return nil, err
	}

	// 2. rollback
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	// 3. validate request
	err = uc.Validate.Struct(req)
	if err != nil {
		return nil, err
	}

	// 4. convert request to entity
	user := &entity.User{
		Username:   strings.TrimSpace(strings.ToLower(req.Username)),
		Fullname:   req.Fullname,
		Email:      strings.TrimSpace(strings.ToLower(req.Email)),
		Password:   req.Password,
		ProfileUrl: sql.NullString{String: "https://upload.wikimedia.org/wikipedia/commons/a/ac/Default_pfp.jpg", Valid: true},
		CreatedAt:  time.Now(),
	}

	// 4. call the repository layer
	createdUser, err := uc.UserRepository.Register(ctx, tx, user)
	if err != nil {
		return nil, err
	}

	// 5. commit transaction
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	// 6. convert the created user to response (bagusnya ini dibuat folder utils khusus convert entity ke response)
	result := &response.CreateUserResponse{
		UserID:    createdUser.ID,
		Username:  createdUser.Username,
		Fullname:  createdUser.Fullname,
		Email:     createdUser.Email,
		CreatedAt: createdUser.CreatedAt,
	}

	// 7. return response
	return result, nil
}
