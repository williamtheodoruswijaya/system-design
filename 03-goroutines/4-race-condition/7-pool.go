package main

import (
	"fmt"
	"sync"
)

var pool sync.Pool
var wg sync.WaitGroup

func main() {
	pool.Put("This's data 1")
	pool.Put("data 2")
	pool.Put("data 3")

	for i := 0; i < 10; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()

			data := pool.Get() // Ketika di get(), maka datanya akan hilang dari Pool
			fmt.Println(data)
			pool.Put(data) // Kalau udah pake data-nya, kita return ke Pool kembali
		}()
	}

	wg.Wait()
}
