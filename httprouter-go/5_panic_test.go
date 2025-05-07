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

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()

	// Ini panic handler-nya
	router.PanicHandler = func(	writer http.ResponseWriter, request *http.Request, i interface{} ) {
		// i interface{} adalah error-nya
		fmt.Fprint(writer, "Panic occurred : ", i)
	}

	// Contoh pembuatan route yang bermasalah (kita jalankan panic handler didalamnya dengan menggunakan `panic()`)
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		panic("Ups") // menjalankan panic handler diatas
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Panic occurred : Ups", string(body))
}