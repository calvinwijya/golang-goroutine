package belajargoroutines_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "new"
		},
	}
	//wait := sync.WaitGroup{}
	pool.Put("calvin")
	pool.Put("vio")
	pool.Put("acin")

	for i := 0; i < 10; i++ {

		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()

	}
	time.Sleep(11 * time.Second)
	fmt.Println("tes berhasil")
}
