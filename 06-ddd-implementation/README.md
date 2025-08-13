# What is DDD (Domain-Driven Design) Architecture?

This is basically Golang Clean Architecture template.

![architecture](architecture.png)

1. External system perform request (HTTP, gRPC, Messaging, etc)
2. The Delivery creates various Model from request data
3. The Delivery calls Use Case, and execute it using Model data
4. The Use Case create Entity data for the business logic
5. The Use Case calls Repository, and execute it using Entity data
6. The Repository use Entity data to perform database operation
7. The Repository perform database operation to the database
8. The Use Case create various Model for Gateway or from Entity data
9. The Use Case calls Gateway, and execute it using Model data
10. The Gateway using Model data to construct request to external system
11. The Gateway perform request to external system (HTTP, gRPC, Messaging, etc)

---

# Setup

1. Initialize Go Modules
```bash
go mod init <nama-folder>
```

2. Install Libraries
```bash
go get github.com/lib/pq
go get github.com/joho/godotenv
go get github.com/go-playground/validator/v10
go get github.com/twmb/franz-go
go get github.com/golang-migrate/migrate/v4
go get -u github.com/golang-jwt/jwt/v5
go get github.com/gofiber/fiber/v2
go get github.com/redis/go-redis/v9
```

3. Create the folder
```bash
C:.
├───.idea
├───api             <- openapi.spec
├───cmd             <- main.go ditaro disini
├───db              <- koneksi ke database
└───internal        <- DDD architecture disini
```

4. Misal kita dah punya `openapi.spec`-nya, kita sekarang bakal pakai kafka dan http backend biasa. Nah, kalau pakai 2 itu, untuk folder `cmd` kita harus pisah jadi `web/main.go` dan `worker/main.go`, dimana web itu ibaratnya main.go yang ngejalanin ke http router kita dan worker itu main.go yang ngejalanin kafka pub-sub nya.
```bash
├───.idea
├───api
├───cmd
│   ├───web
│   └───worker
├───db
└───internal
```

5. Nah, sekarang kita taruh migrations dari Database kita di folder `db/migrations`. Dengan struktur seperti ini jadinya:
```bash
├───.idea
├───api
├───cmd
│   ├───web
│   └───worker
├───db
│   ├───migrations    <- tabel-tabel migrations
|   ├───migrate.go    <- buat create migrations-nya
│   └───connect.go    <- buat koneksi ke database-nya
└───internal
```
Cara-nya tinggal:
1. Buat .env dulu, and taro di bagian luar:
```bash
├───.idea
├───api
├───cmd
│   ├───web
│   └───worker
├───db
│   ├───migrations
|   ├───migrate.go
│   └───connect.go
├───internal
└───.env
```
2. Buat isi dari `connect.go`:
[Copas aja isi dari file ini](db/connect.go)

3. Buat isi dari migrate.go-nya:
[Copas juga dari file ini](db/migrate.go)

4. Buat tabel-tabel yang diinginkan manual pada `db/migrations` dengan format penamaan `00000<nomer-urutan>_create_<nama-table>.<up>.sql`. Misal karena kita mau fokus  ke register dan login user aja. Maka, tabelnya kek gini:
```bash
├───.idea
├───api
├───cmd
│   ├───web
│   └───worker
├───db
│   ├───migrations
│   │   ├───000001_create_users.down.sql
│   │   └───000001_create_users.up.sql
|   ├───migrate.go
│   └───connect.go
├───internal
└───.env
```
- [Lihat isi dari file `up` disini](db/migrations/000001_create_users.up.sql)
- [Lihat isi dari file `down` disini](db/migrations/000001_create_users.down.sql)
- Mereka berdua harus ada, dimana kondisi `up` dan `down`-nya akan kita setting di file `internal/config/db.go` pada  baris ini:

Notes: urutan tabel diperhatikan, kayak misal tabel Address punya Foreign key ke Users tapi tabel Users belum ada. Otomatis bakal terjadi error.
  
5. Nah, setelah itu, buat file `main.go` pada folder web dan worker. Tapi ingat, kita sekarang akan fokus establish tabel-tabelnya dulu di supabase. Jadi, kita hanya akan fokus jalanin file `connect.go` dan `migrate.go` di file `main.go` melalui file `internal/config/db.go`.
```bash
├───.idea
├───api
├───cmd
│   ├───web
│   │   └───main.go
│   └───worker
│       └───main.go
├───db
│   ├───migrations
│   │   ├───000001_create_users.down.sql
│   │   └───000001_create_users.up.sql
|   ├───migrate.go
│   └───connect.go
├───internal
└───.env
```

