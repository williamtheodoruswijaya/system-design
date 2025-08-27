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

	PublishStream(context.Background(), client)
	CreateConsumerGroup(context.Background(), client)
	GetStream(context.Background(), client)

	// Kalau udah delete keysnya
	client.Del(context.Background(), "members")
}

func PublishStream(ctx context.Context, client *redis.Client) {
	for i := 0; i < 10; i++ {
		if err := client.XAdd(ctx, &redis.XAddArgs{
			Stream: "members",
			Values: map[string]interface{}{
				"name":    "William",
				"address": "Indonesia",
			},
		}).Err(); err != nil {
			panic(err)
		}
	}
}

func CreateConsumerGroup(ctx context.Context, client *redis.Client) {
	client.XGroupCreate(ctx, "members", "group-1", "0")                  // context, stream-names, consumer-group, start (--from-beginning)
	client.XGroupCreateConsumer(ctx, "members", "group-1", "consumer-1") // context, stream-names, consumer-group, consumer-name
	client.XGroupCreateConsumer(ctx, "members", "group-1", "consumer-2") // context, stream-names, consumer-group, consumer-name
}

func GetStream(ctx context.Context, client *redis.Client) {
	result := client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    "group-1",
		Consumer: "consumer-1",
		Streams:  []string{"members", ">"}, // > indicates read unread messages
		Count:    2,                        // indicates how many messages we want to read
		Block:    time.Second * 5,          // block indicates the waiting time for consumer to read a data if the broker is empty
	}).Val()

	for _, stream := range result {
		for _, message := range stream.Messages {
			fmt.Println(message.Values)
		}
	}
}
