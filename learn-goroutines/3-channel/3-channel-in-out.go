package main

import (
	"fmt"
	"time"
)

func SendData(channel chan<- string, data string) {
	time.Sleep(1 * time.Second)
	channel <- data
}

func ReceiveData(channel <-chan string, data string) {
	data = <-channel
	fmt.Println(data)
}

func main() {
	channel := make(chan string)
	defer close(channel)

	go SendData(channel, "Hello World")
	go ReceiveData(channel, "Hello World")

	time.Sleep(3 * time.Second)
}
