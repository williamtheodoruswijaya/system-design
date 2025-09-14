# Object-Pool

<img width="969" height="346" alt="image" src="https://github.com/user-attachments/assets/330f76d7-20f8-49e8-a9a2-82672a624b4a" />

Sederhanannya, design pattern Pool ini digunakan untuk menyimpan data, selanjutnya untuk menggunakan datanya, kita bisa mengambil dari Pool, dan setelah selesai menggunakan datanya, kita bisa menyimpan kembali ke Poolnya.

### Contoh Kasus:

Misal kita tau dalam membuat sebuah connection ke database, kita bisa pakai Singleton agar object Connectionnya tidak dibuat per-request tapi dipakai berkali-kali oleh setiap Request. Nah, misal nanti requestnya ada 100, maka akan terjadi sebuah antrian sebanyak 100 kali untuk menggunakan object Connection tersebut secara berulang-ulang. Jadi meskipun secara implementasi Singleton itu cukup bagus, tapi pada kenyataannya ini tetap menghasilkan sebuah flaw yaitu terjadinya sebuah antrian. Di satu sisi, kita gabisa asal langsung set aja tiap request untuk membuat sebuah object Connection. Oleh karena itu, ada Design Pattern yang lebih sesuai untuk kasus Database seperti ini yang juga sering digunakan dikebanyakan database service provider seperti Supabase, yaitu Connection Pool. Connection Pool sendiri mirip seperti ilustrasi diatas, jadi kita menyediakan sebuah Pool/Kumpulan Connection yang udah kita sediakan dan alih-alih mengantri satu-per-satu, request tinggal menggunakan connection sebanyak n object connection yang tersedia di pool. Nah, ini contoh dari Object Pool Design Pattern ini dan biasanya kalau mau diimplementasikan secara manual itu sangat tidak direkomendasikan, karena udah ada library bawaannya.

Contoh kalau di Go (Pool digunakan untuk menyimpan Goroutines kalau di Go):
```go
var pool sync.Pool
var wg sync.WaitGroup

// cara naro data ke dalam Pool
pool.Put("Test 1")
pool.Put("Test 2")
pool.Put("Test 3")

wg.Add(10)

for i := 0; i < 10; i++ { // ini udah aman dari Race Condition jadi gaada 1 pool dipakai 2 goroutines. Ini bakal ada 7 goroutines yang ga kebagian data di pool dan mereka akan dapat nilai default.
    go func() {
        defer wg.Done()
        // cara ambil data dari Pool
        data := pool.Get()
        fmt.Println(data)
        // cara naro data ke Pool (kalau udah pake data-nya harus di put kembali karena pool ini konsepnya Minjem)
        pool.Put(data)
    }()
}
wg.Wait()
```
Output (bisa beda-beda):
```yaml
Goroutine 3 dapet: Test 1
Goroutine 7 dapet: Test 2
Goroutine 0 dapet: Test 3
Goroutine 5 dapet: Default
Goroutine 6 dapet: Default
Goroutine 2 dapet: Default
Goroutine 8 dapet: Default
Goroutine 9 dapet: Default
Goroutine 1 dapet: Default
Goroutine 4 dapet: Default
```