# Kafka Golang DDD Architecture Implementation

## Description

This is golang clean architecture made only for Kafka implementation in DDD architecture.

![Clean Architecture](architecture.png)

1. External system perform request (HTTP, gRPC, Messaging, etc)
2. The Delivery creates various Model from request data
3. The Delivery calls Use Case, and execute it using Model data
4. The Use Case create Entity data for the business logic
5. The Use Case calls Repository, and execute it using Entity data
6. The Repository use Entity data to perform database operation
7. The Repository perform database operation to the database
8. The Use Case create various Model for Gateway or from Entity data
9. The Use Case calls Gateway, and execute it using Model data
10. The Gateway using Model data to construct request to external system
11. The Gateway perform request to external system (HTTP, gRPC, Messaging, etc)

## Tech Stack

- Golang : https://github.com/golang/go
- Apache Kafka : https://github.com/apache/kafka

## Framework & Library

- Franz Kafka Golang : https://github.com/twmb/franz-go


### Run web server

```bash
go run cmd/web/main.go
```

### Run worker

```bash
go run cmd/worker/main.go
```

## Documentation

Oke jadi pertama kalau mau integrasi Kafka ke DDD architecture Golang. Pertama kita perlu ngerti dulu sedikit dari tutorial mulai dari cara buat:
- consumer
- producer
- send message
- consume message

1. Cara buat consumer (ada di `config/kafka.go`):
```go
func NewConsumer(consumer_group, selected_topic string) *kgo.Client {
    // step 1: inisialisasi kafka client configuration
    brokers :=  []string{os.Getenv("KAFKA_CLIENT")}		// address message broker (kalau local ya localhost:9092)
    groupId :=  consumer_group				        // group id untuk consumer, bisa diisi sesuai keinginan (consumer group id)
    topic :=    selected_topic				        // nama topic yang akan di-consume
    
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
```

2. Cara buat producer (juga ada di `config/kafka.go`)
```go
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
```

Nah sekedar menyimpulkan, pembuatan producer dan consumer sedikit memiliki perbedaan dimana consumer tersebut di-inisialisasi dengan `topic` dan `consumer-groupId` sementara producer hanya `brokers`-nya saja.
Lalu, bagaimana producer bisa tau pengiriman message ke topic apa? Nah, kita dapat membuat producer menjadi `Generic Type`.
```go
// step 1: buat struct Producer dengan generic type T yang merupakan model.Event -> cara pakai-nya Producer[UserEvent]
/*
	Keuntungannya adalah type safety (keamanan tipe).
	Kita tidak akan bisa secara tidak sengaja mengirim OrderCreatedEvent
	menggunakan producer yang dirancang khusus untuk UserRegisteredEvent.

	Ibaratnya seperti ini:
	- Jika kita punya `Producer[UserEvent]`, kita hanya bisa mengirim `UserEvent`.
	- Jika kita punya `Producer[OrderEvent]`, kita hanya bisa mengirim `OrderEvent`.
	Jadi, kita tidak akan pernah salah mengirim event yang salah ke topic yang salah.
*/
type Producer[T model.Event] struct {
	Producer *kgo.Client
	Topic    string
}

// step 2: buat function buat dapet topic dari producer (getter ini optional, tapi bisa berguna untuk kedepannya, jadi ywd taro aja)
func (producer *Producer[T]) GetTopic() string {
	return producer.Topic
}

// step 3: buat function untuk publish event ke topic yang sudah ditentukan
func (producer *Producer[T]) Publish(event T) error {
	// step 4: convert event ke json
	value, err := json.Marshal(event)
	if err != nil {
		fmt.Printf("Error marshalling event: %v\n", err)
		return err
	}

	// step 5: buat record untuk dikirim ke topic
	record := &kgo.Record{
		Topic: producer.Topic,
		Key:   []byte(event.GetId()), // Gunakan ID dari event sebagai key
		Value: value,
	}

	// step 6: publish record ke topic
	ctx := context.Background()

	// Gunakan ProduceSync untuk memastikan bahwa message sudah berhasil dikirim sebelum melanjutkan
	// Ini akan mengembalikan error jika terjadi kesalahan saat mengirim message
	err = producer.Producer.ProduceSync(ctx, record).FirstErr()
	if err != nil {
		fmt.Printf("Error producing record: %v\n", err)
		return err
	}

	// step 7: return nil jika tidak ada error
	return nil
}
```
Nah, liat bagian ini aja ges:
```go
type Producer[T model.Event] struct {
	Producer *kgo.Client
	Topic    string     // ini jadi alasan kenapa kita tidak menentukan topic pada saat inisialisasi producer
}
```
Contoh implementasinya kek begini:
```go
type UserProducer struct {
	Producer *Producer[*model.UserEvent] 	// Gunakan generic type untuk UserEvent
}

func NewUserProducer(producer *kgo.Client) *UserProducer {
	return &UserProducer{
		Producer: &Producer[*model.UserEvent]{
			Producer: producer,
			Topic:    "users",
		},
	}
}
```
Artinya, producer dengan UserEvent akan mengirim topic ke "users". Dah, intinya begitu aja.

