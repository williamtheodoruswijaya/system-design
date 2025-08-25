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

[Code](3-RWmutex.go)


## Deadlock

- Ini kejadian yang bisa terjadi kalau kita salah implementasi Mutex atau Locking.
- Deadlock adalah keadaan dimana sebuah proses Goroutine saling menunggu lock sehingga tidak ada satupun goroutine yang bisa jalan.

Contoh simulasi [Deadlock](4-deadlock-simulation.go)


## Sync.WaitGroup

- Fitur yang bisa digunakan untuk menunggu sebuah proses selesai dilakukan. Biasa, WaitGroup digunakan untuk menunggu semua proses Goroutine selesai. Nah biasa kita menggunakan `thread.Sleep(1 * time.Second)`, nah tapi kan sebenernya selesai-nya Goroutine itu tidak tau (kalau pake Sleep itu tebak" aja).
- Hal ini kadang diperlukan, misal kita ingin **semua proses selesai terlebih dahulu sebelum aplikasi kita selesai**.
- Kasus seperti ini bisa menggunakan `WaitGroup`.
- Jadi ga lagi pake `thread.Sleep()` lagi.
- Untuk menandai bahwa ada proses goroutine, kita bisa menggunakan method:
  - `Add(value int)`: dimana value-nya adalah jumlah goroutine yang akan ditunggu. Ketika Add sudah 0, maka dia akan selesai.
  - Setelah proses goroutine selesai, kita bisa gunakan method `Done()` dan dia akan melakukan decrement terhadap value pada `Add()`.
- Untuk menunggu semua proses selesai, kita bisa menggunakan method `Wait()` dimana dia akan mengurangi value dalam `Add()` per selesainya sebuah goroutine.

[Code](5-waitgroup.go)


## Once

- Once adalah fitur di Golang yang bisa kita gunakan untuk memastikan bahwa sebuah function hanya dieksekusi maksimal sekali.
- Jadi berapa banyak pun goroutine yang mengakses, bisa dipastikan bahwa goroutine yang pertama yang bisa mengeksekusi function nya.
- Goroutine yang lain akan di hiraukan, artinya function tidak akan dieksekusi lagi.

[Code](6-once.go)


## Pool

- Implementasi Design Pattern bernama Object Pool Pattern di Go.
- Mirip sama Connection Pool di database, design pattern pool ini digunakan untuk **menyimpan data**, selanjutnya untuk menggunakan datanya, kita bisa mengambil dari Pool, dan setelah selesai menggunakan datanya, kita bisa menyimpan kembali ke Poolnya.
- Implementasi Pool ini intinya buat database itu ga di hit secara terus-menerus, jadi dia ibaratnya kita ambil 1 connection aja terus taro datanya ke pool, terus jadi kalau butuh data tinggal ambil pool instead create connection secara terus-menerus di Database.
- Implementasi Pool di Go-Lang ini sudah aman dari problem race condition.

[Code](7-pool.go)


## sync.Map

- Map ini mirip Golang map, namun yang membedakan adalah map ini dibuat khusus untuk concurrency menggunakan goroutine.
- Ada beberapa function dalam Map:
  - Store(key, value): untuk menyimpan data ke Map
  - Load(key): untuk mengambil value
  - Delete(key)
  - Range(function(key, value)) digunakan untuk melakukan iterasi terhadap seluruh data di Map.

[Code](8-map.go)


## sync.Cond

- Implementasi locking berbasis kondisi
- Cond membutuhkan Locker (Mutex atau RWMutex) untuk implementasi lockingnya, berbeda dengan locker biasa, di Cond terdapat function Wait() untuk menunggu (ibaratnya Goroutine yang sedang di Lock disuruh nunggu sampai diberi signal buat lakuin proses/code dibawahnya)
- Cond punya function Signal() buat ngasih tau Goroutine tidak perlu menunggu lagi.
- Ada juga function Broadcast() buat ngasih tau seluruh Goroutine
- Untuk membuat Cond, kita bisa menggunakan functionnya sync.NewCond(Locker)

[Code](9-cond.go)

Ibaratnya Cond ini ngatur setelah Locking, apakah suatu Goroutine boleh jalanin code-nya atau tidak berdasarkan perintah `Signal()` atau `Broadcast()`


## Atomic

- Intinya Race Condition problem sebelumnya yang pake Mutex buat solve nya karena manipulasi dari dua Goroutines bisa di solve dengan ganti tipe data int jadi Atomic.
- Codingan baca sendiri di dokumentasi aja (keknya gaakan kepake di project ?)


## GOMAXPROCS

- Kita tau goroutines berjalan di dalam sebuah thread. Pertanyaannya, gimana cara kita tau bahwa ada berapa Thread yang ada di Golang ketika aplikasi kita berjalan?
- Untuk mengetahuinya, kita bisa menggunakan GOMAXPROCS (konfigurasi dari Golang) yang bisa kita gunakan untuk mengubah jumlah thread atau mengambil jumlah thread.
- Secara default, jumlah thread di Golang itu sebanyak jumlah CPU di komputer kita.
- Bisa pake runtime.NumCPU() buat nyari tau CPU kita ada berapa core.

Cara ubah" total thread bisa pake `runtime.GOMAXPROCS(int)`. Value di dalam kalau dibawah 0 artinya kita pake default dari CPU, kalau diatas 0 itu kita ubah sendiri. Nah, normalnya biasa pake -1 biar bisa default tapi udah gaush setting-setting gini-ginian dah.

Kita bisa cek jumlah Goroutine pake `runtime.NumGoroutine()`.