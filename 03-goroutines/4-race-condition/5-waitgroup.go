package main

import (
	"fmt"
	"sync"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	// step 1: pastikan selalu Done() ketika sebuah progress selesai.
	defer group.Done()

	// step 2: add goroutines-nya
	group.Add(1)

	// Contoh progress-nya:
	fmt.Println("Running a goroutine")
	time.Sleep(1 * time.Second)
}

func main() {
	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go RunAsynchronous(group)
	}

	// buat function .Wait() yang membuat sebuah goroutine dijalankan terlebih dahulu sampai value dalam add() == 0 baru jalankan code dibawahnya.
	group.Wait() // instead pake Thread.sleep(5 * time.Second)
	fmt.Println("Complete...")
}
