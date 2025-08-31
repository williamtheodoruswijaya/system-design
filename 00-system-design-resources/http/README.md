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
