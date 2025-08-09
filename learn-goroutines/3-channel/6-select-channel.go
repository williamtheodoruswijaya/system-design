package main

import (
	"fmt"
	"strconv"
)

func SendDataIterate(channel chan<- string) {
	for i := 1; i <= 10; i++ {
		channel <- "Perulangan ke-" + strconv.Itoa(i)
	}
}

func main() {
	channel1 := make(chan string)
	defer close(channel1)

	channel2 := make(chan string)
	defer close(channel2)

	go SendDataIterate(channel1)
	go SendDataIterate(channel2)

	counter := 0
	for {
		select {
		// intinya ketika ada
		case data := <-channel1:
			fmt.Println("Data dari Channel 1: ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2: ", data)
			counter++
		}
		if counter == 2 { // ini buat jaga-jaga aja biar for loopnya ga infinite
			break
		}
	}

	// Kalau ga pake select terus for range satu"kan cape ya baca dari channel 1 dulu abis tu channel 2
}
