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

- Fiber juga bisa digunakan untuk mengembalikan resposne File atau `byte[]` (data type File di Golang, biasa kek waktu bikin variable dari Golang Embed)
- Sama, methodnya juga ada banyak tergantung kegunaannya.

|    Method Download File    |                                 Keterangan                                 |
|:--------------------------:|:--------------------------------------------------------------------------:|
| `Download(file, filename)` | Mengubah response menjadi isi dari file-nya (dan dipaksa agar di download) |
|      `Send([]byte)`        |                  Mengubah response menjadi data `[]byte`                   |
| `SendFile(file, compress)` |                    Mengubah response menjadi isi file                      |


## Routing Group

- Saat kita membuat aplikasi web dengan banyak endpoint, kadang-kadang kita perlu melakukan Grouping beberapa routing
- Hal ini agar lebih rapi dan mudah dikembangkan ketika membuat banyak routing
- Kita bisa membuat grup dengan menggunakan `Group()` di `fiber.App`
- Dan kita bisa menambahkan Routing ke `Group()` yang sudah kita buat.

Contoh kode:
```go
func TestRoutingGroup() {
	helloWorld := func(c *fiber.Ctx) error { // dummy handler function
		return c.SendString("Hello World!")
    }
	
	// group 1: "/api/hello"
	api := app.Group("/api")
	api.Get("/hello", helloWorld)
	api.Get("/world", helloWorld)
	
	// group 2: "/api/web"
	web := app.Group("/web")
	web.Get("/hello", helloWorld)
	web.Get("/world", helloWorld)
}
```


## Pre Fork

- Secara default, saat aplikasi berjalan, dia akan berjalan secara standalone dalam satu proses.
- Kadang-kadang, ketika kita jalankan dalam sebuah server yang jumlah CPU nya banyak, mungkin ada baiknya kita menjalankan beberapa proses agar semua CPU terpakai dengan optimal.
- Namun yang jadi masalah, jika kita jalankan beberapa aplikasi Fiber secara bersamaan, maka kita harus menggunakan port yang berbeda-beda.
- Dengan Pre Fork ini, kita bisa membuat server menjalankan lebih dari 1 aplikasi Fiber sesuai dengan jumlah CPUnya tanpa menggunakan port yang baru.
- Hal ini yang mungkin membuat Fiber lebih cepat daripada Gin dan framework lain karena konfigurasi ini (?)
- Fiber memiliki konfigurasi bernama `PreFork`, yang defaultnya adalah false, dan bisa kita ubah menjadi True.
- PreFork adalah fitur menjalankan beberapa proses Fiber namun menggunakan Port yang sama.
- Ini bawaan dari sistem operasi Linux, Unix, Map.
- Teknik ini biasa digunakan di NGINX. Namanya, Socket Sharding.

Cara-nya literally:
```go
app := fiber.New(fiber.Config{
	Prefork: true
})
```
Misal, processor kita ada 10, maka ada 10 fiber server yang berjalan. Tapi, diantara semuanya, ada yang namanya Child dan Parent, dimana Child adalah Worker dan Parent adalah yang manage-nya. Biasa Child/Worker itu yang jumlahnya sesuai dengan CPU kita.


## Error Handling

- Kita bisa perhatikan kalau kita lihat Fiber Handler, itu return value-nya error, yang artinya kalau error, kita tinggal return errornya.
- Dan ketika terjadi error, secara otomatis akan ditangkap oleh ErrorHandler pada `fiber.Config`
- Kita bisa mengubah implementasi dari Error Handler-nya kalau kita mau

```go
app := fiber.New(fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		c.Status(fiber.StatusInternalServerError)
		return c.SendString("Error": + err.Error())
    }
})
```

Biasa sih kalau best-practices kita pisah jadi kek begini:
```go
func NewFiber(config *viper.Viper) *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName:      "nama aplikasi ceritanya",
		ErrorHandler: NewErrorHandler(),
		Prefork:      True,
	})

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return ctx.Status(code).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}
```


## Template

- Fiber secara default tidak memiliki template engine seperti Web Framework kebanyakan (ga kek Laravel ada Blade misal jadi bisa ngoding HTML, CSS, JS di Laravel).
- Tapi Fiber bisa integrasi sama template engine yang udah terkenal.
- Kita bisa liat semua yang di dukung disini: https://docs.gofiber.io/template
- Kek contoh pake mustache di Fiber (ya kalo mo nyoba coba aja ada di Video PZN 01:17:41)


## Middleware

- Fiber juga mendukung Middleware
- Dengan Middleware, kita bisa membuat Handler yang bisa melakukan **sebelum** dan **setelah** _request itu dikerjakan oleh Handler yang memproses request-nya_.
- Fiber sendiri menyediakan banyak Middleware.
- Tapi sekarang kita akan mencoba untuk membuat Middleware terlebih dahulu.
- Membuat Middleware itu sederhana, cukup membuat Handler seperti pada Routing.
- Namun dalam Handlernya, jika kita ingin meneruskan Request ke Handler selanjutnya, kita perlu memanggil `Next()` pada Ctx.
  
Contoh:
```go
func middlewareExample(c *fiber.Ctx) error {
	fmt.Println("I'm middleware before processing request")
	
	// jalankan handler/controller-nya
	err := c.Next()
	
	fmt.Println("I'm middleware after processing request")
	
	return err
}
```
Jadi, kunci-nya disini ada di penempatan `c.Next()` yaitu pemanggilan controller/handler yang kita ingin wrap dalam middleware. Biasanya, sebelum masuk ke controller, kita bisa passing ID dari user atau pun melakukan autentikasi middleware seperti untuk API ini dipanggil, kita harus memasukkan JWT Token terlebih dahulu baru bisa lanjut ke handlernya, dsbnya.


## Prefix

- Middleware kalau kita pake `.Use()` di group tertentu, itu biasanya kita pakai tu middleware di semua endpoint yang berkorelasi dalam group tersebut.
  
Contoh:
```go
app.Use("/api", func(c *fiber.Ctx) error {
	fmt.Println("I'm middleware before processing request")
	err := c.Next()
	fmt.Println("I'm middleware after processing request")
})

app.Get("/hello", func(c *fiber.Ctx) error {
	return c.SendString("hello world")
})
```
Artinya kek yang diatas, kalau kita hit "/hello", itu gaakan jalan middlewarenya, tapi kalau pake "/api/hello", baru jalan middlewarenya.


## Middleware Lainnya.

- Fiber sudah menyediakan banyak sekali middleware yang bisa kita gunakan secara langsung, seperti RequestId, BasicAuth, FileSystem, ETag, dan sebagainya.
- Kita bisa lihat dokumentasi penggunaannya di https://docs.gofiber.io/category/-middleware