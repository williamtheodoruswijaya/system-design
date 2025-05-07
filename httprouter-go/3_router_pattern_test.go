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

func TestRouterPatternNamedParameter(t *testing.T) {
	// step 1: inisialisasi router
	router := httprouter.New()

	// step 2: buat method HTTP-nya
	router.GET("/products/:id/items/:itemId", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		// step 3: ambil parameter dari URL-nya
		id := params.ByName("id")
		itemId := params.ByName("itemId")

		// step 4: tulis response-nya (contoh aja karena ini buat unit test)
		text := "Product " + id + " Item " + itemId
		fmt.Fprint(writer, text)
	})

	// step 5: Test dengan membuat request dan recorder baru
	request := httptest.NewRequest("GET", "http://localhost:3000/products/1/items/1", nil)
	recorder := httptest.NewRecorder()

	// step 6: Buat sebuah request ke router-nya
	router.ServeHTTP(recorder, request)

	// step 7: Ambil hasilnya
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	// step 8: Bandingkan hasilnya dengan yang diharapkan
	assert.Equal(t, "Product 1 Item 1", string(body))
}

func TestRouterPatternCatchAllParameter(t *testing.T) {
	// langsung aja lah ya, ga usah diulang-ulang lagi
	router := httprouter.New()
	router.GET("/images/*image", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		image := params.ByName("image") // ambil parameter yang ditangkap oleh *image (basically sama aja kayak :image cuman dia bakal ngambil semua url dibelakang-nya)
		text := "Image: " + image
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/images/small/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Image: /small/profile.png", string(body))
}