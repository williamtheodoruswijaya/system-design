package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	/* 	BASIC GO
	1. ini ibaratnya int main()-nya di C/C++
	2. Kalau print disini kita pake fmt.Println("...")
	3. Kalau mau dijalanin, kita pake go run main.go
	*/

	/* 	VARIABLES IN GO
	1. var x int = 10 // cara lama
	2. var x = 10 // cara baru, infer type
	3. x := 10 // cara paling baru, infer type, short declaration
	4. x = 20 // assign value to x
	5. var x, y int = 10, 20 // multiple variable declaration
	6. var x, y = 10, 20 // multiple variable declaration with infer type
	7. x, y := 10, 20 // multiple variable declaration with short declaration

	OPERATORS IN GO
	1. x++ // increment x
	2. x-- // decrement x
	3. x += 10 // add 10 to x
	4. x -= 10 // subtract 10 from x
	5. x *= 10 // multiply x by 10
	6. x /= 10 // divide x by 10
	7. x %= 10 // modulus x by 10
	8. x &= 10 // bitwise AND x with 10
	9. x |= 10 // bitwise OR x with 10
	10. x ^= 10 // bitwise XOR x with 10
	11. x <<= 10 // left shift x by 10
	12. x >>= 10 // right shift x by 10
	13. x &=^ 10 // bit clear x with 10
	*/

	/* API IN GO
	1. Pertama install dulu package "fiber/v2" (ibaratnya ini template-nya kalau di express.js)
		- Caranya: "go get github.com/gofiber/fiber/v2"
	2. Terus kita import package-nya di atas (import "github.com/gofiber/fiber/v2") bisa pake ctrl+space buat auto-complete

	3. Terus kita bisa bikin instance-nya gini:
		app := fiber.New()
	4. Terus kita bisa bikin route-nya gini:
		- buat function handler-nya dulu:
			func namaHandler(ctx *fiber.Ctx) error {
				// logic api-nya disini
			}
		- terus, lu tinggal jalanin aja si app-nya:
			app.Get("/namaRoute", namaHandler) <- ini contoh dari jalanin GET request dari handler yang udah dibuat tadi
			app.Post("/namaRoute", namaHandler) <- ini contoh dari jalanin POST request dari handler yang udah dibuat tadi
			app.Put("/namaRoute", namaHandler) <- ini contoh dari jalanin PUT request dari handler yang udah dibuat tadi
			app.Delete("/namaRoute", namaHandler) <- ini contoh dari jalanin DELETE request dari handler yang udah dibuat tadi

			app.Listen(":3000") <- ini contoh dari jalanin server di port 3000 (port bisa kita ganti-ganti)

		Contoh implementasi ada dibawah sini:
	*/

	app := fiber.New()

	app.Get("/", contohGetHandler) // ini contoh jalanin GET request

	// Baca new task 2 dibawah func Main() ini baru lanjutin codingan dibawah sini (kalau belum skip ke app.Listen(":4000"))
	// Ini bagian dari new task 3

	// step 1: setelah lu download godotenv, kita load dulu file-nya
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	// step 2: kalau ga error, kita ambil dotenv variable-nya (ambil PORT-nya sebagai contoh)
	PORT := os.Getenv("PORT")

	// Batas akhir new task 3 (bawah ini new task 2 soal API tanpa koneksi ke database)

	// Buat array dari Todo ini
	todos := []Todo{}

	// Sekarang buat post handler-nya (satuin sama App aja keknya ini best practice-nya)
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		// Buat template dari Todo (variable yang akan menampung data dari POST)
		todo := &Todo{}

		// Bind data dari POST ke variable todo (ini mirip body-parser di express.js)
			// Notes: disini try-catch-nya kita ganti pake if-else
		if err := c.BodyParser(todo); err != nil { // kalau c.BodyParser(todo) terjadi error, kita return ke variable err
			return err
		}

		// Handle juga kalau gaada data yang di-POST
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Body is required",
			})
		}

		// Nah, kalau semua udah oke, kita simpen ke array tadi
		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})

	// Buat update handler-nya
	app.Put("/api/todos/:id", func(c *fiber.Ctx) error {
		// Ambil id dari parameter-nya
		id := c.Params("id")

		// Cari todos sesuai dengan id-nya
		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = !todos[i].Completed
				return c.Status(200).JSON(todos[i])
			}
		}

		// Kalau gaada, kita return error
		return c.Status(404).JSON(fiber.Map{
			"error": "Todo ID not found",
		})
	})

	// Buat delete handler-nya
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		// Ambil id dari parameter
		id := c.Params("id")

		// Cari todos sesuai dengan id-nya
		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				// Hapus instead of Update
				todos = append(todos[:i], todos[i+1:]...)
				/*
					Logic diatas itu mirip kayak ini:
					tambahin value dari 0 sampai sebelum index i yg mau dihapus
					terus tambahin value dari index i+1 sampai akhir array-nya
					terus kita append ke array todos-nya
				*/
				return c.Status(200).JSON(fiber.Map{
					"message": "Todo deleted",
				})
			}
		}

		// Kalau id yang dicari gaada, ywd kita return error
		return c.Status(404).JSON(fiber.Map{
			"error": "Todo ID not found",
		})
	})

	// Appendix, buat handler GET semua todos-nya + GET todos by ID
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	app.Get("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for _, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				return c.Status(200).JSON(todo)
			}
		}

		return c.Status(404).JSON(fiber.Map{
			"error": "Todo ID not found",
		})
	})

	app.Listen(":" + PORT) // ini contoh jalanin server di port 3000
}

// 2. New Tasks - Kita akan coba buat API sederhana tanpa koneksi ke database

// Anggep aja ini adalah bentuk class dari data yang kita mau simpan di database (ini mirip class kalau dalam konsep OOP)
type Todo struct {
	ID			int 	`json:"id"`
	Completed 	bool 	`json:"completed"`
	Body 		string 	`json:"body"`
}

func contohGetHandler(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"data": "Ini contoh data-nya",
		"status": "success",
	})
}

// 3. New Tasks - Buat API dengan koneksi ke database
/*
1. Install package "github.com/joho/godotenv" (ini buat ngambil env variable dari .env file)
	- Caranya: "go get github.com/joho/godotenv"
*/