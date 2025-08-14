package usecase

import (
	"06-ddd-implementation/internal/entity"
	"06-ddd-implementation/internal/gateway/messaging"
	"06-ddd-implementation/internal/model/request"
	"06-ddd-implementation/internal/model/response"
	"06-ddd-implementation/internal/repository"
	"06-ddd-implementation/internal/utils"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
)

type UserUsecase interface {
	// Auth
	Register(ctx context.Context, req *request.CreateUserRequest) (*response.CreateUserResponse, error)
	Login(ctx context.Context, req *request.ValidateUserRequest) (*string, error)

	// Find User
	FindByUserID(ctx context.Context, userID int) (*response.CreateUserResponse, error)
	FindByUsername(ctx context.Context, username string) (*response.CreateUserResponse, error)
	FindByEmail(ctx context.Context, email string) (*response.CreateUserResponse, error)
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

	// 3. validate request, check email exists, check username exists
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errRequest, errEmail, errUsername error
	var existingUsername, existingEmail *entity.User

	wg.Add(3)
	go func() {
		defer wg.Done()

		mu.Lock()
		errRequest = uc.Validate.Struct(req)
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()

		mu.Lock()
		existingUsername, errUsername = uc.UserRepository.FindByUsername(ctx, uc.DB, req.Username)
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()

		mu.Lock()
		existingEmail, errEmail = uc.UserRepository.FindByEmail(ctx, uc.DB, req.Email)
		mu.Unlock()
	}()

	wg.Wait()

	if errRequest != nil {
		return nil, fmt.Errorf("error request body: " + errRequest.Error())
	}
	if errUsername == nil && existingUsername != nil {
		return nil, fmt.Errorf("username %s already exists", req.Username)
	}
	if errEmail == nil && existingEmail != nil {
		return nil, fmt.Errorf("email %s already exists", req.Email)
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

	// 5. call the repository layer
	createdUser, err := uc.UserRepository.Register(ctx, tx, user)
	if err != nil {
		return nil, err
	}

	// 6. commit transaction
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	// 7. convert the created user to response (bagusnya ini dibuat folder utils khusus convert entity ke response)
	result := utils.ConvertUserResponse(createdUser)

	// 8. publish user created event (concurrently)
	go func() {
		if uc.UserProducer.Producer != nil {
			// 7.1 ubah entity ke struct event
			userEvent := utils.ConvertUserEvent(result)

			// 7.2 send event
			err = uc.UserProducer.Producer.Publish(userEvent)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()

	// 9. return response
	return result, nil
}

func (uc *UserUsecaseImpl) Login(ctx context.Context, req *request.ValidateUserRequest) (*string, error) {
	// step 1: validate user request
	err := uc.Validate.Struct(req)
	if err != nil {
		return nil, fmt.Errorf("error request body: " + err.Error())
	}

	// step 2: find user by username
	user, err := uc.UserRepository.FindByUsername(ctx, uc.DB, req.Username)
	if err != nil {
		return nil, err
	}

	// step 3: check if user is found
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	// step 4: else, check the password
	if user.Password != req.Password {
		return nil, fmt.Errorf("invalid password")
	}

	// step 5: if correct, generate token
	userResponse := utils.ConvertUserResponse(user)
	token, err := utils.GenerateToken(userResponse)
	if err != nil {
		return nil, err
	}

	// step 6: publish event (concurrently)
	go func() {
		if uc.UserProducer.Producer != nil {
			// 6.1 ubah response ke struct event
			userEvent := utils.ConvertUserEvent(userResponse)

			// 6.2 send event
			err = uc.UserProducer.Producer.Publish(userEvent)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()

	// step 7: if success, return token
	return token, nil
}

func (uc *UserUsecaseImpl) FindByUserID(ctx context.Context, userID int) (*response.CreateUserResponse, error) {
	// step 1: check if userID is <= 0
	if userID <= 0 {
		return nil, fmt.Errorf("invalid user id")
	}

	// step 2: call repository layer
	foundedUser, err := uc.UserRepository.FindByUserID(ctx, uc.DB, userID)
	if err != nil {
		return nil, err
	}

	// step 3: convert result from repository layer (entity) to response
	result := utils.ConvertUserResponse(foundedUser)

	// step 4: return response
	return result, nil
}

func (uc *UserUsecaseImpl) FindByUsername(ctx context.Context, username string) (*response.CreateUserResponse, error) {
	// step 1: check if username == ""
	if username == "" {
		return nil, fmt.Errorf("invalid username")
	}

	// step 2: call repository layer
	foundedUser, err := uc.UserRepository.FindByUsername(ctx, uc.DB, username)
	if err != nil {
		return nil, err
	}

	// step 3: convert result from repository layer to response
	result := utils.ConvertUserResponse(foundedUser)

	// step 4: return response
	return result, nil
}

func (uc *UserUsecaseImpl) FindByEmail(ctx context.Context, email string) (*response.CreateUserResponse, error) {
	// step 1: check if email != ""
	if email == "" {
		return nil, fmt.Errorf("invalid email")
	}

	// step 2: call repository layer
	foundedUser, err := uc.UserRepository.FindByEmail(ctx, uc.DB, email)
	if err != nil {
		return nil, err
	}

	// step 3: convert result from repository layer to response
	result := utils.ConvertUserResponse(foundedUser)

	// step 4: return response
	return result, nil
}
