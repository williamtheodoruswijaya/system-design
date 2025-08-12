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


## Ctx `*fiber.Ctx`

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


## Route Parameter `.Params(<name>)`

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


## Request Form (Body - form-data) `.FormValue(<key>)`

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


## Multipart Form `.FormFile(<"key-name">)` & `.SaveFile(<file-variable>, <"target-directory">)`

- Untuk mengambil data file yang terdapat di MultipartForm, contoh kita mau mengirimkan file, kita bisa menggunakan method `FormFile()` di Ctx, namun method ini bisa menghasilkan error.
  
Mksd, dari mengembalikan error itu kek gini:
```go
// instead of writing it like this:
file := ctx.FormFile("<key-name>")

// we write it like this
file, err := ctx.FormFile("<key-name>")
if err != nil {
	panic(err)
}
```

- Selain itu, Ctx juga memiliki method `SaveFile()` untuk menyimpan file.
  
Contoh:
```go
err = ctx.SaveFile(file, "./target/"+file.Filename)
if err != nil {
	panic(err)
}
```

[Code-Lengkap](06-multipart-form/main.go)

## Request Body `.Body()`
(ini yang biasa paling sering dipake buat send data dari client ke server)

- Saat kita membuat RESTful API, kita pasti biasanya akan mengambil informasi request body yang dikirim oleh client, misal JSON, XML, dsbnya.
- Kita bisa mengambil informasi request body menggunakan method `.Body()` tanpa parameter apapun
- Kita bisa konversi menggunakan `json.Unmarshal(<body-variable>, <request-struct>)`

[Code](07-request-body/main.go)


## Body Parser `.BodyParser(<request-struct-variable>)`

- Melakukan konversi request body dari []byte menjadi struct sangat menyulitkan jika harus dilakukan manual terus-menerus
- Untungnya, Fiber bisa melakukan parsing otomatis sesuai tipe data yang dikirim, dan otomatis dikonversi ke struct.
- Ada beberapa Content-Type yang didukung oleh Fiber, caranya dengan menambahkan tag pada structnya.
- Ini bisa pakai method `.BodyParser(<request-struct-variable>)`
  
|                  Content Type                   | Struct Tag |
|:-----------------------------------------------:|:-----------|
|       `application/x-www-form-urlencoded`       | form       |
|              `multipart/form-data`              | form       |
| **`application/json`** (ini yang paling sering) | json       |
|                `application/xml`                | xml        |
|                   `text/xml`                    | xml        |
  
Contoh (kita pakai tag ini di struct-nya):
```go
type RegisterRequest struct {
    Username string `json:"username" xml:"username" form:"username"`
    Password string `json:"password" xml:"password" form:"password"`
    Name     string `json:"name" xml:"name" form:"name"`
}
```
Nanti otomatis bakal di mapping sama si Fiber-nya sesuai dengan tag yang ada di struct dengan body-nya.
```go
func RegisterController(c *fiber.Ctx) error {
	// 1. initialize request struct as a variable
	var request RegisterRequest

	// 2. parse body from json to our defined struct
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	// 3. call use-case process, etc...

	// 4. return response
	return c.JSON(request)
}

func main() {
    app := fiber.New()
    
    app.Post("/register", RegisterController)
    
    err := app.Listen(":8080")
    if err != nil {
        panic(err)
    }
}
```


## HTTP Response

- Selain sebagai representasi dari HTTP Request, *fiber.Ctx juga digunakan sebagai representasi HTTP Response.
- Sebelumnya biasa kita pakai `.SendString()` buat mengembalikan response body dalam bentuk string.
- Tapi sebenarnya kalau backend kita biasa pakai `.JSON(<struct-defined-response-body>)` tapi biasanya ada banyak lagi.
  
|  Response Method  | Keterangan                                   |
|:-----------------:|:---------------------------------------------|
| `Set(key, value)` | Mengubah header ke response                  |
| `Status(status)`  | Mengubah response status                     |
| `SendString(body` | Mengubah response body menjadi string        |
|    `XML(body)`    | Mengubah struct yang udah di define jadi XML |
|   `JSON(body)`    | Mengubah struct yang udah di define jadi XML |
| `Redirect(path)`  | Mengubah response menjadi redirect ke path   |
| `Cookie(cookie)`  | Menambah cookie ke response                  |

Contoh yang udah kita sudah praktek-an pada function `RegisterController(c *fiber.Ctx)` di[atas](08-body-parser/main.go).
  
Atau
  
liat code dibawah:  
[Code](09-http-response/main.go)
  
Tapi biasanya kita return-nya pake generic struct yang udah di define macam ini nih:
```go
type WebResponse[T any] struct {
    Data   T             `json:"data"`
    Paging *PageMetadata `json:"paging,omitempty"`
    Errors string        `json:"errors,omitempty"`
}

type UserResponse struct {
    ID        string `json:"id,omitempty"`
    Name      string `json:"name,omitempty"`
    Token     string `json:"token,omitempty"`
    CreatedAt int64  `json:"created_at,omitempty"`
    UpdatedAt int64  `json:"updated_at,omitempty"`
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	// process (call usecase layer, parse body, take value, initialize context)...
    
	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}
```


## Download File