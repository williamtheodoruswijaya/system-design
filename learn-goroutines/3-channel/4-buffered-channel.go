package main

import "fmt"

func main() {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "First Data"
	channel <- "Second Data"
	channel <- "Third Data"
	//channel <- "Fourth Data" // If this are uncommented, then our channel will create a block mechanism since our buffer capacity is 3 and won't reach the end of the program.

	// we also can wrap the sending data into a goroutine so the sending will be done without having to wait the above line to finish sending the data
	// uncomment code below to test:
	//go func() {
	//	channel <- "First Data"
	//	channel <- "Second Data"
	//	channel <- "Third Data"
	//}()
	fmt.Println("Done") // This will be run because channel didn't wait for the receiver since the buffer size are 3 so as long as there are no more than 3 value in channel, the channel won't await.
}
