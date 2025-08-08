## Pengenalan Goroutines

- Goroutine adalah sebuah thread ringan (running di dalam thread) yang dikelola oleh Go Runtime.
- Ukuran Goroutine itu kecil banget (2Kb) jauh lebih kecil daripada Thread yaitu sekitar 1000Kb.
- Namun tidak seperti thread yang berjalan parallel, goroutine berjalan secara concurrent.
- Saat goroutine dijalankan, dia akan dijalankan di dalam sebuah thread, kemudian dia akan dijalankan secara ganti-gantian (kadang-kadang Goroutine A, belum beres, pindah Goroutine B, balik A, dst...)


## Cara Kerja Goroutine

- Goroutine dijalankan oleh Go-Scheduler dalam sebuah thread. Dimana di dalam Go-Scheduler ada banyak thread dan jumlah thread-nya diatur berdasarkan GOMAXPROCS (biasanya sejumlah core CPU) <- nanti dibahas.
- Goroutine tidak bisa dibilang sebagai pengganti thread karena Goroutine sendiri berjalan diatas thread.
- Namun yang mempermudah kita adalah, kita tidak perlu melakukan manajemen Thread secara manual, semua sudah diatur oleh Go Scheduler.


Dalam Go-Scheduler, ada beberapa terminologi:
- G: Goroutine
- M: Thread (Machine)
- P: Processor

Setiap kita membuat sebuah Goroutine, maka akan masuk ke queue. Ada 2 jenis queue yaitu:

1. Global Queue
2. Local Queue

Thread akan mengambil goroutine dalam Local Queue dan Global Queue. Ketika Goroutinesnya habis, dia akan mencuri dari Queue lain. Concurrencynya ada dimana? Nah, ketika Goroutines diambil oleh thread dan eksekusinya lama, maka Goroutines tersebut akan di pause dan dikembalikan ke queue lalu mengambil Goroutines yang lain. Jadi, ketika ada Goroutines yang lama progress-nya, dia tidak akan menjadi blockingan untuk komputernya.