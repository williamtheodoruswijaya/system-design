# Redis Data Structure

## Pengenalan Redis Data Structure

<ol>
  <li>Sebelumnya kita sempat menggunakan dan tau salah satu data structures di Redis yaitu data structure String.</li>
  <li>Nah, sebenarnya, ada banyak sekali struktur data yang bisa kita gunakan di Redis.</li>
  <li>Dan di Redis, ada banyak jenis struktur data dengan kegunaan masing-masing pada kasus-kasus yang berbeda.</li>
</ol>

| Struktur Data | Keterangan                                                                 |
|---------------|----------------------------------------------------------------------------|
| Lists         | Struktur data Linked List yang berisi data string                          |
| Sets          | Koleksi data string yang tidak berurut dan harus unique                    |
| Hashes        | Struktur data key-value                                                    |
| Sorted Sets   | Struktur data seperti Sets, namun berurut                                  |
| Stream        | Struktur data seperti log yang selalu bertambah di belakang (queue)        |
| Geospatial    | Struktur data koordinat                                                    |
| HyperLogLog   | Struktur data untuk melakukan estimasi jumlah unique data di Set           |

## Lists

<ol>
  <li>Lists adalah struktur data berupa Linked List yang berisi string.</li>
  <li>Lists mirip dengan array, jadi tiap data punya index, tapi dia dinamis aja.</li>
  <li>Lists bisa digunakan untuk membuat Queue atau Stack.</li>
  <li>
    <img width="711" height="91" alt="image" src="https://github.com/user-attachments/assets/c735d782-585e-488b-8c0b-5c378065949f" />
  </li>
  <li>Command untuk List dapat dicek disini: https://redis.io/commands/?group=list</li>
  <li>Operasi yang biasa sering digunakan disini seperti:
    <pre>
      LPUSH key element [element ...]
      RPUSH key element [element ...]
      LPOP key [count]
      LRANGE key start stop
    </pre>
  </li>
  <li>Semua yang berhubungan dengan head (pushHead, popHead, dsbnya) diawali dengan huruf <b>L</b></li>
  <li>Semua yang berhubungan dengan tail (pushTail, popTail, dsbnya) diawali dengan huruf <b>R</b></li>
  <li>Sebenernya asal udah paham konsep Linked List, harusnya dah aman sih buat bagian ini.</li>
</ol>

## Sets

<ol>
  <li>Sets adalah struktur data mirip seperti Lists, yang membedakan adalah, pada Sets, isi data harus unik.</li>
  <li>Jika data tidak unik, maka data tersebut tidak akan diterima.</li>
  <li>Intinya, kalau data yang sebelumnya sudah ada, maka otomatis data tersebut tidak akan diterima.</li>
  <li>Data di Sets itu tidak berurutan sesuai waktu kita masuk data ke Sets, jadi kita tidak bisa menjamin urutan dalam Sets itu selalu berurutan.</li>
  <li>Command untuk Set dapat dicek disini: https://redis.io/commands/?group=set</li>
  <li>Operasi yang biasa sering digunakan disini seperti:
    <pre>
      SADD key member [member ...]          // menambahkan member ke dalam set
      SCARD key                             // menghitung jumlah data dalam sets
      SMEMBERS key                          // mengambil seluruh data di sets
      SREM key member [member...]           // menghapus data dari sets
    </pre>
  </li>
  <li>Key di Lists atau Set anggep aja seperti nama variable dari Lists atau Set tersebut.</li>
</ol>

#### Membandingkan Set

<ol>
  <li>Karena Sets adalah struktur data yang berisi nilai-nilai yang unik, jadi terdapat operasi yang bisa kita gunakan untuk membandingkan antar Sets.</li>
  <li>SDIFF -> untuk melihat perbedaan (different) dari Sets pertama dengan Sets lainnya.</li>
  <li>SINTER -> untuk melihat kesamaan (intersect) dari beberapa Sets.</li>
  <li>SUNION -> untuk melihat gabungan unik (union) dari beberapa Sets.</li>
