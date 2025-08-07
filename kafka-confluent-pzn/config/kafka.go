package config

import (
	"os"

	"github.com/twmb/franz-go/pkg/kgo"
)

func NewConsumer() *kgo.Client {
	// step 1: inisialisasi kafka client configuration
	brokers := []string{os.Getenv("KAFKA_CLIENT")}		// address message broker (kalau local ya localhost:9092)
	groupId := "test-group-0"							// group id untuk consumer, bisa diisi sesuai keinginan (consumer group id)
	topic := "test-topic-belajar-kafka"					// nama topic yang akan di-consume

	// step 2: buat kafka client/consumer (disini khusus consumer maupun producer, sebutannya client kalau pake franz-go)
	// perbedaan ada di konfigurasi, kalau consumer ada group id dan topic yang akan di-consume
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.ConsumerGroup(groupId),
		kgo.ConsumeTopics(topic),
	)
	if err != nil {
		panic(err)
	}

	// step 3: return client
	/* 

		- client ini akan digunakan untuk melakukan consume message dari topic yang sudah ditentukan
		- client ini juga akan digunakan untuk melakukan commit offset setelah message berhasil diproses
		- commit offset ini penting agar consumer tidak mengulang membaca message yang sudah diproses sebelumnya
		- jika tidak melakukan commit offset, consumer akan terus membaca message yang sama setiap kali dijalankan
		- sehingga menyebabkan duplikasi data atau proses yang tidak diinginkan
		
		- client ini juga akan digunakan untuk melakukan close connection setelah selesai digunakan
		- close connection ini penting agar tidak terjadi memory leak atau resource yang tidak terpakai
		- jika tidak melakukan close connection, maka connection akan tetap terbuka dan menghabiskan resource
		- sehingga menyebabkan aplikasi menjadi lambat atau tidak responsif
		- jadi pastikan untuk selalu melakukan close connection setelah selesai menggunakan client
	
	*/ 
	return client
}

func NewProducer() *kgo.Client {
	// step 1: inisialisasi kafka client configuration
	brokers := []string{os.Getenv("KAFKA_CLIENT")} // address message broker (kalau local ya localhost:9092)

	// step 2: buat kafka client/producer
	// producer tidak membutuhkan group id dan topic, karena akan mengirim message ke topic yang ditentukan saat publish, jadi cukup dengan seed brokers saja
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
	)
	if err != nil {
		panic(err)
	}

	// step 3: return client/producer
	/*

		- client ini akan digunakan untuk melakukan publish message ke topic yang sudah ditentukan
	
	*/
	return client
}
