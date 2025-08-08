package messaging

import (
	"context"
	"fmt"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

// step 1: define a handler function type for consuming messages (disini string topic akan di define dari objek Record yang bisa diakses pake record.Topic)
type ConsumerHandler func(record *kgo.Record) error

func ConsumeTopic(ctx context.Context, consumer *kgo.Client, handler ConsumerHandler) {
	// step 1: mulai for loop untuk consume message dari topic yang sudah ditentukan

	/*
		kenapa harus pake context? karena kita butuh untuk cancel atau stop consume message ketika aplikasi sudah tidak berjalan lagi
		jadi kita bisa cancel context ini ketika aplikasi sudah tidak berjalan lagi

		kenapa harus pake select? karena kita butuh untuk menunggu message yang masuk dari topic yang sudah ditentukan
		jadi kita bisa pake select untuk menunggu message yang masuk dari topic yang sudah ditentukan

		kenapa harus pake default? karena kita butuh untuk terus menunggu message yang masuk dari topic yang sudah ditentukan
		jadi kita bisa pake default untuk terus menunggu message yang masuk dari topic yang sudah ditentukan

		kenapa harus pake for loop? karena bentuk consume message dari topic yang sudah ditentukan adalah terus menerus
		jadi kita bisa pake for loop untuk terus menerus consume message dari topic yang sudah ditentukan
	*/
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

// Notes: kalau bingung/lupa, lihat ke cmd/worker/main.go, disitu ada contoh implementasi untuk consume message dari topic yang sudah ditentukan
// Topic ditentukan pada saat membuat consumer sehingga tidak perlu lagi menentukan topic disini