</ol>

## Hashes

<ol>
  <li>Hashes adalah struktur data berbentuk pair juga seperti Redis (key-value).</li>
  <li>Dengan menggunakan struktur data Hashes ini, kita bisa menentukan key untuk value yang kita ingin gunakan.</li>
  <li>Berbeda dengan Lists yang menggunakan index, pada Hashes, kita bisa menggunakan key apapun yang kita mau.</li>
  <li>Anggep aja kalau Lists kita kayak begini: arr[0], arr[1].</li>
  <li>Nah, kalau Hashes kita jadinya begini: arr["ini key"], arr["ini key 2"]</li>
  <li>Command pada Hashes ini bisa dilihat di sini: https://redis.io/commands/?group=hash</li>
  <li>Contoh yang sering:
    <pre>
      HSET key field value [field value ...]                // membuat key-value
      HGETALL key                                           // mengambil seluruh data yang ada di dalam key
      HGET key                                              // mengambil data yang ada di key secara spesifik
    </pre>
  </li>
  <li>Karna Redis sendiri bentuknya seperti Hashmap/Associative array. Hashes disini bisa kita anggap sebagai Hash Table yang ada banyak dan di map ke dalam sebuah Hashmap lagi.</li>
  <li>Tapi kita bisa visualisasikan Hashes ini dengan bentuk ini:</li>
  <img width="802" height="450" alt="image" src="https://github.com/user-attachments/assets/5fdfc09d-4ea7-4da5-b49d-322e6e20b76e" />
  <li>Contoh Penggunaan Simple:<br/>
    <pre><code>
      hset "ini-nama-hash-tablenya" key-1 "value-1" key-2 "value-2" key-3 "value-3"
      hget "ini-nama-hash-tablenya" key-1
    </code></pre>
  </li>
  <li>Meskipun ini cenderung cocok untuk dijadikan sebagai cache, tapi pada kenyataannya, saya sendiri seringnya langsung store key-value aja tanpa penggunaan hash table dimana key itu key untuk cachenya dan valuenya itu json value dari data yang sering diquery.</li>
  <img width="1272" height="720" alt="image" src="https://github.com/user-attachments/assets/1599e098-fc7c-47cf-9c54-e8a2708dcdfb" />
  <li>Atau anggep aja dari key ga langsung ke hash table key tapi langsung ke value.</li>
  <li>Jadi ada multiple key dalam 1 database redis.</li>
</ol>

## Sorted Sets

<ol>
  <li>Sorted Sets adalah struktur data mirip Sets, namun datanya diurutkan sesuai nilai score yang kita tentukan.</li>
  <li>Kalau terdapat score dengan nilai yang sama, maka akan diurutkan secara otomatis (lexicographically).</li>
  <li>Dalam sorted sets, data dikatakan unik jika value dan scorenya sama. Artinya, kita bisa memasukkan value yang sama asal scorenya itu berbeda.</li>
  <li>Tidak seperti Sets yang otomatis reject value yang udah ada di dalam sets.</li>
  <li>Jadi kalau Sorted Sets, value yang sudah ada bisa dimasukkan kembali asalkan nilai dari scorenya berbeda.</li>
  <li>Score bernilai number, dan diurutkan secara ascending (terkecil ke terbesar).</li>
</ol>

#### Rate Limiter

Sekarang gimana implementasinya terkait Sorted Sets untuk Rate Limiter? Nah pertama-tama kita perlu tau dulu apa itu Rate Limiter dan kenapa kita harus pakai itu. Bayangin dengan kasus ada satu user yang spam 200 request dalam 2 detik yang memakan habis resources pada website. Gimana caranya agar services kita bisa stay running & healthy disaat ratusan client hit servicesnya secara concurrent?  
  
