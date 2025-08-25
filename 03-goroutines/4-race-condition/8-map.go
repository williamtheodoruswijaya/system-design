package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var data sync.Map
	var addToMap = func(value int) {
		data.Store(value, value)
	}

	for i := 0; i < 100; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()

			addToMap(i)
		}()
	}

	wg.Wait()
	fmt.Println("Adding Data Complete...")
	data.Range(func(key, value interface{}) bool { // return bool ini ibaratnya indikasi kalau kita mau lanjut ke iterasi selanjutnya.
		fmt.Println(key, value)
		return true // return true kalau kita mau lanjut ke iterasi selanjutnya.
	})
}
