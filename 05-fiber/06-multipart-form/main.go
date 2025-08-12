package main

import (
	_ "embed"

	"github.com/gofiber/fiber/v2"
)

// 1. pertama, kita pakai Golang Embed untuk mengambil data file-nya biar bisa diassign ke variable dibawah dan dikonversi jadi byte array.

//go:embed source/contoh.txt
var contohFile []byte

func FormUploadController(c *fiber.Ctx) error {
	// 2. ambil file menggunakan c.FormFile() method
	file, err := c.FormFile("file") // "file" disini adalah key-name dari file tersebut disimpan.
	if err != nil {
		return err
	}

	// 3. simpan pada directory yang sudah disiapkan menggunakan method .SaveFile(variable-file, target-directory)
	err = c.SaveFile(file, "./target/"+file.Filename)
	if err != nil {
		return err
	}

	return c.SendString("Upload Success")
}

func main() {
	// 4. siapkan route-nya
	app := fiber.New()
	app.Post("/upload", FormUploadController)

	// 5. test (postman)
	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
