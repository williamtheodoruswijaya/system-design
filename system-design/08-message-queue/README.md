# Message Queue (RabbitMQ)

## Apa itu Message Queue?

Ini adalah sejenis service-to-service communication yang terjadi secara asynchronous.<br/>
Alurnya sesimple:

<img width="677" height="340" alt="image" src="https://github.com/user-attachments/assets/1e27db73-a40a-445d-af54-fab7c401f4fa" />

### Q: Kenapa kita perlu Message Queue?

Kan bisa aja service A dan service B berkomunikasi via metode RPC seperti RESTful API, kenapa harus pakai Message Queue? Karena tidak semua hal itu bisa diselesaikan dengan cepat. Ibaratnya kalau dengan metode RPC (RESTful API), misal Service A dan Service B itu saling berhubungan dan ketika service A punya banyak request, otomatis prosesnya akan lambat (more response time), otomatis service B yang misal aja kosong, akan menunggu service A untuk selesai, menyebabkan kedua service ikut menjadi lama, yang padahal seharusnya service A saja yang response timenya lambat. Dengan message queue, kita bisa membuat service B bekerja tanpa menunggu service A selesai memproses semua data-nya.<br/>

### 📸 Contoh kasus

Misalnya aplikasi edit foto:

* User upload foto → Service A (API).
* Setelah edit, foto perlu di-render (proses berat).

Kalau pakai REST API biasa:

* Service A harus nunggu Service B selesai render.
* Kalau banyak user upload sekaligus → Service A lambat, user harus nunggu lama.

Kalau pakai Message Queue:

* Service A langsung taruh “tugas render” di queue, balas ke user: “Foto sedang diproses”.
* Service B ambil tugas dari queue, render foto di background.
* User bisa lanjut pakai aplikasi tanpa delay panjang.

Notes: dengan Message Queue, data yang ada di queue akan terus ada sampai data tersebut di-consume. Dan setiap message akan diconsume oleh consumer at-least once.

### Advantages:

Apa sih keuntungan menyimpan HTTP Request dalam sebuah antrian?<br/>
1. **Performance**: Karena message queues membuat request dihandle secara **asynchronous**, maka **producer can add request to queue without waiting for the requests to be processed.**

ex.) Sebuah service yang menerima request dari service lain butuh 10 detik untuk memproses sebuah request. Sehingga, kalau misal ada 10 request, maka user ke-10 akan menunggu 100 detik untuk menunggu service mengirim request tersebut diproses oleh service lain. Nah, tapi dengan Message Queues, setiap request bisa di proses secara asynchronously (artinya, service dari client-side bisa mengirim request ke service B, meskipun service B sedang sibuk memproses sebuah request).

2. **Reliability**: Data yang disimpan di message queues, dalam konteks ini yaitu request, bersifat persistent (artinya, pasti akan ada terus selama belum di consume).

3. **Decoupling**: Producer (Service A) hanya bertugas untuk mengirim message (dalam konteks ini berarti request) tanpa peduli berapa lama atau bagaimana consumer mengonsume datanya. Menghilangkan dependencies antar components.
ex.) Kalau gaada Message Queues, kita harus expect gimana request-nya untuk dihandle oleh Consumer (kalau error maka otomatis service yang mengirim request tersebut akan error juga, menyebabkan seluruh proses gagal). Dengan Message Queues, perubahan di producer tidak akan mempengaruhi Consumer dan Vice versa. (Sekalipun ada error di Consumer, Producer tetap akan berjalan dan mengirim data ke Queue, cuman paling data di queue-nya aja stop di consume).

4. **Scalability**: Dengan adanya Message Queue, data yang banyak banget di producer, kita bisa tetap kirim semua itu ke dalam Message Queue, dan dengan Consumer yang pasti akan process Message Queues secara satu per satu, menjamin message queues, producer, dan consumer tidak akan overload. **When workloads peak, consumer/producer not overload.**

### Kenapa Message Queue bisa Grow Significantly Without Having a Slow Performance?

Ada teknik yang namanya **Backpressure Technique**. Kita tau kalau Message Queue punya banyak data, bisa aja performancenya ngaruh, tapi pada kenyataannya tidak ngaruh sama sekali, Kenapa? Karena Backpressure melimit queue size dari Message Queue.<br/>
Workflow-nya: se-simple kasih limit ke Message Queues, berapa Message yang bisa dia tampung, misal 10, setelah 10, clients akan mendapat error server busy 503 yang basically nyuruh client buat nyoba di lain waktu (ofc mengurangi reliablity sebuah app). Dan Message Queue punya Consumer yang hanya bisa mengonsumsi 1 request (1 consumer cuman bisa memproses 1 request).

### How Message Queues handle Queue?

1. **First-In First-Out (FIFO)**

2. **At-least-once + FIFO (redundancy)**<br/>
ex.) Message Queues membuat beberapa message yang duplicates. Kegunaannya, kalau 1 message error ketika diconsume, dia ada backup. Analogi sederhananya, kalau ada 1000 message dalam message queues, maka kita akan kalikan 2 untuk setiap message, dan each message akan di consume "At-least once".

3. **Exactly once (filter duplicates)**<br/>
Secara otomatis menghapus duplicate request ketika sebuah request berhasil di consume.
