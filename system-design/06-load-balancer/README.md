# Load Balancer

## Apa itu Load Balancer?

<img width="861" height="361" alt="image" src="https://github.com/user-attachments/assets/feedd938-43d6-4cd0-a451-b8ea39eb1171" />

Fungsi utama:<br/>
Distribute incoming network traffic across multiple servers ensuring high availability and **scalability**.

### Q: Gimana Load Balancer bisa improve Scalability?

<img width="1385" height="362" alt="image" src="https://github.com/user-attachments/assets/3c8cfa7c-8c3b-4dc2-b629-6ff2c15150ec" />

#### Conclusions:

Intinya, kita bisa pakai Load Balancer ke semua network distributions, bahkan proses read/write ke database yang punya banyak replication.

### Q: Why'd we need Load Balancer? Why not make it into only one server?

Server can be overworked, which can reduce performance. Therefore, we need multiple server, **but how'd we distribute the request?** **USE LOAD BALANCER!!**

## Routing Algorithms

<img width="3735" height="3573" alt="0251-lb-algorithms" src="https://github.com/user-attachments/assets/531174ed-858b-4cee-afee-8f26a4883290" />

<ol>
  <li>
    <b>Round Robin</b><br/>
    Distribute requests in rotation (abis ke server 1, ke server 2, ke server 3, dst...)
  </li>
  <li>
    <b>Sticky Round Robin</b><br/>
    Ini improvement dari Round Robin biasa, kalau Request dari Alice pertama-tama udah diarahkan ke server 1, maka semua request selanjutnya dari Alice akan diteruskan ke server 1.
  </li>
  <li>
    <b>Weighted Round Robin</b><br/>
    Assign Weight to each server (Round Robin + If command to check weight for server where weight = how many request one server can handle)
  </li>
  <li>
    <b>Hashing: hash client ip address.</b>
  </li>
  <li>
    <b>Least Connections</b><br/>
    Request will be sent to server with fewest connections.
  </li>
  <li>
    <b>Least Response Time</b><br/>
    Request will be sent to server with least response time.<br/>
    Kita tau kalau semakin banyak request terhadap sebuah server, semakin berat performance-nya. Artinya, response time-nya juga meningkat. Nah algoritma ini akan mendistribusikan request ke sever dengan response time tercepat.
  </li>
</ol>

## Disadvantages of Load Balancer

If a Load Balancer fails (single point of failure), **that's why we have redundant load balancer (n load balancer).**<br/>
Jadi sebenarnya, diagram Load balancer kita gambarin kayak begini:

<img width="995" height="492" alt="image" src="https://github.com/user-attachments/assets/3d886d5a-9d3f-4543-93ce-c5659f002221" />

ex.) Load Balancer:

- Azure Load Balancing
- Amazon Elastic Load Balancing
- Digital Ocean
- Nginx

## Advantages of Load Balancer

<ol>
  <li><b>Autoscaling:</b> Add/Remove instances automatically.</li>
  <li><b>Encryption:</b> Encrypted Connection (SSL) -> Kalau pake Load Balancer bisa jadi https dari http kek pas pake AWS.</li>
  <li><b>Health Check:</b> Deteksi server mati/engga (periodically) otomatis matiin server, terus spawn new server.</li>
  <li><b>Compression:</b> Compress size biar ga gede.</li>
  <li><b>Sticky Session:</b> Assign same user/device to same resources.</li>
  <li><b>Persistence Connections (Websocket):</b> allow server to create persistent connection to client.</li>
</ol>

# Consistent Hashing

Misal ketika kita pake Amazon EC2 Auto-Scale, nah jumlah setiap server kan bisa berubah-ubah (bisa nambah/berkurang) sendiri tergantung traffic sebuah system. Nah, gimana distribusi requestnya agar bisa tetap efficient, common approachnya kan menggunakan hash dari request tersebut `Hash(key) mod N` dimana `N` adalah jumlah server tersebut. Metode ini oke, tapi begitu jumlah server berubah, maka akan terjadi redistribusi request secara besar-besaran. Consistent Hashing ini adalah metode buat handle problem ini dimana dia ini membuat ketika terjadi sebuah perubahan terhadap jumlah server, redistribusi request hanya terjadi pada sedikit request saja.

## Traditional Hashing Problem

Ketika kita membuat high-traffic web application yang handle millions of user, kita akan membuat multiple server dan menggunakan load balancer untuk distribusi request tersebut (anggep aja hash-based load balancer). Misal, sistem kita terdiri dari 5 backend server (S0, S1, S2, S3, dan S4) dan setiap request diarahkan menggunakan hash function.

