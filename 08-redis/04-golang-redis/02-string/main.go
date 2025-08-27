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

	// set string data structures in redis example:
	client.SetEx(context.Background(), "name", "William Theodorus", time.Second*3)

	// get string data via key in redis
	result, err := client.Get(context.Background(), "name").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	// check the data if the key are expired
	time.Sleep(time.Second * 5)
	result, err = client.Get(context.Background(), "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
