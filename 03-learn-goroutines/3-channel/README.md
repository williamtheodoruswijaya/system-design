## Apa itu Channel?

- Tempat komunikasi secara synchronous yang bisa dilakukan oleh goroutine.
- Ini yang menjadi solusi ketika kita membutuhkan return value dari sebuah function yg dijalankan oleh goroutine.
- Kita tau function yang dipanggil dengan keyword `go`, value yang dikembalikan tidak bisa direturn.
- Di channel terdapat pengirim dan penerima, biasanya pengirim dan penerima adalah goroutine berbeda. Jadi ada goroutine pengirim value (function yg return valuenya) dan goroutine yang penerima value (goroutine yang menerima value yang direturn via channel).
- Saat melakukan pengiriman data ke Channel, goroutine akan ter-block (goroutine tidak akan jalan sampai data di dalam channel ada yang mengambilnya).
- Maka dari itu, Channel disebut sebagai alat komunikasi  yang syncrhonous.
- Channel ini cocok sebagai alternatif seperti async await. Dimana kali ini **async** dalam Golang itu adalah Goroutine dan **await** dalam Golang itu adalah Channel.

![channel-diagram](channel-diagram.png)


## Karakteristik Channel

- Secara default, channel hanya bisa menampung 1 data, tunggu sampai itu diambil, baru bisa menambah data lagi.
- Channel hanya bisa menerima satu jenis data (tentukan tipe data-nya).
- Channel bisa diambil dari lebih dari satu goroutine. (Pengirim bisa banyak tapi hanya 1 value yang bisa diterima pada saat yang bersamaan).
- Channel harus di close (best practices) jika tidak digunakan untuk menghindari memory leak.


## Cara membuat channel

- Channel di Golang direpresentasikan dengan tipe data Chan.
- Untuk membuat Channel terbilang sangat mudah, kita bisa tinggal menggunakan **`make()`**, mirip ketika membuat map.
- Namun saat pembuatan channel, kita harus tentukan tipe data apa yang bisa dimasukkan ke dalam channel tersebut.

```go
// contoh channel dengan tipe data string

channel := make(chan string)

// ...

close(channel)
```

## Mengirim dan Menerima Data dari Channel

- Channel bisa digunakan untuk mengirim dan menerima data.
- Untuk mengirim data, kita bisa menggunakan kode: **`channel <- data`**.
- Sedangkan untuk menerima data, kita bisa menggunakan kode: **`data <- channel`**.
- Setelah selesai, selalu close channel.


## Channel sebagai parameter

- Dalam kenyataan pembuatan aplikasi, seringnya kita akan mengirim channel ke function lain via parameter.
- Sebelumnya kita tau bahkan di Golang by default, parameter adalah pass by value, artinya value akan diduplikasi lalu dikirim ke function parameter, sehingga jika kita ingin mengirim data asli, kita biasa menggunakan pointer (pass by reference).
- Berbeda dengan channel, kita tidak perlu melakukan hal tersebut.

Contoh kode:
```go
func sendData(channel chan string, data string) {
	time.Sleep(1 * time.Second)
	channel <- data
}

func main() {
	// 1. create the channel (take this as a variable that will change it's value)
	channel := make(chan string)
	defer close(channel)

	// 2. send the data via goroutines
	go sendData(channel, "This's the data")

	// testing purposes
	fmt.Println("I'm runned before sendData()")

	// 4. receive the data
	random_variable := <-channel

	// testing purposes
	fmt.Println("I'm waiting the data received...")

	// 5. line below 25 will be runned after the data are received...
	fmt.Println(random_variable)
}
```

## Channel In and Out

- Saat kita mengirim channel sebagai parameter, isi function tersebut bisa mengirim dan menerima data dari channel tersebut.
- Lalu bagaimana cara kita memberi tahu terhadap function, bahwa channel tersebut hanya digunakan untuk mengirim data, atau hanya dapat digunakan untuk menerima data?
- Hal ini bisa kita lakukan di parameter dengan cara menandai apakah channel ini digunakan untuk in (mengirim data) atau out (menerima data).

Contoh code:

- Channel In (`channel chan<- string`)
```go
func SendData(channel chan<- string, data string) {
	channel <- data
}
```

- Channel Out (`channel <-chan string`)
```go
func ReceiveData(channel <-chan string, *data string) {
	data := channel
}
```

-  Usage Example:
```go
func SendData(channel chan<- string, data string) {
	time.Sleep(1 * time.Second)
	channel <- data
}

func ReceiveData(channel <-chan string, data string) {
	data = <-channel
	fmt.Println(data)
}

func main() {
	channel := make(chan string)
	defer close(channel)

	go SendData(channel, "Hello World")
	go ReceiveData(channel, "Hello World")

	time.Sleep(3 * time.Second)
}
```

## Buffered Channel

- Secara default channel hanya bisa menerima 1 data. Artinya, kita sebenarnya bisa mengubah default pengaturan ini.
- Jika kita menambah data ke-2, maka kita akan diminta menunggu sampai data ke-1 ada yang mengamil.
- Kadang-kadang ada kasus dimana pengirim lebih cepat dibanding penerima, dalam hal ini jika kita menggunakan channel, maka otomatis pengirim akan ikut lambat juga. (Penerima lambat, Pengirim pun akan ikut lambat).
- Untuknya ada buffered channel, yaitu buffer yang bisa digunakan untuk menampung data antrian di Channel.


## Buffer capacity

