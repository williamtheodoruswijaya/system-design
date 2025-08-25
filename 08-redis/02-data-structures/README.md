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

## Sorted Sets (Rate Limiter)

<ol>
  <li>Sorted Sets adalah struktur data mirip Sets, namun datanya diurutkan sesuai nilai score yang kita tentukan.</li>
  <li>Kalau terdapat score dengan nilai yang sama, maka akan diurutkan secara otomatis (lexicographically).</li>
  <li>Dalam sorted sets, data dikatakan unik jika value dan scorenya sama. Artinya, kita bisa memasukkan value yang sama asal scorenya itu berbeda.</li>
  <li>Tidak seperti Sets yang otomatis reject value yang udah ada di dalam sets.</li>
  <li>Jadi kalau Sorted Sets, value yang sudah ada bisa dimasukkan kembali asalkan nilai dari scorenya berbeda.</li>
  <li>Score bernilai number, dan diurutkan secara ascending (terkecil ke terbesar).</li>
</ol>

## Conclusion Visualizations

<img width="876" height="392" alt="image" src="https://github.com/user-attachments/assets/cecfff9e-b098-46e1-b4dc-3a378e0c28db" />
