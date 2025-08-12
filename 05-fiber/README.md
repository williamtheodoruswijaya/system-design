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


## Ctx

- Saat kita membuat handler di Fiber Routing, kita cukup menggunakan parameter `fiber.Ctx`.
- `Ctx` ini merupakan representasi dari Request dan juga Response di Fiber. Jadi, Request dan Response semuanya disimpan di Ctx ini.
- Oleh karena itu, kita bisa __**mendapatkan informasi HTTP Request**__, dan juga bisa __**membuat HTTP Response menggunakan `fiber.Ctx`**__
  
Contoh sederhana penggunaan context dalam mengambil query parameter dari path
```go
func exampleController(c *fiber.Ctx) {
	name := c.Query("name", "Guest")
	return c.SendString("Hello " + name)
}

func main() {
	app := fiber.New()
	
    app.Get("/hello", exampleController) // tar cara akses-nya '/hello?name=William'
	
	err := app.Listen("localhost:8080")
	if err != nil {
		panic(err)
    }
}
```

Jadi selain mengambil query parameter, context dari Fiber juga bisa **mengambil Request Body** juga.


## HTTP Request

- Representasi dari HTTP Request di Fiber adalah `Ctx`.
- Untuk mengambil informasi dari HTTP Request, kita bisa menggunakan banyak sekali method yang terdapat di Ctx.
- Daftar method ada di: `https://pkg.go.dev/github.com/gofiber/fiber/v2#Ctx`

Contoh:
```go
func ExampleHTTPRequest(c *fiber.Ctx) error {
	// mengambil data dari header
	first := c.Get("firstname")
	
	// mengambil data dari cookies
	last := c.Cookies("lastname")
	
	// return response string
	return c.SendString("Hello " + first + " " + last)
}

func main() {
	app := fiber.New()
	
	app.Get("/request", ExampleHTTPRequest)
	
	err := app.Listen(":8080")
	if err != nil {
		panic(err)
    }
}
```


## Route Parameter

- Kita bisa menambahkan parameter di PATH Urlnya.
- Biasa cocok untuk pengiriman data via PATH Url.
- Saat membuat Route Parameter, kita perlu memberi nama di Route, dan di Ctx, kita bisa mengambil seluruh data menggunakan method `AllParams()`, atau menggunakan method `Params(<nama-route-parameter>)`.
  
Contoh:
```go
func TestRouteParameter(c *fiber.Ctx) error {
	userId := c.Params("userId")
	orderId := c.Params("orderId")
	return c.SendString("Get Order " + orderId + " From User " + userId)
}

func main() {
	app := fiber.New()
	
	app.Get("/users/:userId/orders/:orderId", TestRouteParameter)
	
	err := app.Listen(":8080")
	if err != nil {
	    panic(err)	
    }
}
```

[Code](04-route-parameter/main.go)


## Request Form (Body - form-data)

- Ketika kita mengirimkan data menggunakan HTTP Form, kita bisa menggunakan method `FormValue(<nama-key>)` pada `Ctx` untuk mendapatkan data yang dikirimnya.
  
Contoh:
```go
func TestFormRequest(c *fiber.Ctx) error {
	name := c.FormValue("name")
	return c.SendString("Hello " + name)
}

func main() {
	app := fiber.New()
	
	app.Post("/hello", TestFormRequest)
	
	err := app.Listen("localhost:8080")
	if err != nil {
		panic(err)
    }
}
```
Notes: ini bukan parse JSON dari body tapi cuman kalau pakai form-data. (cara test-nya di Postman pake form-data bukan Raw)


## Multipart Form

- P


## Request Body 
(ini yang biasa paling sering dipake buat send data dari client ke server)