package config

import (
	"time"

	"github.com/segmentio/kafka-go"
)

func NewKafkaConsumer(topic string, groupID string) *kafka.Reader {
    // step 1: create configuration-nya
	config := kafka.ReaderConfig{
        Brokers:        []string{"localhost:9092"},
        Topic:          topic,              // topic yang mau di subscribe
        GroupID:        groupID,            // consumer-group ID
        StartOffset:    kafka.FirstOffset,  // mulai dari offset pertama --from-beginning
        MinBytes:       10e3,               // 10KB
        MaxBytes:       10e6,               // 10MB
        MaxWait:        time.Second,        // latency
    }

    // step 2: buat reader-nya / consumer-nya
    consumer := kafka.NewReader(config)

    // step 3: return reader-nya
    return consumer
}

func NewKafkaProducer(topic string) *kafka.Writer {
    // step 1: buat writer-nya (config di satukan dalam inisialisasi writer)
    writer := &kafka.Writer{
        Addr:           kafka.TCP("localhost:9092"),
        Async:          true,
        Topic:          topic,                              // topic yang mau di publish
        RequiredAcks:   kafka.RequireOne,                   // mengharuskan setidaknya 1 broker menerima pesan (At-Least-Once Delivery)
        Balancer:       &kafka.LeastBytes{},                // menggunakan LeastBytes balancer untuk distribusi pesan (Load Balancing)
        BatchSize:      100,                                // ukuran batch untuk pengiriman pesan
        BatchTimeout:   10 * time.Millisecond,              // timeout untuk batch
        Compression:    kafka.Lz4,                          // menggunakan kompresi LZ4 untuk mengurangi ukuran pesan (Compression Algorithm)
    }

    // step 2: return writer-nya
    return writer
}