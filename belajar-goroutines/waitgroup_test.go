package belajargoroutines_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {

	group.Add(1)
	fmt.Println("hello world")

	time.Sleep(1 * time.Second)
	group.Done()
}

func TestAsynchronous(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}
	group.Wait()

	fmt.Println("selesai")
}

var counter1 = 0

func OnlyOnce1() {
	counter++
}
func TestOnce1(t *testing.T) {
	//	once := sync.Once{}
	mutex := sync.Mutex{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		mutex.Lock()
		OnlyOnce()
		mutex.Unlock()
		group.Done()

	}
	group.Wait()
	fmt.Println("counter: ", counter1)
}
