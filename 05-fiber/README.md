## Apa itu Fiber?

- Golang punya HTTPRouter buat bikin API. Tapi, terkadang ada beberapa hal yang masih dilakukan secara manual menggunakan bawaan dari Golang HTTP. Contoh, konversi JSON itu harus setup sendiri di Middleware.
- Maka dari itu, ada Web Framework yang bisa konversi response ke JSON secara otomatis.
- Fiber adalah Web Framework untuk Golang yang terinspirasi dari ExpressJS, oleh karena itu cara penggunaannya sangat mudah dan sederhana kayak ExpressJS.
- https://gofiber.io/


## Cara pakai Fiber gimana?

```go
go mod init <nama_folder>
	
go get github.com/gofiber/fiber/v2
```


## Fiber App

- Saat kita menggunakan Fiber, hal pertama yang perlu kita buat adalah fiber.App
- Untuk membuatnya kita bisa menggunakan `fiber.New(fiber.Config)` yang menghasilkan pointer `*fiber.App`
- `fiber.App` adalah representasi dari aplikasi Web Fiber.
- Setelah membuat `fiber.App`, selanjutnya untuk menjalankan webnya, kita bisa menggunakan method `Listen(address)`.
  
Contoh:
```go
func main() {
	app := fiber.New(fiber.Config{
	    // ...	
    })
	
	err := app.Listen(fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
	    log.Fatalf("Failed to start server: %v", err)	
    }
}
```


## Configuration

- Saat kita membuat `fiber.App` menggunakan `fiber.New()` terdapat parameter `fiber.Config{}` yang bisa kita gunakan
- Ada banyak sekali konfigurasi yang bisa kita ubah.
- Contohnya, yaitu mengubah konfigurasi timeout yaitu konfigurasi untuk timeout otomatis (misal request write lebih dari 5 detik).

[Code](01-create-fiber-app/main.go)


## Routing

- Saat kita menggunakan web framework, pertama yang harus kita pikirkan adalah bagaimana cara membuat endpointnya?
- Di Fiber, untuk membuat Routing, sudah disediakan semua Method di fiber.App yang sesuai dengan HTTP Method.
- Parameternya membutuhkan 2, yaitu path-nya dan juga Handler/Controller-nya (`fiber.Handler`)
  
Contoh:
```go
app.Get('/register', c.UserController.Create)
```

Biasanya routing-routing ini kita simpan di directory `delivery/http/route/route.go` kalau mengikuti DDD (Domain-Driven Architecture) yang ada.