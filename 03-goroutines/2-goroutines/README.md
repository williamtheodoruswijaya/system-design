## Membuat Goroutines

- Untuk membuat Goroutine di Golang sangat sederhana
- Kita hanya cukup menambahkan perintah `go` sebelum memanggil function yang akan kita jalankan dalam goroutine.
- Saat sebuah function kita jalankan dalam goroutine, function tersebut akan berjalan secara asyncrhonous, artinya tidak akan ditunggu sampai function tersebut selesai.
- Aplikasi akan lanjut berjalan ke kode program selanjutnya tanpa menunggu goroutine yang kita buat selesai.
- Kalau aplikasi mati, maka Goroutines juga akan ikut mati terlepas proses-nya beres atau tidak.

Contoh:
```go
package main

import (
	"fmt"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func main() {
	go RunHelloWorld()
	fmt.Println("Ups...")

	time.Sleep(1 * time.Second)
}
```

Output:
```md
Ups...
Hello World
```
Nah, ini menunjukkan bahwa code block selanjutnya itu dijalankan bahkan sebelum function tersebut beres. Artinya, Goroutines membuat sebuah function tersebut di skip dan di ganti ke function bawahnya.

Notes: Goroutines akan membuat function dengan return value tidak bisa ditangkap valuenya.