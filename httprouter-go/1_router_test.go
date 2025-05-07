package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// Contoh pembuatan unit test untuk httprouter (unit test buat API)
func TestRouter(t *testing.T) {
	// step 1: inisialisasi router
	router := httprouter.New()

	// step 2: buat method HTTP-nya
	/*
		template:
			router.METHOD('route', func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
				// logic handler
			})
		Penjelasan parameter:
		- writer: untuk menulis response ke client
		- request: untuk mendapatkan request dari client
		- _: untuk mendapatkan parameter dari route (jika ada ubah jadi "params" biar bisa kita ambil parameter dari URL-nya)
	*/
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		fmt.Fprint(writer, "Hello World")
	})

	// step 3: Test dengan membuat request baru
	request := httptest.NewRequest("GET", "/", nil)

	// step 4: buat recorder-nya, berfungsi untuk merekam response dari server
	recorder := httptest.NewRecorder()

	// step 5: panggil router-nya dengan request dan recorder yang sudah dibuat
	router.ServeHTTP(recorder, request)

	// step 6: ambil hasilnya
	result := recorder.Result()
	
	// step 7: baca body-nya
	body, err := io.ReadAll(result.Body)
	if err != nil {
		log.Fatal("Error")
	}

	// step 8: lakukan assert dengan expected results
	// Sekalian siapa tau mau liat hasil dari response-nya kek gimana:
	fmt.Println(body)
	assert.Equal(t, "Hello World", string(body))
}