# Golang Redis

<ol>
    <li>Redis adalah salah satu In Memory Database (Database pake RAM) yang digunakan untuk Cache, Rate Limiter, dsbnya.</li>
    <li>Untuk pake Redis di Golang, kita harus download di sini: https://github.com/redis/go-redis </li>
    <li>Nanti tinggal ikutin dokumentasinya atau run <b>go get github.com/redis/go-redis/v9</b>.</li>
</ol>

## Client

<ol>
    <li>Hal pertama yang perlu dilakukan jika ingin menggunakan redis dari Golang, adalah membuat koneksi ke Redisnya.</li>
    <li>Untuk membuat koneksi ke Redis, kita perlu membuat object redis.Client</li>
    <li>Kita bisa menggunakan function redis.NewClient(redis.Options)</li>
    <li>Kita bisa tentukan konfigurasi menggunakan redis.Options</li>
</ol>

Contoh:  
```go
var client = redis.NewClient(&redis.Options{
	Addr:   "localhost:6379",
	DB:     0,
})
```

## Command

<ol>
    <li>Golang Redis sangat mudah digunakan karena semua command Redis CLI bisa kita gunakan sebagai method di Golang.</li>
    <li>Sebagai contoh kalau kita pake command ZADD di Redis CLI, disini kita bisa pake client.Zadd()</li>
    <li>Tiap Command di Redis bisa dilakukan di Client, dengan format PascalCase, misalnya Ping(), Set(), Get(), dsbnya.</li>
    <li>Karena ini Golang, jadi parameter yang paling sering di pass adalah <b>Context</b>.</li>
    <li>Karena Redis ini adalah pair database (key-value), kita bisa mengambil value dari sebuah key dengan method <b>.Result()</b></li>
</ol>

## String

<ol>
    <li>Struktur data ini adalah struktur data yang paling sering digunakan di Redis.</li>
    <li>Sama kek di Redis CLI, commandnya juga sama cuman bedanya disini dalam bentuk method aja, seperti: <b>Set(), SetEx(), Get(), MGet(), dsbnya.</b></li>
</ol>

Contoh code: [code](02-string/main.go)

## List

Nah, kalau List kek gini nih: [code](03-list/main.go)

## Set

Nah, sama kek di CLI, di CLI, kita biasa pake SADD untuk add value ke dalam set, dan pakai SCARD untuk melihat total/panjang sebuah Set, dan kadang juga menggunakan SMEMBERS untuk melihat isi dari set tersebut. 
Nah, di Golang juga sama, kita pakai Command-command tersebut dalam bentuk method. Dimana untuk setiap methodnya, kita mempass context sebagai parameter tambahan.
  
Contoh kode: [code](04-set/main.go)

## Sorted Set (ZSET)

Nah, ini sekarang lebih ke basic aja karena sebenarnya kita udah implementasi pembuatan Rate Limiter di Go menggunakan Sorted Set ini dengan memanfaatkan ZADD, ZREMRANGEBYSCORE, dan ZCARD.
  
Contoh kode: [code](05-zset/main.go)

## Hash

Nah, kalau Hash juga kurang lebih sama kek di CLI dimana biasa kita pake HSET, HGETALL, dsbnya.

Contoh kode: [code](06-hash/main.go)

## Pipeline

<ol>
    <li>Pipeline digunakan untuk mengirim beberapa perintah secara langsung tanpa harus menunggu respon dari perintah tersebut secara satu per satu.</li>
    <li>Di Golang kita bisa gunakan method <b>Client.Pipelined(callback)</b></li>
    <li>Di dalam Callback Function, kita bisa melakukan semua command yang akan dijalankan.</li>
</ol>

Contoh kode: [code](07-pipeline/main.go)

## Transaction

<ol>
    <li>Kita tau bahwa menggunakan Redis bisa melakukan transaction dengan menggunakan perintah MULTI dan COMMIT.</li>
    <li>Namun, harus dalam koneksi yang sama.</li>
    <li>Karena Golang Redis melakukan maintain connection pool secara internal, jadi kita tidak bisa dengan mudah menggunakan MULTI dan COMMIT menggunakan redis.Client</li>
    <li>Kita harus menggunakan function TxPipelined(), dimana parameternya bisa berupa Callback function yang berisikan function-function yang ingin di execute.</li>
    <li>Atau bisa juga kita kosongkan parameternya dan ambil return-value dari <b>TxPipeline()</b> lalu gunakan return value tersebut sebagai client (instead pake variable berisikan redis.Client).</li>
