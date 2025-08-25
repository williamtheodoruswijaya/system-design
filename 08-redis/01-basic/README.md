# Redis

## Apa itu Redis?

<ol>
  <li>Redis sendiri adalah database berbentuk <b>hash-map</b> yaitu database berbasis key-value (associative array) dan menyimpan datanya di RAM.</li>
  <li>https://redis.io</li>
</ol>

#### Apa itu Key-Value Database?

<ol>
  <li>Sebelumnya kita sering mendengar adanya database seperti Relational Database seperti SQL, tapi ada juga yang NoSQL seperti contohnya MongoDB yang berbasis Document.</li>
  <li>Redis sendiri termasuk ke NoSQL yang berbasis key-value.</li>
  <li>Key-value adalah sistem dimana data disimpan dalam bentuk pair (key - value).</li>
  <li>Key mirip dengan primary key dalam database, sedangkan value adalah isi dari datanya.</li>
  <li>Kita bisa mencari data di redis dengan menggunakan keynya.</li>

  <img width="304" height="349" alt="image" src="https://github.com/user-attachments/assets/25b737e4-212f-4507-bc23-1116118aca85" />

  <li>Intinya kita cuman bisa mengambil data yang disimpan di dalam Redis dengan menggunakan Key-nya.</li>
</ol>

#### Apa itu Memory Database? Kenapa kita simpan datanya di Memory? Kenapa ga di Hardisk?

<ol>
  <li>Redis itu menyimpan data-nya di Memory (RAM), tapi kita bisa minta datanya secara permanen via Disk.</li>
  <li>Ingat tapi setiap kali kita mengambil data di Redis, itu tetap datanya di ambil dari memory.</li>
  <li>Jadi kapasitas Redis sendiri tetap menggunakan size dari RAM.</li>
  <li>Biasanya data di Disk dijadikan Backup oleh Redis agar Redis bisa restart tanpa kehilangan data yang sempat disimpan.</li>
  
  <img width="301" height="293" alt="image" src="https://github.com/user-attachments/assets/9ad0574f-86ce-4d98-bc3b-109c0ec495c1" />

</ol>

#### Kapan Butuh Redis?

<ol>
  <li>Saat kita membuat aplikasi secara pertama kali, itu biasanya kita ga langsung pakai Redis.</li>
  <li>Redis biasa digunakan ketika ada kasus tertentu.</li>
  <li>Implementasi Redis biasanya cenderung mahal karena harga Memory itu lebih mahal ketimbang Disk.</li>
  <li>Untuk menggunakan Redis, kita harus melihat kasusnya secara detail.</li>
  <li>Contoh:
    <ul>
      <li>
        Ketika Database Utama Lambat<br/>
        <img width="465" height="254" alt="image" src="https://github.com/user-attachments/assets/2a41ff1d-bb20-432b-9fe6-bdaf5de9fade" /><br/>
        Biasa ketika Database lambat, kita akan menyimpan data yang paling sering diambil/diread, untuk disimpan di cache, yaitu di Redis.
      </li>
      <li>
        Dan sebagainya:<br/>
        <img width="282" height="386" alt="image" src="https://github.com/user-attachments/assets/5f604a6b-f6cf-4c83-a5a1-0755a22712d4" /><br/>
        Bisa juga Redis kita pakai sebagai Rate Limiter.
      </li>
    </ul>
  </li>
</ol>

#### Cara Install

<ol>
  <li>Link: <b>https://redis.io/docs/getting-started/installation</b></li>
  <li>Saat kita menginstall Redis, ada 2 aplikasi yang sebenarnya terinstal, yaitu <b>Redis Server dan Redis CLI</b>.</li>
  <li>Redis server adalah aplikasi Redis itu sendiri. Jadi, Redis berjalan berbasis Client-Server, dimana kita harus menjalankan Server terlebih dahulu dan kita akan menggunakan Redis Client untuk bisa berkomunikasi ke server.</li>
  <li>Redis CLI adalah aplikasi command line untuk client, dimana client lah yang akan berkomunikasi dengan server.</li>
  
  <img width="697" height="256" alt="image" src="https://github.com/user-attachments/assets/285c838c-e3f1-4a42-b95b-25cdef18a4d1" />

  <li>Cara cek sebenernya tinggal ketik <b>redis-server</b> buat menjalankan server redisnya secara manual (atau bisa otomatis via docker).</li>
  <li>Saya sendiri senengnya via Docker jadi tinggal download aja Imagenya via Docker Registry dengan cara <b>docker pull redis:latest</b></li>
  <li>Cara menjalankan Redis Client berbasis CLI: <b>redis-cli -h localhost -p 6379</b></li>
</ol>

## Database

