package main

import (
	"fmt"
	"time"
)

func main() {
	var x = 0
	for i := 1; i <= 1000; i++ {
		// akan terdapat 1000 goroutine
		go func() {
			// setiap goroutine akan menjalankan penjumlahan sebanyak 100 kali
			for j := 1; j <= 100; j++ {
				x = x + 1
			}
		}()

		// berarti setiap goroutine akan menambahkan nilai x sebanyak 100
		// hingga akhir dari nilai x harusnya 1000 * 100 yaitu 100000
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter: ", x)
}
