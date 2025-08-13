## What is DDD (Domain-Driven Design) Architecture?

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


## Setup

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


## Step-by-step

Nah, sebenernya gaada aturan pasti mau mulai darimana, cuman biasanya biar alurnya bagus, kalau saya sih gini:

1. Pertama, buat config-nya dulu. 
   1. [`app.go`](internal/config/app.go), isinya inisialisasi semua layer yang kita punya (ini nanti aja berarti terakhir).
   2. [`db.go`](internal/config/db.go), isinya file untuk build connection pool ke database, and run migrationnya. Disini, kita akan run file `migrate.go` dan `connection.go` disini.
   3. [`fiber.go`](internal/config/fiber.go), isinya configure `fiber.App` (`fiber.App` kita initialize disini juga pakai `fiber.New()`).
   4. [`kafka.go`](internal/config/kafka.go), isinya initialize consumer dan producer dari kafka disini. (NewConsumer() & NewProducer()).
   5. [`redis.go`](internal/config/redis.go), isinya initialize client Redis.
   6. [`validator.go`](internal/config/validator.go), isinya initialize validator.
2. Sekarang, untuk test connection ke database, kita coba buat beresin `cmd/web/main.go` dah gitu coba run filenya.