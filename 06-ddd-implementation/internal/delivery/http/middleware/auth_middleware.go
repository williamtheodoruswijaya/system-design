package middleware

import (
	"06-ddd-implementation/internal/model/response"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type Claims struct {
	User *response.CreateUserResponse `json:"user"`
	jwt.RegisteredClaims
}

func Authenticate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// step 1: ambil header authorization dari request (ambil JWT Token-nya)
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return fiber.ErrUnauthorized
		}

		// step 2: validate token (sekalian ambil claims untuk digunakan sebagai data user (jadi user gaush pass id terus di client)
		parsedClaims, err := ValidateToken(authHeader)
		if err != nil {
			return fiber.ErrUnauthorized
		}

		// step 3: Simpan ke context bawaan Go
		ctx := c.UserContext()
		if parsedClaims.User != nil {
			ctx = context.WithValue(ctx, "username", parsedClaims.User.Username)
			if parsedClaims.User.UserID != 0 {
				ctx = context.WithValue(ctx, "userid", parsedClaims.User.UserID)
			} else {
				return fiber.ErrUnauthorized
			}
		} else {
			return fiber.ErrUnauthorized
		}

		// step 4: update context di Fiber
		c.SetUserContext(ctx)

		// step 5: simpan token ke Fiber locals biar gampang di akses
		c.Locals("token", parsedClaims)

		// step 6: lanjut ke handler selanjutnya
		return c.Next()
	}
}

func ValidateToken(authHeader string) (*Claims, error) {
	// step 1: pastiin token diawali dengan "Bearer "
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, fmt.Errorf("authorization header format must be Bearer {token}")
	}

	// step 2: ambil token dari header
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == "" {
		return nil, fmt.Errorf("token string is empty")
	}

	// step 3: load JWT secret key dari .env file
	err := godotenv.Load()
	if err != nil {
		err = godotenv.Load()
		if err != nil {
			log.Println("Warning: .env file not found, relying on environment variables")
		}
	}
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		log.Println("FATAL: JWT_SECRET_KEY environment variable is not set")
		return nil, fmt.Errorf("server configuration error: JWT_SECRET_KEY is not set")
	}

	// step 4: parse token menjadi claims dengan mengubah ke bentuk []byte dimana []byte adalah tipe data yang dibutuhkan oleh jwt.ParseWithClaims
	jwtSecret := []byte(jwtSecretKey)

	// step 5: buat claims untuk menyimpan data user
	claims := &Claims{}
	jwtToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// step 6: pastikan token menggunakan algoritma HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	// step 7: jika ada error saat parsing token (termasuk jika token expired), kembalikan error (ERROR HANDLING)
	if err != nil {
		log.Printf("Token parsing error: %v\n", err)
		if err == jwt.ErrTokenExpired {
			return nil, fmt.Errorf("token expired")
		}
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	// step 8: jika token tidak valid, kembalikan error
	if !jwtToken.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	// step 9: jika semua langkah berhasil, kembalikan claims yang berisi data user (tambah validasi sedikit)
	if claims.User == nil {
		return nil, fmt.Errorf("user data missing in token claims")
	}
	if claims.User.Username == "" {
		return nil, fmt.Errorf("username missing in token claims")
	}

	return claims, nil
}
