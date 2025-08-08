
## Parallel Programming

- Kita hidup di zaman dimana CPU dan GPU semua sudah multicore.
- Perkembangan hardware ini biasa akan diikuti oleh perkembangan software yang juga berkembang menyesuaikan dengan hardware.
- Parallel programming memungkinkan sebuah software untuk memanfaatkan lebih dari 1 core CPU untuk menjalankan sebuah tugas.
- Analoginya mirip seperti dapur dimana ada lebih dari 1 koki yang mengerjakan tugasnya masing-masing.

Dalam parallel programming terdapat istilah yaitu Process vs Thread.


## Process vs Thread

Process:
1. Sebuah eksekusi program. Contoh ketika kita menjalankan aplikasi Word, nah itu disebut sebagai Process.
2. Process mengonsumsi memory yang sangat besar.
3. Process itu saling terisolasi dengan process lain.
4. Process lama untuk dijalankan dan dihentikan (contoh menjalankan aplikasi Adobe Photoshop)

Thread:
1. Segmen/Bagian kecil dari Process. Contoh, kita membuka Google Chrome. Nah tab dalam Google Chrome disebut sebagai thread sementara Google Chromenya sendiri adalah process.
2. Thread mengonsumsi memory yang lebih kecil dikarenakan thread sendiri adalah sebagian kecil dari Process dan Process adalah gabungan berbagai threads.
3. Thread bisa saling berhubungan jika dalam process yang sama. Contoh, kita bisa mengirimkan data dari satu thread ke thread yang lain.
4. Thread mudah untuk dijalankan dan dihentikan.


## Parallel vs Concurrency

1. Berbeda dengan parallel (menjalankan berbagai pekerjaan secara bersamaan), concurrency adalah menjalankan beberapa pekerjaan secara **bergantian**.
2. Dalam parallel biasanya kita membutuhkan banyak Thread, sedangkan dalam concurrency, kita hanya membutuhkan sedikit Thread. (Karena dalam concurrency itu kita fokusnya ke bergantian mengerjakan sesuatu).

Contoh concurrency:

Saat kita makan di cafe, kita bisa makan, lalu ngobrol, lalu minum, makan lagi, ngobrol lagi, minum lagi, dan seterusnya. Tetapi kita tidak bisa pada saat yang bersamaan minum, makan, dan ngobrol, hanya bisa melakukan satu hal pada satu waktu, namun bisa berganti kapanpun kita mau.

## CPU-bound

- Banyak algoritma dibuat yang hanya membutuhkan CPU untuk menjalankannya. Algoritma jenis ini biasanya sangat tergantung dengan kecepatan CPU.
- Contoh yang paling populer adalah Machine Learning, oleh karena itu sekarang banyak sekali teknologi Machine Learning yang banyak menggunakan GPU karena core yang lebih banyak dibanding CPU biasanya.
- Jenis algoritma seperti ini tidak ada benefitnya menggunakan Concurrency Programming, namun bisa dibantu dengan implementasi Parallel programming.


## I/O-bound

- I/O-bound adalah kebalikan dari sebelumnya, dimana biasanya aplikasi sangat tergantung dengan kecepatan input output devices yang digunakan.
- Contoh aplikasinya seperti membaca data dari file, database, dan sebagainya.
- Kebanyakan software engineer membuat aplikasi dengan jenis seperti ini.
- Aplikasi jenis I/O bound, walaupun tetap bisa terbantu dengan Parallel Programming, tapi benefitnya akan lebih baik jika menggunakan Concurrency Programming.
- Bayangkan kita membaca data dari database, dan Thread harus menunggu 1 detik untuk mendapat balasan dari database, padahal waktu 1 detik itu jika menggunakan Concurrency Programming, bisa digunakan untuk melakukan hal lain lagi.
