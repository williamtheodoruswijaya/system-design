package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex

	x := 0
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock() // tiap goroutines yang berjalan akan mencoba menjalan locking ini, leaving only 1 that able to lock and able to run the code under.
				x = x + 1
				mutex.Unlock() // after goroutines update X value, the mutex will be unlocked, giving chances for other goroutines to lock updating X value. (Mutex can only receive one lock from one goroutines only)
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(x)
}
