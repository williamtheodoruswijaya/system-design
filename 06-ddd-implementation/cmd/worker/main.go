package main

import (
	"06-ddd-implementation/internal/config"
	"06-ddd-implementation/internal/delivery/messaging"
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func RunUserConsumer(ctx context.Context) {
	userConsumer := config.NewConsumer("users-group", "users")
	userHandler := messaging.NewUserConsumer()
	messaging.ConsumeTopic(
		ctx,                 // context (use for sending stop signal, so ctx.Done() can return a value on ConsumeTopic)
		userConsumer,        // Consumer-nya
		userHandler.Consume, // Handler function to process messages that were fetches on messaging.ConsumeTopic() method
	)
}

func main() {
	// initialize waitGroup
	var wg sync.WaitGroup

	// initialize context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// run a consumer as a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		RunUserConsumer(ctx)
	}()

	// create a terminate signal so cancel() function can be called
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

	// wait for the goroutines to be complete
	wg.Wait()
}
