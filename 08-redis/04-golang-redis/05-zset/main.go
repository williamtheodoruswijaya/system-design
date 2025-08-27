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

	// Inget kalau ZSet itu harus pake redis.Z{} buat masukin membernya (ini adalah enum dengan score dan value)
	client.ZAdd(context.Background(), "users", redis.Z{Score: 100, Member: "Ini urutan pertama"})
	client.ZAdd(context.Background(), "users", redis.Z{Score: 85, Member: "Ini urutan ketiga"})
	client.ZAdd(context.Background(), "users", redis.Z{Score: 95, Member: "Ini urutan kedua"})

	// Buat dapetin valuenya kita bisa pake ZRange
	fmt.Println(client.ZRange(context.Background(), "users", 0, -1).Val())

	// Buat dapetin jumlah valuenya
	fmt.Println(client.ZCard(context.Background(), "users").Val())

	// Buat pop dari ZSet (ZPOPMAX (buang head), ZPOPMIN (buang tail))

	client.Expire(context.Background(), "users", 10*time.Second)
	client.Del(context.Background(), "users")
}