<img width="2127" height="1183" alt="image" src="https://github.com/user-attachments/assets/315633af-e6e5-4e98-a009-fda43a8a8dc3" />
  
Nah, solusinya adalah dengan menggunakan konsep yang namanya Rate Limiter. Di Redis Sorted Sets, kita bisa menjadikan **timestamp sebuah request sebagai skor**. Kita tau bahwa Sorted Sets ini akan melakukan sort terhadap sets nya secara otomatis, sehingga solusi dari ini adalah dengan menciptakan sebuah Sliding-window rate limiting.

<img width="1208" height="992" alt="image" src="https://github.com/user-attachments/assets/cdd995c9-c5c9-4606-b89f-8e3382012ffd" />
  
Cara kerjanya:<br/>
<ol>
  <li><b>MASUKKAN TIMESTAMP SEBAGAI VALUE DAN SCORE KE DALAM REDIS SORTED SETS</b>. Dalam kasus ini, kita akan contohkan dengan menggunakan Redis CLI dulu, kedepannya tinggal diimplementasikan ke codingannya.<br/>
    <pre><code>
      ZADD key now now
    </code></pre>
    <br/>
    Contoh:
    <pre><code>
      ZADD user:123:requests now now
    </code></pre>
  </li>
  <li>
    <b>SWEEP THE WINDOW</b>, artinya, semua yang lebih tua dari window akan dihapus. Misal, window = 5000 ms (5 detik), maka semua request yang udah lebih dari 5 detik akan dihapus. Jadi isi sorted sets cuman request dari 5 detik terakhir.<br/>
    <pre><code>
      ZREMRANGEBYSCORE key -inf now-window
    </code></pre>
    <br/>Dalam kasus ini, berarti kita bisa tulis kek gini:
    <pre><code>
      ZREMRANGEBYSCORE user:123:requests -inf now-window
    </code></pre>
  </li>
  <li>
    Appendix: Hitung jumlah request (ZCARD) yang tersisa setelah sweeping
  </li>
  <li>Karena score = waktu. Jadi otomatis “sliding window” ke depan, tanpa butuh counter per user atau cron job buat hapus data lama → data lama auto-hapus tiap request masuk</li>
</ol>  

Ini bentuk implementasi dengan menggunakan Golang:
```go
package util

import (
	"context"
	"fmt"
	"golang-clean-architecture/internal/model"
	"time"

	"github.com/redis/go-redis/v9"
)

type RateLimiterUtil struct {
	Redis      *redis.Client
	MaxRequest int64
	Duration   time.Duration
}

func NewRateLimiterUtil(redis *redis.Client) *RateLimiterUtil {
	return &RateLimiterUtil{
		Redis:      redis,
		MaxRequest: 1,
		Duration:   time.Second * 1,
	}
}

func (u RateLimiterUtil) IsAllowed(ctx context.Context, auth *model.Auth) bool {
	key := auth.ID

	increment, err := u.Redis.Incr(ctx, key).Result()
	if err != nil {
		fmt.Println("Error incrementing:", err)
		return false
	}

	if increment == 1 {
		err := u.Redis.Expire(ctx, key, u.Duration).Err()
		if err != nil {
			fmt.Println("Error setting expiration:", err)
			return false
		}
	}

	return increment <= u.MaxRequest
}
```

