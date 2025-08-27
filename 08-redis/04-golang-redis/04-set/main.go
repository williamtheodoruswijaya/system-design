package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	defer client.Close()

	// Adding a data to a set (SADD)
	client.SAdd(context.Background(), "users", "test-1")
	client.SAdd(context.Background(), "users", "test-1")
	client.SAdd(context.Background(), "users", "test-2")
	client.SAdd(context.Background(), "users", "test-2")
	client.SAdd(context.Background(), "users", "test-3")
	client.SAdd(context.Background(), "users", "test-3")

	// Get Length of the Set (SCARD)
	lengthSet := client.SCard(context.Background(), "users")
	fmt.Println(lengthSet)

	// Get each values of the Set (SMEMBERS)
	stringArr := client.SMembers(context.Background(), "users")
	fmt.Println(stringArr)
}
