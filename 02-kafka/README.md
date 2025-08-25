# Kafka - Message Broker

## Publish-Subscribe

Saat kita membuat aplikasi, pasti ada komunikasi antar aplikasi dengan mekanisme RPC (Remote Procedure Call). Contohnya kek RESTful API. Keuntungan aplikasi dengan RPC ini tuh bisa memanggil dengan sync dan async.
  
* Contoh Kasus Aplikasi E-Commerce:
  * Dalam aplikasi E-Commerce, ada aplikasi:
  1. Product Services
  2. Promo Services
  3. Shopping Cart Services
  4. Order Services
  5. Logistic Services
  6. Payment Services
  
Untuk mendapatkan informasi product, shopping cart services wajib **berkomunikasi** dengan product services untuk mendapatkan informasi product dan selebihnya.

* Diagram RPC untuk aplikasi E-Commerce:

<img width="1015" height="388" alt="image" src="https://github.com/user-attachments/assets/858af4cb-3a23-4c7e-96e2-5e7734c71d0c" />
  
* Perubahan kasus

<p style="text-align: justify;">
  Misal dalam kasus diatas, kita mau menambahkan fraud detection yang akan dikirimkan setelah order services. Hal ini akan membuat aplikasi kita tidak scalable karena membuat aplikasi semakin berat/tidak mudah di scale karena harus manual menambahkan komunikasi.
</p>
  
* Keuntungan **Publish-Subscribe**:

<p style="text-align: justify;">  
  RPC adalah mekanisme komunikasi dimana pengirim menentukan siapa yang menerima data. Nah, ada alternatif untuk decoupling proses ini yaitu Publish-Subscribe. Alih-alih mengirimkan data ke consumer. Publish-Subscribe memiliki mekanisme messaging dimana pengirim tidak harus menentukan siapa penerimanya, tetapi dia hanya perlu mengirim ke yang namanya **message broker**.
  
  **Message Broker** sendiri akan berperan sebagai perantara yang dimana membuat tugas pengirim data tidak harus tau kepada siapa pengirimnya tetapi langsung ke parantara aja (biar dia yang handle). Sehingga dengan demikian Diagram Messagingnya akan menjadi seperti dibawah ini:
  
  <img width="1016" height="517" alt="image" src="https://github.com/user-attachments/assets/24087d0c-a115-4eac-8c3b-03675c471e38" />
  
  Dengan begitu ini akan jadi lebih optimal dalam hal **scalability** maupun **flexibility**.
</p>
  
* Advantages:
  1. **Code Decoupling**: Pengirim (`Producer`) tidak perlu tau kompleksitas yang akan dilakukan penerima data (`consumer`).
  2. Setiap terjadi perubahan jumlah penerima (`consumer`) data, pengirim tidak perlu tau.
       
     ex. ) kasus penambahan penerima data seperti Fraud Detection yang mengharuskan kode dalam Order Services dan Fraud Detection diubah. Dengan publish-subscribe, kode yang harus diubah hanya Fraud Detection untuk menerima data dari message broker. Pengirim juga cuman perlu mengirim data sebanyak 1x.

* Disadvantages:
  1. Tidak realtime seperti RPC, ada jeda/delay yang membuat data tidak konsisten secara langsung (eventually consistent). Contohnya, bisa aja producer mengirim datanya sudah selesai tapi consumer belum selesai mengonsumsi + memproses data tersebut.
  2. Ketika terjadi kegagalan pengiriman data, pengirim tidak akan tau. Oleh karena itu, harus ada `retry method`.
  
* Publish-Subscribe Model

<img width="1218" height="576" alt="image" src="https://github.com/user-attachments/assets/f3b2ef27-96ec-453a-a2fc-5c29346a6884" />

* **Publisher**: ga peduli cara kerja subscriber-nya, yang dia peduliin adalah publish message (data) aja. (Bahkan kalau message tersebut tidak di proses sama sekali, dia tidak peduli).
* **Subscriber**: ga perlu tau cara kerja services dari publisher, data-nya darimana. Literally, ada topic dimana topic itu representasi dari messages-messages dengan kategori tertentu, maka dia akan proses hal tersebut.

## Kafka

Kafka adalah contoh dari **Message Broker** untuk mekanisme messaging. Kafka juga sering digunakan di distributed commit log & streaming services.  

* Kenapa Kafka?
  1. **Scalability**: Kafka mampu menerima **`overload traffic`** dengan baik.
  2. **High Performance**: Banyak data dalam message broker itu tidak ngaruh terhadap performance pada Message Brokernya.
  3. **Persistance**: Data dalam message broker itu pasti ada meskipun penerima data gagal memproses.

### Topics

Saat kita mengirim data, nah itu kita kirim ke yang namanya **topic**. Topic itu mirip table dalam database, yang digunakan untuk **menyimpan data** yang dikirim oleh pengirim data. (**`Producer -> Topic -> Consumer`**)
  
* Diagram Topic

<img width="1020" height="720" alt="image" src="https://github.com/user-attachments/assets/f9348421-be0f-445e-88cb-6ce0d13dc8a5" />
  
Nah, misal order ngirim data di topic order. Nah, siapapun yang membutuhkan data order akan diambil dari topic order. Begitu pun seperti Wallet Services & Notification Services yang mau mengambil data dari member.
  
