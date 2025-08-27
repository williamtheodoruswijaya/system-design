package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	defer client.Close()

	var wg sync.WaitGroup

	wg.Add(1)

	go SubscribePubSub(context.Background(), client)
	PublishPubSub(context.Background(), client)

	wg.Wait()
}

func SubscribePubSub(ctx context.Context, client *redis.Client) {
	// subscribe ke topic channel-1
	pubSub := client.Subscribe(ctx, "channel-1")

	// close kalau udah selesai
	defer pubSub.Close()

	// read message via method ReceiveMessage()
	// message can be read via Payload
	for i := 0; i < 10; i++ {
		message, _ := pubSub.ReceiveMessage(ctx)
		fmt.Println(message.Payload)
	}
}

func PublishPubSub(ctx context.Context, client *redis.Client) {
	for i := 0; i < 10; i++ {
		err := client.Publish(ctx, "channel-1", "Hello-"+strconv.Itoa(i)).Err()
		if err != nil {
			panic(err)
		}
	}
}
