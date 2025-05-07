package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// Contoh pembuatan middleware untuk logging
type LogMiddleware struct {
	http.Handler
}

// Ini middleware-nya
func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// Cukup simple aja kita print deh si request-nya
	fmt.Println("Receive Request")

	// forward ke router handler-nya
	middleware.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		fmt.Fprint(writer, "Middleware Test")
	})
	// Buat middleware-nya (wrap router ke dalam middleware)
	middleware := LogMiddleware{Handler: router}

	// Buat request dan recorder-nya
	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	// Jalankan middleware-nya
	middleware.ServeHTTP(recorder, request)

	// Ambil response-nya
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	// Cek apakah response-nya sesuai
	assert.Equal(t, "Middleware Test", string(body))
}