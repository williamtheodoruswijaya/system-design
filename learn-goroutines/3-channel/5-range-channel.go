package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	channel := make(chan string)
	defer close(channel) // intinya kalau gaada ini, maka akan error karena loopnya gaakan berhenti

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
	}()

	go func() {
		for data := range channel { // intinya kalau ga diclose, maka channel tidak akan pernah stop karena dia akan menerima data terus.
			fmt.Println(data)
		}
	}()

	time.Sleep(1 * time.Second)
}