<ol>
  <li>Redis memiliki konsep database seperti pada relational database mysql atau postgre.</li>
  <li>Di Redis, kita bisa membuat database dan menggunakan databasenya.</li>
  <li>Tapi yang membedakan adalah, kalau relational database, kita bisa membuat database dengan menyebutkan nama databasenya, nah kalau di redis, kita cuman bisa membuat database menggunakan angka.</li>
  <li>Secara default database di redis itu namanya 0.</li>
  <li>Maksimal database di redis secara default itu 16, meskipun kita bisa atur di file confignya.</li>
  <li>Biasanya 1 aplikasi bakal pakai 1 database aja.</li>
  <li>Gimana cara kita pindah-pindah database? well, tinggal pindah aja ke cli-nya dan jalankan <b>select 0</b>, yang artinya kita memilih database nomer 0, kita bisa pindah via nomer ya ges.</li>
</ol>

## Strings

<ol>
  <li>Redis mendukung banyak struktur data, salah satunya itu <b>Strings</b>.</li>
  <li>Tapi yang paling sering digunakan justru adalah <b>Strings</b></li>
  <li>Tapi saya sendiri sering nyimpen `json` data di redis.</li>
</ol>

#### Operasi Data String

| Operasi            | Keterangan                                                                             |
|--------------------|----------------------------------------------------------------------------------------|
| `set key value`    | mengubah string value dari key                                                         |
| `get key`          | mendapatkan value menggunakan key                                                      |
| `exists key`       | mengecek apakah key memiliki value                                                     |
| `del key [key ...]`| menghapus menggunakan key (kalau mau lebih dari satu tinggal kasih spasi aja)          |
| `append key value` | menambah data value ke key (redis mirip hash table, 1 key bisa punya multiple value)   |
| `keys pattern`     | mencari key menggunakan patterns                                                       |
  
#### Operasi Range Data Strings

| Operasi                   | Keterangan                                                                                                                                          |
|----------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------|
| `setrange key offset value`| mengubah value dari offset yang ditentukan, misal mengubah value yang posisinya ada di tengah-tengah (string berfungsi sebagai array of character, offset dianggap sebagai index) |
| `getrange key start end`   | mengambil value dari range yang ditentukan, misal mengambil value dari awal ke tengah                                                               |

#### Operasi Multiple Data Strings

| Operasi                       | Keterangan                                                                                                                |
|-------------------------------|----------------------------------------------------------------------------------------------------------------------------|
| `mget key [key ...]`          | Get the values of all the given keys (kalau kita mau mengambil banyak value dari berbagai key)                            |
| `mset key value [key value ...]` | Set multiple keys to multiple values (atau kalau mau mengubah banyak value dari banyak key secara sekaligus)              |

**Notes:** `[key value ...]` artinya pasangan key dan value dipisahkan dengan spasi.<br/><br/>

Contoh:  
```bash
mset budi "100" joko "200" ucok "300"
```

## Expiration

<ol>
  <li>Secara default, ketika kita menyimpan data ke redis, data tersebut akan terus berada di Redis sampai kita menghapusnya.</li>
  <li>Kadang kita ingin menghapus data di redis secara otomatis dalam kurun waktu tertentu.</li>
  <li>Contoh, kita mau menyimpan data cache tersebut dalam waktu 10 menit saja.</li>
  <li>Setelah lewat 10 menit tersebut, data akan kembali disimpan di redis secara otomatis tetapi melalui query ulang ke database.</li>
  <li>Kita bisa atur waktu expired-nya agar redis bisa menghapus data tersebut kalau sudah lewat waktu expire.</li>
  <li>Contoh: <b>expire ini-key 10</b> artinya key yang bernama "ini-key" akan di delete setelah 10 detik. Ini digunakan ketika datanya sudah terlanjur ada di redis.</li>
  <li>Kalau datanya belum ada di redis tapi kita ingin sekaligus mengatur expire time-nya, kita bisa dengan cara: <b>setex key seconds value</b>, contoh <b>setex ini-key 10 ini-value</b>, maka value dengan key "ini-key" akan dibuat di redis sekaligus akan di hapus setelah 10 detik.</li>
  <li>Kita juga bisa cek time limit yang sebuah key-value miliki menggunakan command <b>ttl</b></li>
</ol>

## Increment & Decrement

<ol>
  <li>Operasi increment/menaikkan angka, atau decrement/menurunkan angka, itu sekilas mudah untuk dilakukan, kita tinggal perlu mengupdate data yang ada di redis dengan data yang baru (data lama ditambah 1).</li>
  <li>Tapi kalau aplikasi kita rame, yang membuat operasi tersebut dilakukan secara paralel, maka hal ini dapat menyebabkan race condition.</li>
  <li>Untungnya, redis memiliki operasi bawaan untuk melakukan increment/decrement yang secara otomatis menghandle race condition tersebut.</li>
