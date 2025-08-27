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

	// Add value to Hash (HSET)
	client.HSet(context.Background(), "users", "name", "John")
	client.HSet(context.Background(), "users", "age", "18")
	client.HSet(context.Background(), "users", "job", "Student")

	// Get all values
	user := client.HGetAll(context.Background(), "users").Val()

	fmt.Println(user)
}
