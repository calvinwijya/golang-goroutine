package belajargoroutines_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {

	x := 0
	var lock sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				lock.Lock()
				x += 1
				lock.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("counter ke ", x)
}

// READ WRITE MUTEX
type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance += amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("total balance = ", account.Balance)
}

// DEADLOCK

type UserBalance struct {
	Mutex   sync.RWMutex
	name    string
	balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.balance += amount
}

func Transfer(out *UserBalance, in *UserBalance, amount int) {

	out.Mutex.Lock()
	fmt.Println("lock", out.name)
	out.Change(-amount)

	in.Mutex.Lock()
	fmt.Println("lock", in.name)
	in.Change(amount)

	out.Unlock()
	in.Unlock()
}

func TestDeadLock(t *testing.T) {
	//group := sync.WaitGroup{}
	user1 := UserBalance{
		name:    "calvin",
		balance: 1000000,
	}
	user2 := UserBalance{
		name:    "vio",
		balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	time.Sleep(5 * time.Second)

	fmt.Println("user: ", user1.name, "balance: ", user1.balance)
	fmt.Println("user: ", user2.name, "balance: ", user2.balance)

}
