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

// Cara penggunaan params untuk mengambil parameter dari URL-nya (contoh: /user/:id)
func TestParams(t *testing.T) {
	// step 1: inisialisasi router
	router := httprouter.New()

	// step 2: buat method HTTP-nya
	router.GET("/products/:params_yg_mau_diambil", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		// step 3: ambil parameter dari URL-nya
		id := params.ByName("params_yg_mau_diambil")

		// step 4: tulis response-nya (contoh aja karena ini buat unit test)
		test := "Product " + id
		fmt.Fprint(writer, test)
	})

	// step 5: Test dengan membuat request dan recorder baru
	request := httptest.NewRequest("GET", "http://localhost:3000/products/1", nil)
	recorder := httptest.NewRecorder()

	// step 6: Buat sebuah request ke router-nya
	router.ServeHTTP(recorder, request)

	// step 7: Ambil hasilnya
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1", string(body))
}