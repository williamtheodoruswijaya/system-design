package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/middleware"
)

type LogMiddleware struct {
	Handler http.Handler
}

// Middleware basically ngikutin http.Handler buat ngejalanin ServeHTTP function
func (lm *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before Execute Handler")

	// ini auto pass ke Handler
	middleware.Handler.ServeHTTP(writer, request)

	fmt.Println("After Execute Handler")
}

func TestMiddleware(t *testing.T) {
	// Ini unit-test lengkapin sendiri dah...
}