</ol>

#### Operasi Increment & Decrement

<ol>
  <li>
    Contoh Kode yang salah: <br/>
    <pre><code class="language-js">
    let value = await redis.get('key');
    value = Number(value) + 1;
    await redis.set('key' value);
    </code></pre>
    Hal ini jelas berbahaya karena bisa aja ada 2 user yang ingin mengupdate value dengan key yang sama.
  </li>
  <li>Cara handlenya sesimple <pre>incr key</pre> dengan syarat value yang disimpan dengan key tersebut harus bersifat angka/integer.</li>
  <li><pre>decr key</pre> juga berfungsi sebaliknya</li>
  <li>Tapi gimana kasusnya kalau kita ingin jumlahnya ga cuman 1 tapi lebih dari 1, misal, 2, 3, dstnya.</li>
  <li>Nah, kita bisa pakai <pre>incrby key jumlah-increment</pre></li>
  <li>Atau kalau mau decrement ya <pre>decrby key jumlah-decrement</pre></li>
</ol>

## Flush

<ol>
  <li>Gimana kalau kita ingin mengosongkan seluruh data di Redis?</li>
  <li>Misal ada error yang mengharuskan kita untuk menghapus seluruh data yang ada di redis.</li>
  <li>Nah, kita bisa pakai operasi flush.</li>
  <li>Buat menghapus seluruh key yang ada di database yang sedang digunakan: <pre>flushdb</pre></li>
  <li>Buat menghapus seluruh key yang ada di seluruh database baik yang sedang digunakan ataupun tidak digunakan: <pre>flushall</pre></li>
</ol>

## Pipeline

<ol>
  <li>Perintah yang dikirimkan dari client ke server, itu menggunakan request/response protocol.</li>
  <li>Artinya setiap request yang dikirim, itu punya response yang bakal dibalas oleh redis secara langsung.</li>
  <li>Gimana kalau kita perlu mengirim banyak data ke redis secara langsung?</li>
  <li>Kalau kita ngirim secara satu per satu akan lambat karena akan mendapat respon secara satu per satu.</li>
  <li>Nah untungnya redis bisa mengirim banyak data secara langsung, metode ini kita kenal dengan nama bulk insert.</li>
  <li>Hal ini bisa diachieve dengan menggunakan operasi pipeline dimana kita bisa mengirim banyak data dalam 1 request.</li>
  <li>Disadvantagesnya adalah server redis tidak akan membuat respon untuk setiap data yang masuk jika dikirim via pipeline.</li>
  <li>Caranya kalau via redis-cli: 
    <pre><code>
      redis-cli -h host -p port -n database --pipe < file
    </pre></code>
  </li>
  <li>
    <img width="516" height="238" alt="image" src="https://github.com/user-attachments/assets/65f1751e-4d17-4c8b-82d5-35ee3246c48f" />
  </li>
</ol>

## Transaction

<ol>
  <li>Seperti database relational, redis memiliki fitur Transaction yaitu proses dimana ketika kita mengirim beberapa perintah, dan perintah tersebut dianggap sukses jika semua perintah yang dikirimkan itu sukses, jika salah satu gagal, maka semua perintah tersebut dibatalkan semuanya.</li>
</ol>

#### Operasi Transaction

| Operasi  | Keterangan                                     |
|----------|------------------------------------------------|
| MULTI    | Mark the start of a transaction block          |
| EXEC     | Execute all commands issued after MULTI        |
| DISCARD  | Discard all commands issued after MULTI        |

<ol>
  <li>Multi disini menandakan awal mulai transaksi. Eksekusi-eksekusi selanjutnya setelah Multi ini tidak akan dimulai langsung.</li>
  <li>Eksekusi dibawahnya akan menunggu dijalankan atau tidak dijalankan dengan command Exec/Discard.</li>
  <li>Contoh:<br/><img width="421" height="238" alt="image" src="https://github.com/user-attachments/assets/a07e6692-3839-4a8d-aecd-8517cf1ff92c" /></li>
  <li>Yang membedakan adalah, ketika kita menuliskan multi, semua command dibawahnya akan dimasukkan ke dalam queue dan tidak akan dijalankan sampai ada command exec.</li>
  <li>Ini yang membedakan dengan transaction pada DB dimana ketika kita memulai transaction, command berikutnya akan dijalankan sampai ketika ada command yang gagal, baru command-command yang sebelumnya akan di rollback.</li>
</ol>

## Monitor

