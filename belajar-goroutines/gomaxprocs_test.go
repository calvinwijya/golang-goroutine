package belajargoroutines_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestCPU(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(5 * time.Second)
			group.Done()
		}()

	}

	totalCPU := runtime.NumCPU()
	fmt.Println("CPU ", totalCPU)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("thread ", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("goroutine ", totalGoroutine)

	group.Wait()
}

func TestChangeThread(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(5 * time.Second)
			group.Done()
		}()

	}

	totalCPU := runtime.NumCPU()
	fmt.Println("CPU ", totalCPU)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("thread ", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("goroutine ", totalGoroutine)

	group.Wait()
}
