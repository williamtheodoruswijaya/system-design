package main

import (
	"fmt"
	"sync"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func main() {
	var once sync.Once
	var group sync.WaitGroup

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println(counter) // Hasilnya pasti 1 karena function ini hanya di jalankan 1 kali oleh 1 Goroutines dari 100 Goroutines yang ada.
}
