# Redis PubSub

## Apa itu PubSub?

<ol>
  <li>Saat kita membuat aplikasi, salah satu yang biasa dilakukan adalah integrasi antar aplikasinya.</li>
  <li>Saat kita membutuhkan data dari aplikasi lain, maka kita akan melakukan integrasi antar aplikasi.</li>
  <li>Misal kita punya aplikasi Order Management, lalu Order Management ini butuh data dari aplikasi User Management.</li>
  <li>Nah, disitu kita butuh integrasi antar aplikasinya.</li>
  <li>Ada banyak cara untuk melakukan integrasi antar aplikasi, mulai dari sharing file (yang paling basic, jadi dari aplikasi A simpan data ke file, dan aplikasi B tinggal baca file itu), atau dari sharing database (aplikasi A simpan data ke database, aplikasi B tinggal baca databasenya), atau menggunakan API atau messaging (Kafka/RabbitMQ).</li>
  <li>Redis Pubsub itu salah satu integrasi antar aplikasi menggunakan messaging.</li>
</ol>

#### Permasalahan di API untuk integrasi aplikasi

<ol>
  <li>Terkadang, ada kasus dimana kita perlu mengirim data ke lebih dari 1 aplikasi.</li>
  <li>Contoh, kita baru register di aplikasi User Service, nah kita harus mengirimkan data dari User yang berhasil di Register ke Service" seperti Notification Service, Email Service, dsbnya.</li>
  <li>Kalau menggunakan API, data yang dikirimkan harus dilakukan manual secara satu per satu.</li>
  <li>Semakin banyak aplikasinya, semakin lama prosesnya.</li>
  <li>Ini jelas scalability issues, oleh karena itu, menggunakan messaging jauh lebih cocok ketimbang menggunakan API.</li>
</ol>

<img width="1566" height="639" alt="image" src="https://github.com/user-attachments/assets/924fe75e-9fc5-407f-8402-aaced62ded47" />

<ol>
  <li>Dengan menggunakan API, kita harus mengirim data dari client sebanyak n kali (dimana n adalah banyaknya aplikasi).</li>
  <li>Kalau menggunakan messaging, kita akan menggunakan aplikasi tambahan di tengahnya, jadi client cukup mengirim data sebanyak 1 kali ke message brokernya (aplikasi di tengah), nanti aplikasi yang ditengah yang akan DIBACA oleh aplikasi yang membutuhkan datanya.</li>
  <li>Jadi tidak ada kasus pengiriman data sebanyak n kali antar aplikasi.</li>
  <li>Ditambah lagi kalau kebutuhan aplikasinya tidak peduli terhadap berapa lama pemrosesan data yang dikirim.</li>
  <li>Kita bisa mengirimkannya ke dalam Message Broker karena pemrosesan data tiap aplikasi yang membaca ke Message Broker  tersebut berjalan secara <b>Asynchronous</b>.</li>
</ol>

## Redis PubSub

<ol>
  <li>Redis PubSub itu salah satu fitur di Redis yang digunakan untuk implementasi integrasi antar aplikasi menggunakan messaging.</li>
  <li>Implementasi ini menggunakan teknik khusus yaitu teknik Publish-Subscribe.</li>
  <li>Cara kerja Redis PubSub ini sedikit berbeda dari aplikasi PubSub pada umumnya. (ga kek Kafka)</li>
  <li>Redis PubSub (Publisher) hanya akan mengirim data ke Message Broker hanya jika ada Consumer/Subscriber.</li>
  <li>Kalau ternyata tidak ada Consumer/Subscriber yang terhubung ke Message Broker, maka Redis tidak akan mengirimkan data apapun ke dalam Message Broker tersebut, resulting in a data loss if all consumer/subscriber are down.</li>
  <li>Oleh karena itu, disarankan untuk tidak menggunakan Redis PubSub sebagai media penyimpanan/antrian.</li>
</ol>

<img width="947" height="964" alt="image" src="https://github.com/user-attachments/assets/c8cee3c7-9d6f-41d3-b55f-c9e0f763d64c" />

#### Database Scoping

<ol>
  <li>Fitur Redis PubSub tidak seperti struktur data Redis lainnya. Dimana kalau dalam struktur data Redis, scope/lokasi data terdapat dalam database.</li>
  <li>Di Data Structure Redis biasa, misal data kita kirim ke database 0, maka data tersebut hanya bisa dibaca dari database 0 itu lagi kan.</li>
  <li>Nah, kalau Redis PubSub, itu tidak terikat dengan database, oleh karena itu, jika kita mengirim data dari database 0, lalu data diterima di database 7, itu bisa2 aja.</li>
  <li>Jadi PubSub itu sifatnya Global dan bisa diakses dari database manapun.</li>
</ol>

## Channel

<ol>
  <li>Channel adalah sebutan untuk key yang digunakan di Redis PubSub.</li>
  <li>Channel/Key disini digunakan untuk mengirim dan menerima data di PubSub.</li>
  <li></li>
</ol>
