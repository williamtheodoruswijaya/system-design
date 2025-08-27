package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	var client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	defer client.Close()

	// Kalau di CLI kan RPUSH, RPOP, LPUSH, dsbnya, nah ini sama cuman dijadiin method aja.
	client.RPush(context.Background(), "name", "William Theodorus")
	client.RPush(context.Background(), "name", "Kirigaya Kazuto")
	client.RPush(context.Background(), "name", "Amano Hina")

	// Kalau di CLI, kita ambil value di List pake LRange / RRange
	result, err := client.LRange(context.Background(), "name", 0, -1).Result()
	if err != nil {
		panic(err)
	}

	for _, v := range result {
		fmt.Println(v)
	}

	// Nanti kalau mau Pop() juga kita bisa pake LPop(), RPop()
	// Kalau mau ngambil valuenya pake .Val()

	// Set Expires
	client.Expire(context.Background(), "name", 10*time.Second)
}