Sekarang, buat `consumer.go`:
```go
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
						}
					}
				})
		}
	}
}

// Notes: kalau bingung/lupa, lihat ke cmd/worker/main.go, disitu ada contoh implementasi untuk consume message dari topic yang sudah ditentukan
// Topic ditentukan pada saat membuat consumer sehingga tidak perlu lagi menentukan topic disini
```

Dan as always, dia akan digunakan di dalam:
```go
package messaging

import (
	"encoding/json"
	"fmt"
	"kafka-confluent-pzn/model"

	"github.com/twmb/franz-go/pkg/kgo"
)

type UserConsumer interface {
	Consume(record *kgo.Record) error
}

type UserConsumerImpl struct {
}

func NewUserConsumer() UserConsumer {
	return &UserConsumerImpl{}
}

func (userConsumer *UserConsumerImpl) Consume(record *kgo.Record) error {
	// step 1: convert record value to model user event (biar bentuknya sesuai)
	userEvent := model.UserEvent{}
	err := json.Unmarshal(record.Value, &userEvent)
	if err != nil {
		fmt.Printf("Error unmarshalling record value: %v\n", err)
		return err
	}

	// step 2: print user event (untuk tutorial/debug purpose saja) bisa loncat ke step 3 kalau app-nya udah ke defined
	fmt.Printf("Received user event: ID=%s, Name=%s, CreatedAt=%d, UpdatedAt=%d\n",
		userEvent.ID, userEvent.Name, userEvent.CreatedAt, userEvent.UpdatedAt)

	// step 3: process user event (disini bisa ditambahkan logic untuk menyimpan ke database atau melakukan proses lainnya)
	// misalnya simpan ke database atau melakukan proses lainnya
	// ...

	// step 4: return nil jika tidak ada error
	return nil
}
```
Nah, karena tadi `producer.go` dan `consumer.go` ada nyinggung Event-event-an. Harusnya 2 file dibawah ini yang harus dibuat terlebih dahulu baru abis tu `producer.go` dan `consumer.go`
- `event.go`:
```go
package model

type Event interface {
	GetId() string // ini bakal berperan sebagai key untuk message yang dikirim ke topic
}
```

- `user_event.go`:
```go
package model

type UserEvent struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

func (u *UserEvent) GetId() string {
	return u.ID
}
```

## Quick Start

Darisini harusnya step-by-stepnya kalau untuk implementasi udah jelas yaitu:
1. Buat file **`event.go`** dan **`user_event.go`** di dalam directory **`model/`**.
2. Buat file **`consumer.go`** dan define handler function-nya di dalam **`user_consumer.go`**, function handler ini akan di define dan di masukkan ke dalam worker pada **`cmd/worker/main.go`**. File `consumer.go` dan `user_consumer.go` akan di masukkan ke dalam directory **`delivery/messaging/`**
3. Buat file **`producer.go`** dan define producer spesifik beserta topic yang akan dia jadikan tempat untuk publish pada **`user_producer.go`**. Function publish akan di gunakan dalam directory `/usecase` jadi gaush dipikirin dulu buat sekarang. Kedua file ini akan ada di dalam directory **`gateway/messaging/`**.
4. Buat file configurasi **`kafka.go`** yang isinya adalah function untuk initialize consumer dan producer di **`config/kafka.go`**. Take notes, langkah ke-4 ini boleh jadi langkah pertama kalau mau. Biasanya juga bisa langsung copas karena configuration dan sebagainya akan di define lagi di `cmd/worker/main.go` sewaktu pemanggilan function-function New... ini.


