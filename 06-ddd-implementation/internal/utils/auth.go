package utils

import (
	"06-ddd-implementation/internal/model/response"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type Claims struct {
	User *response.CreateUserResponse `json:"user"`
	jwt.RegisteredClaims
}

func GenerateToken(user *response.CreateUserResponse) (*response.ValidateUserResponse, error) {
	// baca dulu si .env-nya
	err := godotenv.Load()
	if err != nil {
		err = godotenv.Load("../.env")
		if err != nil {
			log.Println("env not found, skipping...")
		}
	}

	// Buat jwt token-nya
	expTime, _ := strconv.Atoi(os.Getenv("JWT_EXPIRATION_TIME"))
	expirationTime := time.Now().Add(time.Duration(expTime) * time.Minute).Unix()
	claims := &Claims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(expirationTime, 0)),
		},
	}

	// Generate token berdasarkan claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	// return token-nya
	result := response.ValidateUserResponse{
		Token: tokenString,
	}

	return &result, nil
}
