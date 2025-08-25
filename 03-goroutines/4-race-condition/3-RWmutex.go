package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	RWMutex sync.RWMutex // RWMutex di dalam struct ini adalah cara kita pakai RWMutex khusus struct yang akan diakses oleh berbagai Goroutines
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	// step 1: lock goroutines ketika ingin write
	account.RWMutex.Lock()

	// step 2: unlock kalau udah beres
	defer account.RWMutex.Unlock()

	// step 3: logic disini...
	account.Balance += amount
}

func (account *BankAccount) GetBalance() int {
	// step 1: lock goroutines ketika ingin read (pakai RLock())
	account.RWMutex.RLock()

	// step 2: unlock kalau udah beres Read-nya. (takut terjadi Race Condition saat Read)
	defer account.RWMutex.RUnlock()

	// step 3: kode dibawah akan dijalankan pada saat Lock
	return account.Balance
}

func main() {
	account := &BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance: ", account.GetBalance())
}
