package main

import (
	"fmt"
	"sync"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1, user2 *UserBalance, amount int) {
	// step 1: lock user 1 &
	user1.Lock()
	fmt.Println("Lock user 1: ", user1.Name)

	// step 2: kurangi balance dari user 1
	user1.Change(-amount)

	// add latency sedikit
	time.Sleep(1 * time.Second)

	// step 3: lock user 2
	user2.Lock()
	fmt.Println("Lock user 2: ", user2.Name)

	// step 4: tambahin balance dari user 2
	user2.Change(amount)

	// add latency sedikit
	time.Sleep(1 * time.Second)

	// step 5: unlock user 1 kalau udah beres
	fmt.Println("Unlock user 1: ", user1.Name)
	user1.Unlock()

	// step 6: unlock user 2 kalau udah beres
	fmt.Println("Unlock user 2: ", user2.Name)
	user2.Unlock()
}

func main() {
	user1 := &UserBalance{
		Name:    "Eko",
		Balance: 100000,
	}
	user2 := &UserBalance{
		Name:    "Budi",
		Balance: 100000,
	}

	go Transfer(user1, user2, 1000)
	go Transfer(user2, user1, 1000)

	time.Sleep(5 * time.Second)

	fmt.Println("User1 Balance: ", user1.Balance)
	fmt.Println("User2 Balance: ", user2.Balance)
}

/*
	Kenapa hasilnya jadi:
		Lock user 1:  Eko
		Lock user 1:  Budi
		User1 Balance:  99000
		User2 Balance:  99000

	Lihat, tidak pernah terjadi Lock user 2 antara Goroutine pertama maupun kedua,
	karena Goroutine kedua langsung lock Budi di saat Eko masih di Lock,
	akibatnya lock user 2 tidak pernah terjadi karena untuk lock user Budi,
	user Budi harus di unlock terlebih dahulu oleh goroutine A, sementara untuk lock user Eko,
	harus di unlock terlebih dahulu oleh Goroutine B.
	kedua lock masih saling menunggu.
	Akibatnya, terjadilah Deadlock.
*/