### How to use:

Contoh isi sewaktu pemanggilan dan inisialisasi messaging, beserta consumer di `cmd/worker/main.go`:
```go
package main

import (
	"context"
	"kafka-confluent-pzn/config"
	"kafka-confluent-pzn/delivery/messaging"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunUserConsumer(ctx context.Context) {
	userConsumer := config.NewConsumer("users-group", "users")
	userHandler := messaging.NewUserConsumer()
	messaging.ConsumeTopic(
		ctx,                 // context (nothing to explain)
		userConsumer,        // Consumer-nya
		userHandler.Consume, // Handler function to process messages that were fetches on messaging.ConsumeTopic() method
	)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go RunUserConsumer(ctx)

	// Template (should learn Goroutines first):
	terminateSignals := make(chan os.Signal, 1)
	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	stop := false
	for !stop {
		select {
		case <-terminateSignals:
			cancel()
			stop = true
		}
	}

	time.Sleep(5 * time.Second) // wait for all consumers to finish processing
}
```
Jadi, inget yang dipanggil di `main.go` itu adalah bagian `messaging` dan `consumer`.


Nah, kalau bagian producer di panggil dimana dong? Nah ini bakal di panggil di use-case dimana kalau mau lanjut ke use-case, lu harus buat repositorynya dulu. Yang jelas, code dibawah ini bakal jadi referensi aja:
Fokus ke bagian yang di highlight aja ya setelah code block ini:
```go
func (c *UserUseCase) Create(ctx context.Context, request *model.RegisterUserRequest) (*model.UserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	total, err := c.UserRepository.CountById(tx, request.ID)
	if err != nil {
		c.Log.Warnf("Failed count user from database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if total > 0 {
		c.Log.Warnf("User already exists : %+v", err)
		return nil, fiber.ErrConflict
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Log.Warnf("Failed to generate bcrype hash : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	user := &entity.User{
		ID:       request.ID,
		Password: string(password),
		Name:     request.Name,
	}

	if err := c.UserRepository.Create(tx, user); err != nil {
		c.Log.Warnf("Failed create user to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if c.UserProducer != nil {
		event := converter.UserToEvent(user)
		c.Log.Info("Publishing user created event")
		if err = c.UserProducer.Send(event); err != nil {
			c.Log.Warnf("Failed publish user created event : %+v", err)
			return nil, fiber.ErrInternalServerError
		}
	} else {
		c.Log.Info("Kafka producer is disabled, skipping user created event")
	}

	return converter.UserToResponse(user), nil
}
```
Bagian yg di**highlight** itu yang ini:
```go
if c.UserProducer != nil {
	event := converter.UserToEvent(user)
	c.Log.Info("Publishing user created event")
	if err = c.UserProducer.Send(event); err != nil {
		c.Log.Warnf("Failed publish user created event : %+v", err)
		return nil, fiber.ErrInternalServerError
	}
} else {
	c.Log.Info("Kafka producer is disabled, skipping user created event")
}
```
Lho, kalo gitu producer-nya di define dimana? Inget aja kalau kita buat ini pakai Factory Pattern, jadi otomatis Producer ada di bagian atas dan use-case ini akan di panggil pada bagian `delivery/http/routes/route.go`:
```go
type UserUseCase interface {
	Create(ctx context.Context, request *model.RegisterUserRequest) (*model.UserResponse, error)
	// other function to implement...
}

type UserUseCaseImpl struct {
    DB                  *sql.DB
    RedisClient         *redis.Client
    Validate            *validator.Validate
    UserRepository      *repository.UserRepository
    UserProducer        *messaging.UserProducer // di define disini
}

func NewUserUseCase( // Remember factory pattern, ini akan ada di routes.go
	db *gorm.DB, 
	validate *validator.Validate, 
	userRepository *repository.UserRepository,
	userProducer *messaging.UserProducer
	) *UserUseCase {
    return &UserUseCase{
        DB:             db,
        Validate:       validate,
        UserRepository: userRepository,
        UserProducer:   userProducer,
    }
}
```