6. Sekarang, ikuti susunan folder internal-nya (**SAMA PERSIS!**), filenya sementara gaush yang penting foldernya dulu aja.
```bash
├───.idea
├───api
├───cmd
│   ├───web
│   └───worker
├───db
│   └───migrations
└───internal
    ├───config                      <- semua initialize app (fiber, db, redis, validator, kafka)
    │   ├───app.go
    │   ├───db.go
    │   ├───fiber.go
    │   ├───kafka.go
    │   ├───redis.go
    │   └───validator.go
    ├───delivery                
    │   ├───http                    <- controller/handler API kita taro disini
    │   │   ├───middleware
    │   │   │   └───middleware.go      
    │   │   ├───route
    │   │   │   └───route.go
    │   │   └───user_controller.go        
    │   └───messaging               <- consumer-consumer dari kafka kita taro disini
    │       ├───consumer.go
    │       └───user_consumer.go
    ├───entity                      <- entity dari tabel-tabel yang ada (model lah istilah lainnya kalau di MVC architecture)
    │   └───user_entity.go
    ├───gateway                     
    │   └───messaging               <- producer dari kafka kita buat disini
    │       ├───producer.go
    │       └───user_producer.go
    ├───model                       <- nah ini isinya bukan entity tapi lebih ke struct untuk request dan response, dan event (untuk kafka).
    │   │───converter               <- ini isinya function converter untuk ubah struct dari User ke Response, atau dari User ke Event, dsbnya.
    │   │    └───user_converter.go
    │   ├───auth.go
    │   ├───event.go
    │   ├───model.go
    │   ├───user_event.go
    │   └───user_model.go
    ├───repository                  <- ini repository layer (isinya query logic langsung ke database)
    └───usecase                     <- ini service/usecase layer (isinya business logic aplikasi)
```

---

# Step-by-step

Nah, sebenernya gaada aturan pasti mau mulai darimana, cuman biasanya biar alurnya bagus, kalau saya sih gini:

1. Pertama, **buat config**-nya dulu di **`internal/config/`**. 
   1. [`app.go`](internal/config/app.go), isinya inisialisasi semua layer yang kita punya (ini nanti aja berarti terakhir).
   2. [`db.go`](internal/config/db.go), isinya file untuk build connection pool ke database, and run migrationnya. Disini, kita akan run file `migrate.go` dan `connection.go` disini.
   3. [`fiber.go`](internal/config/fiber.go), isinya configure `fiber.App` (`fiber.App` kita initialize disini juga pakai `fiber.New()`).
   4. [`kafka.go`](internal/config/kafka.go), isinya initialize consumer dan producer dari kafka disini. (NewConsumer() & NewProducer()).
   5. [`redis.go`](internal/config/redis.go), isinya initialize client Redis.
   6. [`validator.go`](internal/config/validator.go), isinya initialize validator.
  
Notes: dengan membuat struktur seperti ini, kita dapat dengan mudah mengatur confignya pada [app.go](internal/config/app.go) pada bagian BootstrapConfig-nya. Jadi, misal kita ga butuh kafka. Nah, ywd tinggal ga ush pake. Jadi, pengaturan kayak gini tuh bikin architecture kita jadi scalable.

2. Sekarang, **untuk test connection ke database, kita coba buat beresin [`cmd/web/main.go`](cmd/web/main.go)** dah gitu coba run filenya.

3. Oke, config udah beres, connection ke database juga dah aman. Sekarang, kita lanjut ke **bagian [`entity`](internal/entity/user_entity.go). Disini, kita akan buat tabel-tabel yang ada pada database jadi sebuah Entitas struct**. Hal ini dilakukan di struktur `internal/entity`.

