package main

import (
	"fmt"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine() {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i) // Try remove the 'go' keywords, and with 'go' keywords,
		// With no 'go' keywords, the number will be sequential
		// However if the 'go' keywords are used, then the number will be unsequential since Goroutines run it without waiting for each other.
	}

	// Add some timer so then our application not going to stop immediately
	time.Sleep(5 * time.Second)
}

func main() {
	//go RunHelloWorld()
	//fmt.Println("Ups...")
	//
	//time.Sleep(1 * time.Second)
	TestManyGoroutine()
}
