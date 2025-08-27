package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	defer client.Close()

	// Ini akan dijalankan secara bulk (bukan satu per satu)
	result, err := client.Pipelined(context.Background(), func(pipeliner redis.Pipeliner) error {
		pipeliner.SetEx(context.Background(), "users", "William", time.Second*5)
		pipeliner.SetEx(context.Background(), "address", "Jakarta", time.Second*5)
		return nil
	})
	if err != nil {
		panic(err)
	}

	// cek hasil
	fmt.Println(result) // <- hasil/respon2 dari request yang dieksekusi
	fmt.Println(client.Get(context.Background(), "users").Val())
	fmt.Println(client.Get(context.Background(), "address").Val())
}
