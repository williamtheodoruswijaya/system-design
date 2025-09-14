# CAP Theorem

<img width="543" height="535" alt="image" src="https://github.com/user-attachments/assets/2faef11b-2424-4fd1-bbf5-26fc8744271d" />

Ini adalah sebuah theorem yang mengatakan trade-off antara Consistency, Availability, dan Partition Tolerance. CAP Theorem basically says:<br/>
> It is impossible for a distributed data store to simultaneously provide all three guarantees.

* **Consistency (C)**: Setiap `READ` mendapatkan data terbaru yang baru di `WRITE` atau dapat sebuah `ERROR`.
* **Availability (A)**: Setiap request (baik itu `READ` ataupun `WRITE`) tidak boleh mendapatkan error, tanpa kepastian bahwa data yang di `READ` adalah data terbaru (most recent `WRITE`).
* **Partition Tolerance (P)**: System harus terus bekerja meskipun terdapat **failure of communication** antara server.

## 3 Pillars of CAP

### 1. Consistency

Sebuah system dikatakan consistent jika **setiap read yang dilakukan, mengembalikan response yang merupakan the most recent write ke database atau kembalikan sebuah error.** Artinya semua nodes yang ada dalam distributed system (dalam kasus ini server-server yang ada) akan mengembalikan data yang sama dalam waktu yang sama pula.

<img width="1166" height="932" alt="image" src="https://github.com/user-attachments/assets/6b4486c0-567d-4931-bf98-3537c2a0e6be" />

Jadi, kalau kita `WRITE` data ke server A, semua `READ` dari server lain secara langsung mengembalikan data yang barusan di `WRITE` ke server A. Consistency ini cocok dan wajib diterapkan pada aplikasi yang mengharuskan data yang diterima adalah data yang most up-to-date, seperti contohnya financial systems, dimana ketika terjadi perubahan balance pada salah satu akun yang ditulis menggunakan 1 server, semua server lain yang READ akun tersebut harus mengembalikan nilai balance yang baru saja berubah.

### 2. Availability

Availability mengharuskan sebuah sistem yang menerima Request (read/write) harus mengembalikan sebuah response (ga boleh error), tapi dengan demikian ga mungkin bisa mengembalikan data yang merupakan most recent write. (Ga mungkin achieve consistency sambil mempertahankan availability). Jadi, sebuah sistem wajib responsif dan operasional meskipun data-nya tidak pernah up-to-date.

<img width="1188" height="868" alt="image" src="https://github.com/user-attachments/assets/baafdaed-01df-4788-b27e-ace7c7657f0b" />

Availability penting untuk aplikasi yang harus tetap operational sepanjang waktu (Online Retail System).

### 3. Partition Tolerance

Partition Tolerance artinya sebuah sistem harus tetap bekerja terlepas jika terdapat failure dalam komunikasi antar node (dalam kasus distributed server biasanya antar server yang harus tetap sync terjadi failure communication).

<img width="844" height="858" alt="image" src="https://github.com/user-attachments/assets/81a0506c-3485-41b4-b31e-204aa8fae2f9" />

Ketika sebuah network failure terjadi, akan terjadi sebuah **network partition** yaitu **sebuah distributed system akan terbagi menjadi 2 grup yang tidak bisa berkomunikasi satu sama lain**. Partition tolerance basically menciptakan sebuah server yang tidak pernah terjadi Network Partition dan tetap dapat operasional antar partition.

## CAP Trade-off

Network Partition yang terjadi karena Network Failure akan mengharuskan sebuah sistem harus memilih salah satu antara **Consistency** dan **Availability** dan tidak bisa keduanya. Contoh skenario:

### CP (Consistency and Partition Tolerance):

Sistem ini memilih Consistency ketika sebuah Network Partition terjadi, dengan harga Availability yang berkurang. Artinya, ketika sebuah Network Partition terjadi, sistem bisa melakukan reject ke beberapa request demi mempertahankan consistency antar node. Contoh:<br/>
> Banking system akan memprioritaskan consistency over availability, jadi ketika ATM ditarik uangnya, system harus bisa mengupdate semua balance yang ada di ATM mana pun saat itu juga.

Database yang memilih trade-off ini adalah MySQL dan PostgreSQL, dimana karena database sejenis ini lebih sering digunakan untuk proses READ yang lebih banyak ketimbang WRITE (cmiiw).

### AP (Availability and Partition Tolerance):

Sistem ini memilih Availability, jadi ketika sebuah Network Partition terjadi, beberapa node bisa aja mengembalikan value yang berbeda. Contoh:
> Amazon's shopping cart system selalu di desain untuk bisa kita masukkan item (ga pernah ada kasus Amazon's shopping cart tidak bisa menerima item bahkan ketika high traffic seperti 9.9 sales, Amazon's shopping cart tetap bisa kita input barang).

Database yang memilih trade-off seperti ini adalah NoSQL database seperti Caassandra dan DynamoDB yang lebih condong ke Heavy-Write meskipun dia jadinya ada peluang buat menampilkan data yang tidak konsisten antar node.

### CA (Consistency and Availability):

Sistem bisa aja consistent dan available, tapi ketika sebuah Network Partition terjadi, akan terdapat 2 group of node yang tidak bisa berkomunikasi. Hal ini jelas ga mungkin kecuali sistem kita hanya running di 1 node aja. (Ga pernah di pakai di Reallife).

## Practical Design Strategies

Distributed System harus bisa memanfaatkan trade-off ini, ada beberapa yang diterapkan:

### 1. Eventual Consistency

Untuk beberapa sistem aplikasi, consistency itu bukan hal wajib (kecuali Bank), jadi data-nya dibiarkan aja ga konsisten dulu buat sementara yang penting nanti konsisten.

### 2. Strong Consistency

Nah kalau yang ini, ketika sebuah `WRITE` berhasil, semua `READ` dari node manapun harus mengembalikan value yang baru di `WRITE` tersebut.

### 3. Tunable Consistency

Kalau ini, sejenis sistem yang bisa di atur level of consistencynya tergantung kebutuhan. Dia bisa milih antara Eventual Consistency dengan Strong Consistency seperti aplikasi E-commerce, dimana Order Service harus punya Strong Consistency, sementara Product Recommendation cukup punya Eventual Consistent aja. Ini bisa di achieve dengan database Cassandra makanya Twitter pake database ini buat service follow dan unfollow mereka.

## Beyond CAP: PACELC

CAP itu ga selalu cover semua scenario, ada yang namanya PACELC yang juga merupakan extension dari CAP dengan adanya atribut LATENCY yang juga jadi bagian dari trade-off antara latency dengan consistency.
