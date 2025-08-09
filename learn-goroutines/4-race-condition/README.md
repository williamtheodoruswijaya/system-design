## Masalah dengan Goroutine

- Saat kita menggunakan goroutine, dia tidak hanya berjalan secara concurrent, tapi bisa parallel juga, karena bisa ada beberapa thread yang berjalan secara parallel.
- Hal ini sangat berbahaya ketika kita melakukan manipulasi data variable yang sama oleh beberapa goroutine secara bersamaan.
- Hal ini bisa menyebabkan masalah yang namanya Race Condition.

Jadi intinya kalau ada beberapa goroutine terus goroutine tersebut melakukan perubahan variable maka bisa aja terjadi yang namanya Race Condition.

Contoh kode:
```go
var x = 0
for i := 1; i <= 1000; i++ {
	// akan terdapat 1000 goroutine
	go func() {
		// setiap goroutine akan menjalankan penjumlahan x sebanyak 100 kali
		for j := 1; j <= 100; j++ {
			x = x + 1
        }   
    }()
	
	// berarti setiap goroutine akan menambahkan nilai x sebanyak 100
	// hingga akhir dari nilai x harusnya 1000 * 100 yaitu 100000
}
time.Sleep(5 * time.Second)
fmt.Println("Counter: ", x)
```

Harusnya hasil x tidak akan sampai 100000, melainkan berubah-ubah setiap running.


Kenapa hal ini bisa terjadi?
Karena bisa aja ada 2 goroutine yang mengakses X di waktu yang sama.
Contoh ketika x = 1000.
Nah ada 2 goroutine mengakses x = 1000 akibatnya, nilai x justru hilang sebagian.
Problem ini kita kenal dengan nama "Race Condition" dimana Goroutinenya saling berbalapan membuat nilai X overlap.

Gimana cara handle ini? Nah, ini ada yang namanya `Mutex`.


## Mutex (Mutual Exclusion)

- Untuk mengatasi masalah race condition, Golang terdapat sebuah struct bernama `sync.Mutex`.
- Mutex bisa digunakan untuk melakukan locking dan unlocking, dimana ketika kita melakukan locking terhadap Mutex, maka tidak ada yang bisa melakukan locking lagi sampai kita melakukan unlock.
- Dengan begitu, jika ada beberapa goroutine melakukan lock terhadap Mutex, maka hanya 1 goroutine yang diperbolehkan melakukan lock tersebut, setelah goroutine tersebut melakukan unlock, baru goroutine selanjutnya diperbolehkan melakukan lock lagi.
- Ini sangat cocok sebagai solusi ketika ada masalah race condition yang sebelumnya kita hadapi.

Contoh ada 1000 goroutine, nah, maka diminta satu" buat lockingnya, tiap goroutine yang udah beres kita unlock, setelah itu kita lock kembali pada goroutine selanjutnya.

[Code](2-mutex.go)


## RWMutex (Read-Write Mutex)

- Kadang ada kasus dimana kita ingin melakukan locking tidak hanya pada proses mengubah data, tapi juga membaca data
- Kita sebenarnya bisa menggunakan Mutex saja, namun masalahnya nanti akan rebutan antara proses membaca dan mengubah.
- Read Write Mutex memiliki perbedaan dengan Mutex biasa dimana dia memiliki dua lock, lock untuk Read dan lock untuk Write.

[code](3-RWmutex.go)


## Deadlock

- Ini kejadian yang bisa terjadi kalau kita salah implementasi Mutex atau Locking.
- Deadlock adalah keadaan dimana sebuah proses Goroutine saling menunggu lock sehingga tidak ada satupun goroutine yang bisa jalan.

Contoh simulasi [Deadlock](4-deadlock-simulation.go)


## Sync.WaitGroup