Idealnya 1 topic = 1 jenis data.

* Log
  
**Data di topic disimpan dalam format log**. Apa itu log? log adalah cara menyimpan data yang paling sederhana (macam Queue), yaitu **`append-only`** sesuai urutan masuk.
  
<img width="912" height="1340" alt="image" src="https://github.com/user-attachments/assets/ecb86018-ee6c-4c38-8904-72c4cc555dbf" />

* Message

Data yang dikirim di Kafka itu kita sebut sebagai **"Message"**. Struktur data dari Messages udah di set oleh Kafka, yaitu kek dibawah ini:  

<img width="1244" height="352" alt="image" src="https://github.com/user-attachments/assets/f7f14450-92dd-4a63-9566-c34400b0c153" />
  
* Producer  

Pengirim data = Producer. (Data yang dikirim pasti ada di belakang (`push()`).  

* Consumer

Penerima data = Consumer. (Consumer akan membaca data secara berurutan, dari No. Message paling awal ke akhir dalam sebuah partition pada sebuah topic (partition ada di dalam topic)). Tapi bisa juga langsung baca data yang paling akhir. Dan proses pembacaan data ini bersifat realtime, jadi ketika ada data baru dalam queue maka data tersebut akan langsung dibaca juga.  

* Consumer Group

Saat consumer membaca data dari topic, maka consumer perlu menentukan consumer group mana yang mau digunakan. Kalau ga ditentukan, Kafka secara otomatis membuat consumer group baru. Tapi, untuk penerapan yang lebih baik, consumer group harus selalu disebutkan. Jika tidak menyebutkan consumer group, consumer group baru akan selalu terbentuk, namun Kafka memegang prinsip **'At-least Once'** yang artinya, data dalam message broker hanya akan dikirim sebanyak 1 kali. **NAH, KALAU CONSUMER GROUP-NYA ADA BANYAK MAKA DATA YANG SUDAH DIKIRIM AKAN DIKIRIMKAN BERKALI-KALI**.  

<img width="1119" height="298" alt="image" src="https://github.com/user-attachments/assets/88f70ab9-7749-4be4-8614-406f17f724e3" />

Anggep Payment Services itu menerima data, dimana data yang diterima akan dikirimkan ke database. Nah, hal ini bisa menyebabkan duplikat primary key, jadi better messagesnya kita kirim 1 kali aja antara ke server 1 atau server 2. Hal ini bisa di achieve dengan membuat sebuah Consumer Group antara kedua Payment Services (Mereka berdua adalah Consumer) tersebut, sehingga data hanya akan diterima oleh 1 Consumer.  

* Offset

Kita tau data di Topic itu disimpan secara berurutan. Nah, kalau semua Consumer kita matiin tapi Producer masih terus mengirim data, ketika Consumer berjalan lagi, darimana consumer akan membaca data? Dari awal? Dari yang terbaru? **atau dari data terakhir semenjak consumer dimatikan?**. Nah, Kafka punya fitur yang namanya **Offset**.  
  
  i. Default:  
    Secara default, consumer akan membaca data terbaru aja. Misal, ada 10 data yang sudah dikirim. Lalu Consumernya dimatikan semua. Ketika dimatikan, ada 2 data yang dikirim yaitu data ke-11 dan data ke-12. Ketika consumer dinyalakan kembali, maka dia akan membaca data ke-13 (kalau ada), dan data ke-11 dan data ke-12 tidak akan dibaca.
    
  ii. Offset:  
    Kafka menyimpan data yang terakhir dibaca yang disebut sebagai **Offset**. Ketika consumer pertama kali dijalankan, data offset itu tidak ada, jadi dia cuman bisa dijalankan dengan 2 cara:

    - `--from-beginning`
    - `default` (dikosongin)  
    
  Offset disimpan sesuai consumer group, jadi kalau consumer group dan consumer yang sama, maka offset-nya juga akan sama. Cara pakai offset, cukup jalankan consumer yang sudah dimatikan dengan `--from-beginning`.
  
* Partition

Saat topic kita buat, topic tersebut akan disimpan di 1 partition. Jumlah partition di Kafka secara default itu 1. Ini yang sebenarnya jadi alasan ketika ada 2 consumer dalam 1 consumer group yang sama, dia dikirim ke 1 consumer aja. Tapi 1 consumer group bisa baca lebih dari 1 partition, dimana 1 partition akan menerima messages. Ketika ada 1 partition dan 2 consumer, maka data hanya akan dibaca oleh 1 consumer, ini menjelaskan kenapa consumer group memiliki fitur tersebut.  

<img width="1211" height="741" alt="image" src="https://github.com/user-attachments/assets/ca009c7e-b870-4c7d-b0db-437b5992d6f8" />

* Routing

Gimana cara menentukan ke partition mana sebuah data akan dikirimkan? Nah, penentuan partition yang dipilih, ditentukan dari **key** yang terdapat pada messages yang kita pilih.  

* Key Routing

Key pada messages digunakan untuk memilih partition mana yang akan dikirim data tersebut dimana key tersebut akan melewati proses hashing. (Key yang sama pasti akan masuk ke partition yang sama).
