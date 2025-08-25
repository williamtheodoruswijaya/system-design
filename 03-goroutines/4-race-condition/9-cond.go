package main

import (
	"fmt"
	"sync"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	cond.L.Lock()
	defer cond.L.Unlock()
	defer group.Done()
	group.Add(1)

	// step 1: tentuin setelah locking kita Goroutine-nya boleh jalan atau ga? Normalnya kondisi dia jalan itu diatur oleh function Signal()
	cond.Wait()

	// step 2: progress-nya
	fmt.Println("Waiting Condition Complete... - ", value)
}

func main() {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal() // kirim signal ke 1 goroutine yang sedang di lock kalau dia boleh jalanin code dibawahnya.
		}
	}()

	group.Wait()
}
