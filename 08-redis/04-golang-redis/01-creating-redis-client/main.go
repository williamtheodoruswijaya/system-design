package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	var client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	defer client.Close()

	testConnection, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(testConnection)
}