4. Next, kita **define [response](internal/model/response/user_response.go) dan [request](internal/model/request/user_request.go) dari setiap tabel pada directory `internal/model/<request/response>`**. Ini, kita biasanya sesuaikan dengan `api.spec` kita yang isinya itu bentuk json dari request dan response untuk masing-masing tabel. 
   1. Biasanya, kita akan buat dulu [1 generic class](internal/model/model.go) yang nampung keseluruhan response, nah bagian data-nya aja yang kita ubah berdasarkan [tabel](internal/model/response/user_response.go) response yang kita mau.
   2. Ini [Generic Class Web Response](internal/model/model.go) khusus untuk `Response` aja ya. Karena semua requestnya udah di define di file masing-masing [`user_request.go`](internal/model/request/user_request.go).

5. Kalau udah, kita lanjut **buat setup Kafka-nya**. Setup kafka sendiri bisa dilihat detailnya di[sini](../02-kafka-franz-go/README.md) atau **khusus projek ini, akan dibahas di file [`KAFKA_SETUP.md`](KAFKA_SETUP.md)**.

6. Sekarang, setelah semua-nya beres, kita bisa mulai coding pada bagian `internal/repository`. Dalam konteks ini, saya akan mulai melakukan query logic pada layer repository di file [`user_repository.go`](internal/repository/user_repository.go).

7. Setelah repository beres, kita bisa lanjut ke business logic-nya di [`user_usecase.go`](internal/usecase/user_usecase.go). Ini ada di `internal/usecase/`.

8. Baru setelah itu, kita balik lagi ke folder 'delivery/http/' buat define [`user_controller.go`](internal/delivery/http/user_controller.go).

9. Sebelum lanjut, kita pertama buat middleware-nya terlebih dahulu untuk mengatur siapa aja yang bisa mengakses api tertentu (menentukan mana yang public api dan mana yang bukan public api). Sekaligus berguna juga untuk verify JWT Token. Ini ada di folder `delivery/http/middleware/`.
   - Khusus penggunaan Fiber, middleware yang kita perlu pikirkan hanya middleware untuk verifikasi JWT Token, dan menyimpan info dari user itu di Context. Ini, kita define di [`auth_middleware.go`](internal/delivery/http/middleware/auth_middleware.go)
   - Middleware lainnya seperti **logger, recovery, dan bahkan rate limiter**. Sudah tersedia dan bisa kita langsung pakai di [`route.go`](internal/delivery/http/route/route.go).
   - Yang bakal kita pakai itu:
     - `Recover` 🛡️: Untuk mencegah server crash jika terjadi panic di salah satu request. Methodnya, **`.Use(recover.New())`**.
     - `CORS` 🌐: (Cross-Origin Resource Sharing) Agar API-mu bisa diakses oleh frontend dari domain yang berbeda. Methodnya, **`.Use(cors.New(cors.Config{...}))`**.
     - `Logger` 📝: Untuk mencatat setiap request yang masuk. Sangat berguna untuk debugging. Methodnya, **`.Use(logger.New())`**.
     - Baru Rate Limiter. Methodnya: **`.Use(limiter.New(limiter.Config{...}))`**.

10. Controller ini akan kita define path api-nya pada folder `delivery/http/route/` di file [`route.go`](internal/delivery/http/route/route.go). S

11. Last, setelah [`route.go`](internal/delivery/http/route/route.go) di define, kita bisa kembali ke [`app.go`](internal/config/app.go) dan menginisialisasi semua layer  (repository, usecase, dan controller), baru ke [`main.go`](cmd/web/main.go) menjalankan server backend kita dan [`main.go`](cmd/worker/main.go) dari worker untuk menjalankan consumer kita.

### Appendix: Configuration Each Layer

- Repository Layer: (ga pakai apa-apa, database ada di service layer)
```go
type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}
```

- Service/Usecase Layer: (database, validate, redis (optional), **repository**, **producer**) -> bisa nambah lagi tergantung services-services yang dipakai
```go
type UserUsecaseImpl struct {
	DB             *sql.DB
	Validate       *validator.Validate
	UserRepository repository.UserRepository
	UserProducer   messaging.UserProducer
}

func NewUserUsecase(db *sql.DB, validate *validator.Validate, userRepository repository.UserRepository, userProducer messaging.UserProducer) UserUsecase {
	return &UserUsecaseImpl{
		DB:             db,
		Validate:       validate,
		UserRepository: userRepository,
		UserProducer:   userProducer,
	}
}
```

- Controller/Handler layer: (**usecase**)
```go
type UserControllerImpl struct {
	UserUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) UserController {
	return &UserControllerImpl{
		UserUsecase: userUsecase,
	}
}
```