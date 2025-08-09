package main

import (
	"context"
	"kafka-confluent-pzn/config"
	"kafka-confluent-pzn/delivery/messaging"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunUserConsumer(ctx context.Context) {
	userConsumer := config.NewConsumer("users-group", "users")
	userHandler := messaging.NewUserConsumer()
	messaging.ConsumeTopic(
		ctx,                 // context (nothing to explain)
		userConsumer,        // Consumer-nya
		userHandler.Consume, // Handler function to process messages that were fetches on messaging.ConsumeTopic() method
	)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go RunUserConsumer(ctx)

	// Template (should learn Goroutines first):
	terminateSignals := make(chan os.Signal, 1)
	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	stop := false
	for !stop {
		select {
		case <-terminateSignals:
			cancel()
			stop = true
		}
	}

	time.Sleep(5 * time.Second) // wait for all consumers to finish processing
}
