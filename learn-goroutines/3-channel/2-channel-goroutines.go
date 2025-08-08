package main

import (
	"fmt"
	"time"
)

func sendData(channel chan string, data string) {
	time.Sleep(1 * time.Second)
	channel <- data
}

func main() {
	// 1. create the channel (take this as a variable that will change it's value)
	channel := make(chan string)
	defer close(channel)

	// 2. send the data via goroutines
	go sendData(channel, "This's the data")

	// testing purposes...
	fmt.Println("I'm runned before sendData()")

	// 4. receive the data
	random_variable := <-channel

	// testing purposes
	fmt.Println("I'm waiting the data received...")

	// 5. line below 25 will be runned after the data are received...
	fmt.Println(random_variable)
}
