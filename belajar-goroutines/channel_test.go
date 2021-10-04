package belajargoroutines_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreatChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "calvin wijay"
		fmt.Println("selesai mengirim data string")
	}()
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// channel as parameter
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "calvin wijaya"
}
func TestParameterChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "calvin wijaya"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)
	time.Sleep(5 * time.Second)
}

// buffered channel
func TestBufferChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "calvin"
		channel <- "wijaya"
		channel <- "phil"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()
	fmt.Println("done")
	time.Sleep(2 * time.Second)
}

// range channel (menerima data terumeneruts dgn looping)

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("perulangan ke ;", strconv.Itoa(i))
		}
		close(channel)
	}()
	for data := range channel {
		fmt.Println("menerima data", data)
	}
	fmt.Println("selesai")
}

// select channel

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	for counter := 0; counter <= 2; counter++ {
		select {
		case data := <-channel1:
			fmt.Println("data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari channel 2", data)
			counter++
		default:
			fmt.Println("nugngu")
		}
	}
}

//	counter := 0
//	for {
//		select {
//		case data := <-channel1:
//			fmt.Println("data dari channel 1", data)
//			counter++
//		case data := <-channel2:
//			fmt.Println("data dari channel 2", data)
//			counter++
//		}
//		if counter == 2 {
//			break
//		}
//	}
