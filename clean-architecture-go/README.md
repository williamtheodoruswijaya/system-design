# Penjelasan Kegunaan Folder-folder (package) dalam Clean Architecture

- cmd: tempat kita menaruh file `main.go` atau buat jalanin `web socket` juga bisa ditaro disini.
- internal: tempat kita menaruh `rest.api` kita (semua yang berhubungan dengan api disimpan disini).
  - api: setup konfigurasi `rest.api`, biasanya disini kita taro 3 jenis function yaitu `SetupRoutes()`, `initRoutes()`, dan `initHandlers()`.
  - repository: ini isinya query-query ke database.
  - service: ini isinya business logic aplikasi kita.
  - handler: handler akan mengakses services (menerima return value dari services) dan mengubahnya menjadi response. (Dia juga akan berperan sebagai request yang akan di hit oleh client).
- infrastructure: berisi konfigurasi-konfigurasi third-party seperti konfigurasi database, konfigurasi redis, dan sebagainya.

# Konfigurasi awal-awal

```bash
go mod init <nama_folder>
```

# Web Framework yang harus di-install

```bash
go get github.com/lib/pq
go get github.com/joho/godotenv
go get -u github.com/gin-gonic/gin
go get github.com/go-playground/validator/v10
```

Appendix: Jangan lupa tambahkan import di file `infrastructure/postgres.go`

```go
import github.com/lib/pq
```

# Penamaan nama function berdasarkan konsep access-modifiers

- huruf besar pada huruf pertama (public)
- huruf kecil pada huruf pertama (private)