- Buffer itu anggepannya sebagai penyimpanan di dalam channel.
- Kita bebas memasukkan berapa jumlah kapsitas antrian dalam buffer.
- Jika kita set misal 5, artinya channel bisa menerima 5 data.
- Jika kita mengirim data ke-6, maka kita diminta untuk menunggu sampai buffer ada yang kosong.
- Ini cocok sekali ketika memang goroutine yang menerima data lebih lambat dari yang mengirim data.

Dari sini kita bisa mengubah Channel menjadi sebuah queue dan menghindari blockingan await yang terlalu lama ketika receiver lamban dalam memproses.

Buffer membuat sebuah channel bisa mengirimkan value tanpa menunggu value tersebut di receive (istilahnya ma gitu).

![channel-buffer](channel-buffer.png)

Cara buatnya:
```go
channel := make(chan string, 3) // by default 1, but we add 3 as a buffer capacity.

fmt.Println(cap(channel)) // melihat panjang buffer (buffer capacity)
fmt.Println(len(channel)) // melihat jumlah data di dalam buffer
```


## Select Channel

- Kadang ada kasus dimana kita membuat beberapa channel dan menjalankan beberapa goroutine untuk tiap channelnya.
- Kalau kita ingin mendapatkan data dari semua channel tersebut, kita bisa menggunakan select channel
- Dengan select channel, kita bisa memilih data tercepat dari beberapa channel, jika data datang secara bersamaan di beberapa channel, maka akan dipilih secara random.

Kalau masih ga kebayang coba baca code ini aja:
```go
counter := 0
for {
	select {
	// intinya ketika ada 
	case data := <-channel1:
		fmt.Println("Data dari Channel 1", data)
		counter++
	case data := <-channel2:
		fmt.Println("Data dari Channel 2", data)
		counter++
    }
	if counter == 2 { // ini buat jaga-jaga aja biar for loopnya ga infinite
		break
    }
}
```
Intinya, kalau ada banyak data yang dikirim ke sebuah channel secara bersamaan, dengan menggunakan select kita bisa membuat Go secara otomatis mengambil data dalam channel tersebut berdasarkan yang tercepat baru setelah itu di perulangan selanjutnya ngambil data di channel lain. (instead of pake multiple for loop buat setiap channel).


## Default Select

- Apa yang terjadi jika kita melakukan select terhadap channel yang ternyata tidak ada datanya? Deadlock bro.
- Atau kita akan menunggu sampai data ada
- Kadang mungkin kita ingin melakukan sesuatu jika misal semua channel tidak ada datanya ketika kita melakukan select channel
- Dalam select, kita bisa menambahkan default, dimana ini akan dieksekusi jika memang di semua channel yang kita select tidak ada datanya.
- **Default akan dijalankan ketika sebuah channel tidak ada datanya.**

Contoh kasus dalam implementasi `consumer.go` pada franz-go library (kafka library in go):
```go
func ConsumeTopic(ctx context.Context, consumer *kgo.Client, handler ConsumerHandler) {
	// step 1: mulai for loop untuk consume message dari topic yang sudah ditentukan
	for {
		select {
		// step 2: kalau aplikasi sudah tidak berjalan lagi, kita bisa cancel context ini, sehingga consumer akan berhenti consume message dari topic yang sudah ditentukan
		case <-ctx.Done():
			fmt.Println("Context cancelled, stopping consumer")
			return

		// step 3: consume message dari topic yang sudah ditentukan
		default:

			// step 4: poll fetches untuk mendapatkan partition2 dari topic yang sudah ditentukan, ini akan mengembalikan partition-partition yang berisi message-message
			fetches := consumer.PollFetches(ctx)

			// Kalau gagal, coba lagi
			err := fetches.Err()
			if err != nil {
				// Appendix: Kasih delay sebelum mencoba lagi sekitar 1 detik untuk menghindari busy loop
				time.Sleep(1 * time.Second)
				fmt.Printf("Error polling fetches: %v\n", err)
				continue
			}

			// step 5: iterasi setiap partition dari topic X dan baca setiap message yang ada di dalamnya, parameter-nya adalah function yang berfungsi untuk membaca setiap message yang ada dalam partition
			fetches.EachPartition(func(partition kgo.FetchTopicPartition) {
				for _, record := range partition.Records {
					fmt.Printf("Received message from topic %s: %s\n", partition.Topic, string(record.Value)) // Print the topic and message for tutorial purpose only

					// step 6: panggil handler function untuk ubah record menjadi model json yang sesuai sekaligus proses tambahan untuk menyimpan ke database atau melakukan proses lainnya
					err := handler(record)
					if err != nil {
						fmt.Printf("Error handling record from topic %s: %v\n", partition.Topic, err)
					} else {
						// step 7: commit records setelah sukses
						err = consumer.CommitRecords(ctx, record)
						if err != nil {
							fmt.Printf("Error committing record to topic %s: %v\n", partition.Topic, err)
						}
					}
				}
			})
		}
	}
}
```
Kode diatas terdapat select dimana dia akan dijalankan ketika `ctx.Done()` return sebuah value yang akan dijalankan ketika contextnya di cancelled (program selesai). Nah, ini membuat for loop tersebut akan dijalankan pada bagian default dimana di bagian default terdapat logic untuk melakukan consume terhadap suatu topic yang di sent oleh publisher.

Kalau belum ngerti kafka, bisa lihat [codingan-nya aja yang lebih simple](7-default-select.go).