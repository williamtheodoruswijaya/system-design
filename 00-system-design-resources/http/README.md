# HTTP

## Apa itu HTTP Server?

<ol>
  <li>HTTP singkatan dari Hypertext Transfer Protocol.</li>
  <li>HTTP merupakan protokol untuk melakukan transmisi hypermedia document (maksudnya itu pengiriman dari satu komputer ke komputer yang lain), seperti HTML, Javascript, CSS, Image, Audio, Video, dsbnya.</li>
  <li>HTTP awalnya di desain untuk komunikasi antara Web Browser dengan Web Server, namun saat ini sering juga digunakan untuk kebutuhan lain seperti komunikasi antara Aplikasi Mobile dengan Backend Server aplikasi tersebut.</li>
</ol>

### Client Server

<ol>
  <li>HTTP mengikuti arsitektur client dan server.</li>
  <li>Client mengirimkan HTTP request untuk meminta atau mengirim informasi ke Server.</li>
  <li>Dan server akan membalasnya dengan HTTP response dari HTTP request yang diterima.</li>
</ol>

<img width="1015" height="195" alt="image" src="https://github.com/user-attachments/assets/be7e0ce3-8ee2-426f-9ec1-16ee4b3627f9" />

### Plain Language and Human Readable

Nah, kenapa HTTP sangat populer? Karena HTTP didesain menggunakan bahasa yang mudah dimengerti oleh bahasa manusia, seperti:

<ol>
  <li>GET</li>
  <li>POST</li>
  <li>PUT</li>
  <li>DELETE</li>
  <li>HEAD</li>
  <li>OPTION</li>
</ol>

### Stateless

<ol>
  <li>HTTP merupakan protokol yang Stateless.</li>
  <li>Artinya, tiap HTTP Request merupakan request yang independen, tidak ada keterkaitan atau hubungan dengan HTTP Request sebelum atau setelahnya.</li>
  <li>Hal ini dilakukan agar HTTP Request tidak harus dilakukan dalam sequence, sehingga client bisa melakukan HTTP request secara bebas tanpa ada aturan harus dimulai dari mana. (ex. buka Youtube bisa langsung liat komentar, tanpa harus ke dashboard home dulu.)</li>
</ol>

### Session

<ol>
  <li>Tapi kalau emang stateless, tapi kenapa kadang harus login dulu kalau mau ke home facebook?</li>
  <li>Untuk menangani permasalahan seperti ini, ada yang namanya Session, dimana ketika User Login, maka session akan dibuat dan digunakan fitur yang bernama HTTP Cookie.</li>
  <li>HTTP Cookie akan memaksa client untuk menyimpan informasi yang diberikan oleh server ke dalam yang namanya Cookie.</li>
</ol>

## HTTP Version

<ol>
  <li>Versi dari HTTP selalu diperbaharui.</li>
  <li>Anggep HTTP sebagai aturan aja yang harus diikuti aplikasi seperti Google Chrome, dsbnya.</li>
  <li>Makanya, Google Chrome bisa buka sebuah website.</li>
  <li>Saat ini kebanyakan web berjalan di HTTP/1.1 atau HTTP/2.</li>
</ol>

### HTTP/1.1 vs HTTP/2

<img width="1076" height="520" alt="image" src="https://github.com/user-attachments/assets/721f3979-ce21-4b1a-bc3a-976adedf734e" />

<ol>
  <li>Saat ini HTTP/1.1 merupakan fallback protocol (kalau tidak support 2, maka baru pindah kesini), dimana Web Browser secara default akan melakukan request menggunakan HTTP/2, jika web server tidak mendukung, maka web browser akan melakukan fallback ke protocol HTTP/1.1</li>
  <li>Secara garis besar, spesifikasi HTTP/2 sama dengan HTTP/1.1, yang membedakan adalah pada HTTP/2, HTTP Request yang dikirim dalam bentuk teks, secara otomatis menjadi binary, sehingga lebih cepat dibandingkan HTTP/1.1</li>
  <li>Selain itu di HTTP/2, menggunakan algoritma kompresi untuk memperkecil size request dan mendukung multiplexing, sehingga bisa mengirim beberapa request dalam 1 connection yang sama.</li>
  <li>Jadi kalau dari client dan server, ada yang namanya Connection, nah, biasanya HTTP 1.1 dari 1 connection cuman bisa mengirim 1 request, nah kalau HTTP/2 dari 1 connection tersebut bisa mengirim lebih dari 1 request.</li>
  <li>Jadi bedanya gaada di aturan, hanya ada di advantages dimana HTTP/2 ada algoritma kompresi + mendukung multiplexing</li>
</ol>

### HTTPS

<ol>
  <li>Secara default, HTTP itu tidak aman.</li>
  <li>Artinya, kalau ada orang (man in the middle), contohnya dengan menggunakan internet umum, nah orang dengan menggunakan internet tersebut bisa aja mencuri informasi dari internet tersebut.</li>
  <li>Nah, kalau HTTPS menggunakan enkripsi.</li>
  <li>Perbedaan HTTP dengan HTTPS, HTTPS menggunakan SSL (Secure Sockets Layer) untuk melakukan enkripsi HTTP Request dan Response.</li>
  <li>Hasilnya, HTTPS jauh lebih aman dibandingkan HTTP biasa.</li>
  <li>Jadi dari Web Browser, datanya di enkripsi dulu, baru dikirim ke web server, tar baru di decrpyt di web server-nya.</li>
  <li>Jadi otomatis kalau ada man in the middle mencoba untuk mencuri informasi request/response kita, dia tidak bisa melihat isinya.</li>
  <li>Web yang menggunakan HTTPS akan menggunakan HTTPS:// pada urlnya, dan yang hanya menggunakan HTTP tanpa enkripsi, akan menggunakan HTTP://</li>
</ol>