Penjelasan:
<ol>
  <li>
    <b>Key = auth.ID</b>
    <br/>
    Jadi limit dihitung per user (ID tertentu).
  </li>
  <li>
    <b>Increment counter</b>
    <br/>
    <pre><code>
      increment, err := u.Redis.Incr(ctx, key).Result()
    </code></pre>
    <ol>
      <li>Tiap request → counter naik.</li>
      <li>Redis INCR aman karena atomic (nggak ke-race walau ada banyak request bersamaan).</li>
    </ol>
  </li>
  <li>
    <b>Set TTL (expire) ketika baru pertama kali dibuat</b>
    <br/>
    <pre><code>
      if increment == 1 {
        u.Redis.Expire(ctx, key, u.Duration)
      }
    </code></pre>
    <ol>
      <li>Kalau baru pertama kali request di window ini → kasih batas waktu sebesar u.Duration, kalau dicode udah kita set = 1 detik. Jadi setelah 1 detik, increment akan jadi 0 lagi.</li>
      <li>Setelah expire, counter auto reset.</li>
    </ol>
  </li>
  <li>
    return increment <= u.MaxRequest (cek limit)
  </li>
</ol>

Nah, cara diatas adalah cara paling sederhana untuk menerapkan Rate Limiter menggunakan Redis. Hanya saja terdapat flaw yaitu, karena kita tidak menggunakan Sorted Set, cara diatas punya kelemahan dimana masih ada kesempatan untuk Hacker untuk melakukan bom request dalam rentang waktu 1 detik tersebut.
Contoh:  
- limit = 5 req / 1 detik.
- User bisa kirim 5 req di detik ke-0.9 dan 5 req lagi di detik ke-1.0 → total 10 req dalam 0.1 detik.
  
Nah, jadi solusi paling ideal adalah dengan kode seperti ini (memanfaatkan command ZADD, ZREMRANGEBYSCORE, dan ZCARD di Golang):
```go
package util

import (
	"context"
	"fmt"
	"golang-clean-architecture/internal/model"
	"time"

	"github.com/redis/go-redis/v9"
)

type SlidingWindowRateLimiter struct {
	Redis      *redis.Client
	MaxRequest int64
	Duration   time.Duration
}

func NewSlidingWindowRateLimiter(redis *redis.Client) *SlidingWindowRateLimiter {
	return &SlidingWindowRateLimiter{
		Redis:      redis,
		MaxRequest: 5,                 // misal 5 request
		Duration:   time.Second * 5,   // dalam 5 detik
	}
}

func (u SlidingWindowRateLimiter) IsAllowed(ctx context.Context, auth *model.Auth) bool {
	key := auth.ID
	now := time.Now().UnixMilli()
	windowStart := now - u.Duration.Milliseconds()

//     0. anggap ini sebagai `multi` keyword yang menandakan awal mula transaction di start dalam Redis
	pipe := u.Redis.TxPipeline()

//     1. tambahkan request sekarang ke sorted set (ZADD key now now)
	pipe.ZAdd(ctx, key, redis.Z{
		Score:  float64(now),
		Member: now,
	})

//     2. hapus request yang sudah di luar window (ZREMRANGEBYSCORE key -inf now-window)
	pipe.ZRemRangeByScore(ctx, key, "-inf", windowStart)

//     3. hitung jumlah request yang tersisa dalam window (ZCARD key)
	count := pipe.ZCard(ctx, key)

//     4. appendix: kasih key expired agar tidak ada key-value yang disimpan di Redis secara permanen.
	pipe.Expire(ctx, key, u.Duration*2)

//     5. Jalankan pipeline transaction (`multi exec`)
	_, err := pipe.Exec(ctx)
	if err != nil {
		fmt.Println("Redis error:", err)
		return false
	}

//     6. ambil hasil count
	reqCount, _ := count.Result()

//     7. true kalau masih di bawah limit
	return reqCount <= u.MaxRequest
}
```
Jadi, tetap aja susunan flownya itu selama dah paham CLInya sedikit-sedikit adalah:
1. ZADD key
2. ZREMRANGEBYSCORE key -inf now-window
3. ZCARD key
  
Command-command penting lainnya bisa di cek disini: https://redis.io/commands/?group=sorted-set

## Visualizations Checkpoints

<img width="876" height="392" alt="image" src="https://github.com/user-attachments/assets/cecfff9e-b098-46e1-b4dc-3a378e0c28db" />

## Streams
