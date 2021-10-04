package belajargoroutines_test

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("hellow world")
}

func TestCreatGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("oh no")

	time.Sleep(1 * time.Second)
}

/////////////////////////////////////////////////////

func DisplayName(number int) {
	fmt.Println("Display", number)
}

func TestLoopName(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayName(i)
	}
	time.Sleep(5 * time.Second)
}