<ol>
  <li>Kadang ada kasus kita ingin melakukan debug aplikasi ketika sedang berkomunikasi dengan redis.</li>
  <li>Misal, kita ingin melihat semua request yang masuk ke redis server.</li>
  <li>Nah, itu kita bisa monitor dengan menggunakan fitur monitoring.</li>
  <li>Dengan fitur ini, kita bisa mendebug ketika ada operasi yang salah ketika ke redis servernya.</li>
  <li>Operasi monitornya itu sesimple: <br/><pre>monitor</pre></li>
</ol>

## Server Information

<ol>
  <li>Kadang kita butuh informasi/statistik dari server redis.</li>
  <li>Seperti jumlah memory yang sudah terpakai, dan konfigurasi dari servernya.</li>
  <li>Cara-nya: <pre>info</pre></li>
  <li>Atau kalau mau dapet informasi statistic secara spesifik kita bisa pakai: <pre>config get key</pre></li>
  <li>Contoh: <pre>config get databases</pre> akan mendapatkan informasi jumlah database yang didapat.</li>
</ol>

## Client Connection

<ol>
  <li>Redis menyimpan semua informasi client di server.</li>
  <li>Hal ini memungkinkan kita untuk melihat daftar client yang terkoneksi, dan bahkan sampai mengecek anomali seperti terlalu banyak koneksi client ke server dsbnya.</li>
  <li>Kita bisa pakai <b>client list</b> untuk mendapatkan list client yang terkoneksi ke server.</li>
  <li>Kita bisa pakai <b>client id</b> untuk mendapatkan id dari client tersebut (ini biasa lanjutannya bakal kita kill)</li>
  <li>Kita bisa pakai <b>client kill ip:port</b> untuk mematikan koneksi client ke server.</li>
</ol>

## Protected Mode

<ol>
  <li>Secara default, ketika kita menyalakan redis server, redis server itu akan mendengarkan request dari semua network interface (misal, ada yang konek dari localhost, VM, dsbnya).</li>
  <li>Hal ini justru berbahaya karena membuat redis server bisa diakses darimanapun.</li>
  <li>Untungnya redis punya second layer protection untuk pengecekan koneksi, yaitu mode protected.</li>
  <li>Secara default, mode protectednya itu aktif yang artinya walaupun redis server bisa diakses darimanapun, tapi redis hanya mau menerima request dari localhost kita.</li>
  <li>Coba aja pake 2 laptop yang ada redis, dah gitu coba konek ke redis server dari laptop yang berbeda dengan ip yang unik dari laptop tersebut, maka redis akan menolak koneksinya.</li>
</ol>

## Security

#### Authentication

<ol>
  <li>Proses verifikasi identitas untuk memastikan bahwa yang mengakses redis server adalah identitas yang benar.</li>
  <li>Redis memiliki fitur authentication, dan kita bisa menambahkan di file konfigurasi di server redis.</li>
  <li>Proses authentication di redis sangat cepat, jadi pastikan password harus sulit biar susah di brute force.</li>
</ol>

#### Authorization

<ol>
  <li>Authorization adalah proses memberi hak akses terhadap identitas yang berhasil melewati proses authentication.</li>
  <li>Pertama kita harus tentukan username dan password untuk authentication.</li>
  <li>Setelah itu kita beri akses control terhadap username dan password tersebut untuk membatasi apa yang bisa username tersebut lakukan.</li>
  <li>Commandnya ada banyak di dokumentasi kalau via redis cli.</li>
  <li>Meskipun kedepannya kita bakal pakai redis client via Go, atau NodeJS, bagian-bagian ini cukup pahami teorinya aja.</li>
</ol>

## Persistence

<ol>
  <li>Media penyimpanan utama di redis itu memory/ram.</li>
  <li>Tapi, kita bisa menyimpan data tersebut ke disk.</li>
  <li>Tapi proses penyimpanannya ini ga realtime, tetapi dia melakukan secara periodic.</li>
  <li>Oleh karena itu, jangan anggap redis sebagai tempat penyimpanan yang persistence karena proses menyimpan datanya yang bertahap.</li>
</ol>

## Eviction

<ol>
  <li>Ketika memory redis penuh, maka redis secara default akan reject semua request ketika kita melakukan penyimpanan data.</li>
  <li>Hal ini mungkin akan menjadi masalah jika kita menggunakan redis sebagai caching saja.</li>
  <li>Artinya data baru tidak akan bisa masuk.</li>
  <li>Kadang akan sangat berguna jika memory penuh, redis bisa secara otomatis menghapus data yang sudah jarang digunakan.</li>
  <li>Fitur ini adalah namanya Eviction dimana dia bisa menghapus data yang sudah mendekati expired.</li>
  <li>Namun untuk mengaktifkan fitur ini, kita perlu memberi tahu redis, maximum memory yang boleh digunakan, dan bagaimana strategi untuk melakukan eviction nya
.</li>
</ol>
