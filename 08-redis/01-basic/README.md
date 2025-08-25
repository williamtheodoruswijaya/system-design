# Redis

## Apa itu Redis?

<ol>
  <li>Redis sendiri adalah database berbentuk <b>hash-map</b> yaitu database berbasis key-value (associative array) dan menyimpan datanya di RAM.</li>
  <li>https://redis.io</li>
</ol>

### Apa itu Key-Value Database?

<ol>
  <li>Sebelumnya kita sering mendengar adanya database seperti Relational Database seperti SQL, tapi ada juga yang NoSQL seperti contohnya MongoDB yang berbasis Document.</li>
  <li>Redis sendiri termasuk ke NoSQL yang berbasis key-value.</li>
  <li>Key-value adalah sistem dimana data disimpan dalam bentuk pair (key - value).</li>
  <li>Key mirip dengan primary key dalam database, sedangkan value adalah isi dari datanya.</li>
  <li>Kita bisa mencari data di redis dengan menggunakan keynya.</li>

  <img width="304" height="349" alt="image" src="https://github.com/user-attachments/assets/25b737e4-212f-4507-bc23-1116118aca85" />

  <li>Intinya kita cuman bisa mengambil data yang disimpan di dalam Redis dengan menggunakan Key-nya.</li>
</ol>

### Apa itu Memory Database? Kenapa kita simpan datanya di Memory? Kenapa ga di Hardisk?

<ol>
  <li>Redis itu menyimpan data-nya di Memory (RAM), tapi kita bisa minta datanya secara permanen via Disk.</li>
  <li>Ingat tapi setiap kali kita mengambil data di Redis, itu tetap datanya di ambil dari memory.</li>
  <li>Jadi kapasitas Redis sendiri tetap menggunakan size dari RAM.</li>
  <li>Biasanya data di Disk dijadikan Backup oleh Redis agar Redis bisa restart tanpa kehilangan data yang sempat disimpan.</li>
  
  <img width="301" height="293" alt="image" src="https://github.com/user-attachments/assets/9ad0574f-86ce-4d98-bc3b-109c0ec495c1" />

</ol>

### Kapan Butuh Redis?

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

### Cara Install

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
</ol>

#### Operasi Data String

| Operasi            | Keterangan                          |
|--------------------|-------------------------------------|
| `set key value`    | mengubah string value dari key      |
| `get key`          | mendapatkan value menggunakan key   |
| `exists key`       | mengecek apakah key memiliki value  |
| `del key [key ...]`| menghapus menggunakan key           |
| `append key value` | menambah data value ke key          |
| `keys pattern`     | mencari key menggunakan patterns    |
