# How to setup Kafka in Golang DDD Architecture:

1. **Buat contract untuk semua Event**. Dilakukan di (`internal/model/`).
   - Ini adalah bagian paling fundamental, yaitu [`**Event Interface**`](internal/model/event.go). Ini adalah kontrak/aturan dasar untuk semua event.
   - Kenapa ini harus ada? Tujuannya adalah memaksa setiap model event agar memiliki function getter yaitu `GetId()`. Fungsi ini krusial karena nilai-nya akan digunakan sebagai `key`.
   - Manfaat: penggunaan key memastikan semua event terkait dengan **entity yang sama masuk ke partition yang sama**.

2. **Buat model event spesifik berdasarkan tabelnya**. Dilakukan di (`internal/model/event/`).
   - Berdasarkan kasus ini, kita akan membuat **[`user_event.go`](internal/model/event/user_event.go)**
   - Setelah menyiapkan aturan/function yang tiap model harus ikuti. `user_event.go` adalah implementasi dari kontrak tersebut sekaligus mendefinisikan data apa saja yang dikirim saat ada event terkait user.
   - Jadi setiap struct `UserEvent` wajib memiliki fungsi `GetId()` untuk memenuhi kontrak event interface.

3. **Buat consumer dan handler untuk message yang diconsume**. Dilakukan di (`internal/delivery/messaging/`).
   - Function `ConsumeTopic()` adalah sebuah loop yang akan secara terus menerus mengambil pesan dari partitions-partitions2 yang ada.
   - [`consumer.go`](internal/delivery/messaging/consumer.go) bertugas untuk:
     - connection ke kafka
     - fetching messages dari setiap partitions dalam 1 topik.
     - menangani sinyal berhenti dari context.
     - mengambil message & meneruskannya ke [`handler`](internal/delivery/messaging/user_consumer.go) atau `user_consumer.go`.
     - Tugas dari `user_consumer.go` biasanya:
       - menerima pesan dari topic
       - melakukan unmarshal dari json ke `UserEvent` struct
       - melakukan tindakan lain (email notification, store database, etc.)

4. **Buat Producer & user_producer**. Dilakukan di (`internal/gateway/messaging`).
   - Sebuah aplikasi bisa aja memiliki banyak jenis event: (UserEvent, OrderEvent, dsbnya). Bagaimana kita bisa memastikan developer tidak salah mengirim ke topic yang salah. (UserEvent (message tentang user) dikirim ke topic Orders).
   - Simple-nya kita harus membuat Producer menjadi sebuah Generic Class agar instance producer yang dibuat terikat pada Event tersebut.
   - Contoh: `producer[UserEvent]` yang artinya producer ini hanya bisa mengirim message `UserEvent` saja ke topic "users".
   - Beda dengan consumer. Consumer pada saat pembuatannya sudah di tentukan ingin mendengarkan topik apa sementara producer engga.