</ol>

Contoh kode: [code](08-transaction/main.go)

#### Conclusions

<ol>
    <li>Jadi ada 2 cara dimana cara pertama itu pake method <b>TxPipelined(callback)</b> dan functionnya kita masukin ke dalam parameternya.</li>
    <li>Dan cara kedua yaitu pake method <b>TxPipeline()</b> tanpa parameter apapun dan jalankan return-valuenya seolah-olah sebagai redis.Client</li>
    <li>Ingat, cara kedua mengharuskan kita untuk menjalankan perintah Exec diakhir.</li>
</ol>

## Stream

#### Publish Stream (XADD)

```go
func PublishStream(ctx context.Context, client *redis.Client) {
    for i := 0; i < 10; i++ {
        if err := client.XAdd(ctx, &redis.XAddArgs{
            Stream: "members",
            Values:  map[string]interface{}{
                "name": "William",
                "address": "Indonesia",
			},
        }).Err(); err != nil {
            panic(err)
        }
    }
}
```

Notes:<br/>
<ol>
    <li>value yang ditambahkan harus dalam bentuk map yang udah dibuat yaitu redis.XAddArgs{}</li>
    <li>Dan tiap operasi XADD akan return sebuah error.</li>
</ol>

#### Create Consumer (XGROUPCREATE + XGROUPCREATECONSUMER)

```go
func CreateConsumerGroup(ctx context.Context, client *redis.Client) {
	client.XGroupCreate(ctx, "members", "group-1", "0")                     // context, stream-names, consumer-group, start (--from-beginning)
	client.XGroupCreateConsumer(ctx, "members", "group-1", "consumer-1")    // context, stream-names, consumer-group, consumer-name
	client.XGroupCreateConsumer(ctx, "members", "group-1", "consumer-2")    // context, stream-names, consumer-group, consumer-name
}
```

#### Get Stream (Consume Data via XREADGROUP)

```go
func GetStream(ctx context.Context, client *redis.Client) {
	result := client.XReadGroup(ctx, &redis.XReadGroupArgs{
	    Group:      "group-1",
		Consumer:   "consumer-1",
		Streams:    []string{"members", ">"},   // > indicates read unread messages
		Count:      2,                          // indicates how many messages we want to read
		Block:      time.Second * 5,            // block indicates the waiting time for consumer to read a data if the broker is empty
    }).Val()
	
	for _, stream := range result {
		for _, message := range stream.Messages {
		    fmt.Println(message.Values)	
        }
    }   
}
```

Contoh kode: [code](09-stream/main.go)

## PubSub (Stream without Consumer Group)

#### Subscribe PubSub

```go
func SubscribePubSub(ctx context.Context, client *redis.Client) {
	// subscribe ke topic channel-1
	pubSub := client.Subscribe(ctx, "channel-1")
	
	// close kalau udah selesai
	defer pubSub.Close()
	
	// read message via method ReceiveMessage() to reading with unlimited waittime / ReceiveTimeOut()  to read with timeout.
	// message can be read via Payload
	for i := 0; i < 10; i++ {
		message, _ := pubSub.ReceiveMessage(ctx)
		fmt.Println(message.Payload)
    }
}
```

Notes: <br/>
<ol>
    <li>Jadi kalau pake <b>ReceiveMessage()</b> dia akan selalu menunggu messagenya bahkan ketika log topicnya sedang kosong.</li>
    <li>Kalau <b>ReceiveTimeout()</b> itu ada batas waktunya kalau misal ga dapet message dalam waktu yang sudah ditentukan, maka akan di stop consumernya.</li>
</ol>

#### Publish PubSub

```go
func PublishPubSub(ctx context.Context, client *redis.Client) {
	for i := 0; i < 10; i++ {
	    client.Publish(ctx, "channel-1", "Hello-"+strconv.Itoa(i))	
    }
}
```

Contoh kode: [code](10-pubsub/main.go)