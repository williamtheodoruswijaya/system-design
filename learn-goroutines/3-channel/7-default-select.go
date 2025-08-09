package main

import "fmt"

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	counter := 0
	for {
		select {
		case <-channel1:
			fmt.Println("Data dari Channel 1: ", <-channel1)
			counter++
		case <-channel2:
			fmt.Println("Data dari Channel 2: ", <-channel2)
			counter++
		default: // sebelumnya kalau gaada ini, kita akan terkena deadlock error, tapi kalau sekarang ada ini ya cuman infinite loop aja tapi intinya ini akan dijalankan terlepas dari sebuah data akan di receive atau tidak (use-case-nya consume di kafka)
			fmt.Println("Waiting for data...")
		}
		if counter == 2 {
			break
		}
	}
}
