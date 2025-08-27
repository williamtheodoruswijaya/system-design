package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func caraPertama(ctx context.Context, client *redis.Client) {
	client.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.SetEx(ctx, "name", "William", time.Second*5)
		pipe.SetEx(ctx, "age", 18, time.Second*5)
		return nil
	})
}

func caraKedua(ctx context.Context, client *redis.Client) {
	pipe := client.TxPipeline()
	// Cara ini mengharuskan kita untuk menggunakan perintah Exec diakhir
	defer func() {
		_, err := pipe.Exec(ctx)
		if err != nil {
			panic(err)
		}
	}()

	pipe.SetEx(ctx, "name", "William", time.Second*5)
	pipe.SetEx(ctx, "age", 18, time.Second*5)
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	defer client.Close()

	// Ganti method antara cara pertama dan kedua
	caraKedua(context.Background(), client)

	// Cek hasil:
	fmt.Println(client.Get(context.Background(), "name").Val())
	fmt.Println(client.Get(context.Background(), "age").Val())
}
