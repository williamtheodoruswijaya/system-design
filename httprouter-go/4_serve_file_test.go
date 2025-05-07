package main

import (
	"embed"
	"io"
	"io/fs"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

// ini namanya Golang Embed buat load file: (komentar diatas var resources wajib ada ya)

//go:embed resources/*
var resources embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()

	// Biar gaush ribet nulis requested routes-nya "files/resources/nama_file" jadi "files/nama_file" buat sub-folder
	directory, err := fs.Sub(resources, "resources") // template-nya fs.Sub(resources, "<nama_folder>")
	if err != nil {
		log.Fatal(err)
	}
	
	// Notes: ini route-nya wajib "files/*filepath"
	router.ServeFiles("/files/*filepath", http.FS(directory))

	// Baru kita bisa lanjut minta request ke server-nya:
	request := httptest.NewRequest("GET", "http://localhost:3000/files/serve_file_test.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Halo, aku di load dari file ini (golang embed)", string(body))
}