<img width="1498" height="1422" alt="image" src="https://github.com/user-attachments/assets/ddd87915-4bd2-419e-b625-541633fa6e37" />

Proses-nya kurang lebih kayak begini:
1. Load Balancer mengambil IP-Address dari User (atau Session ID).
2. Hash Function akan maps IP dengan menjumlahkan banyaknya Byte dalam IP tersebut yang kemudian di modulo oleh jumlah server yaitu 5.
3. Hasil angkanya akan menunjukkan server yang harus ditunjuk.

<img width="2114" height="1148" alt="image" src="https://github.com/user-attachments/assets/8ef44aa8-6f30-4f40-842e-f2e04cd591d8" />

Sementara semua akan baik-baik aja, sampai harus di-scale. Misal, kita tambahin S5. Otomatis, hash function kita juga akan berubah dan menyebabkan redistribusi requestnya jadi seperti ini:

<img width="2118" height="1290" alt="image" src="https://github.com/user-attachments/assets/ade81a05-3e47-44f2-9f25-4e3d30255735" />

Darisini aja requestnya jadi terdistribusi secara ga imbang (redistribusi bukan distribusi request baru) artinya request lama ada yang berubah handlingnya. Sekarang, gimana kalau S4 fails dan dihapus.

<img width="2116" height="1124" alt="image" src="https://github.com/user-attachments/assets/1edb3165-6b82-4d31-9d39-43b0273f1ed7" />

Maka, sebagian besar user akan mengalami redistribusi request. Ini fatal karena bisa menyebabkan session loss (kalau stateful architecture, user akan logout setiap kali ada perubahan jumlah server), cache juga akan invalidated, sehingga akan terjadi massive traffic ke database, dan akan terjadi performance degradation pada seluruh aplikasi. 

## Consistent Hashing

Dengan Consistent Hashing, kita bisa membuat redistribusi request tersebut hanya terjadi pada beberapa user aja (ga massive kek sebelumnya). Alih-alih menempatkan server dalam posisi vertical, ini kita tempatkan dalam yang namanya **circular hash space (hash ring)** dengan besar lingkaran yang konstan. Contoh kalau ada 5 server, bentuknya jadi kek gini:

<img width="1536" height="1362" alt="image" src="https://github.com/user-attachments/assets/9b34e8f9-fe81-4f36-baaa-73de86568b4c" />

Nah, nanti cara kerja mapnya itu posisi dari request akan ditempatkan berdasarkan hash function sebelumnya. Jadi bukan request yang diarahkan ke server mana tapi posisi request buat ditempatkan di circular hash tersebut.

<img width="1636" height="1362" alt="image" src="https://github.com/user-attachments/assets/6c97faa2-752b-492e-bc57-acba603a4e3c" />

Request tersebut akan diarahkan ke server terdekat secara clockwise. Dan ketika ada server baru, hanya beberapa request aja yang akan didistribusikan ke server baru.

<img width="1700" height="1364" alt="image" src="https://github.com/user-attachments/assets/5dac7552-3f71-46dc-b8c4-ad2d83580dcd" />

Apalagi kalau kita hapus, maka hanya segelintir request aja yang merasakan effectnya dan diarahkan ke server baru.

<img width="1756" height="1362" alt="image" src="https://github.com/user-attachments/assets/cbd4237c-bbfd-4288-95c4-6668b62ef380" />

Tapi kalau kek gini doang, ada problem dimana bisa aja server-server yang dihash, memiliki posisi clustered yang menyebabkan adanya hot spots. Dan ketika server di hapus, akan terjadi shifting secara besar-besaran pada request karena ada-nya hotspot tersebut. Makanya ada yang namanya **Virtual Nodes**. Jadi, alih-alih mengassign server berdasarkan posisi secara 1 aja:

<img width="1660" height="1338" alt="image" src="https://github.com/user-attachments/assets/31a05005-0fd1-4303-94b2-c0710d484641" />

Kita hash server-server yang udah di hash secara berkali-kali:

<img width="1838" height="1390" alt="image" src="https://github.com/user-attachments/assets/5474f799-7bac-40f1-acbf-ba123e483740" />

Dengan begitu, ketika ada perubahan jumlah server, redistribusi ke S2 dan S3 tidak akan menyebabkan massive shifting.

> Sources: https://blog.algomaster.io/p/consistent-hashing-explained
