package main

import (
	"context"
	"fmt"
	"runtime"
)

func CreateCounter() chan int {
	destination := make(chan int)

	// goroutines to fill a channel
	go func() {
		defer close(destination)
		counter := 1

		// This's an infinite loop
		for {
			destination <- counter
			counter++
		}
	}()

	// return the channel
	return destination
}

func CorrectCreateCounter(ctx context.Context) chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done(): // stop the goroutines when the application are finished.
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()
	return destination
}

func main() {
	// check how many goroutines are running before and after CreateCounter()
	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())

	// Correct way:
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent) // The cancel function here, if it's called, it will make ctx.Done() return a value making the goroutines stop

	// Run the goroutine
	destination := CorrectCreateCounter(ctx)

	// Check how many value are in the channel
	for n := range destination {
		fmt.Println("Counter: ", n)
		if n == 10 {
			break
		}
	}
	cancel() // makes the ctx.Done() returns a value (stopping the running Goroutines)

	// check how many goroutines are running before and after CreateCounter()
	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())

	// The problem here is that after the function is finish, the goroutine are still running
	// This's a goroutine leak problem.
	// To handle this, we can use context.WithCancel() to cancel a goroutine when the process are finished or the goroutines are no longer used.
	// The function should have context as a parameter and inside the infinite loop, insert a select and default channel
	// And we just need to send the context into the parameter
}
