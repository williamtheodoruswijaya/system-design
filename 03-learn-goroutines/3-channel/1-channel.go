package main

import "fmt"

func main() {
	// 1. initialize a channel
	channel := make(chan string)

	// 2. sending data to channel
	channel <- "Hello"

	// 3. receiving data in channel
	data := <-channel

	// 4. optional: we can also send the data into a parameter
	fmt.Println(<-channel)

	// 5. always close channel to avoid memory leak
	close(channel)
}
