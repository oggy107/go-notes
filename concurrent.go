package main

import (
	"fmt"
	"sync"
)

// using sync.WaitGroup is suitable for waiting for multiple goroutines
// to wait for single goroutine simple channel can be used

// using sync.WaitGroup for waiting
func greetConcurrently(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	// simulating some work
	i := 0
	sum := 0
	for i < 10000000000 {
		sum += i
		i++
	}

	fmt.Printf("Hello, %s!\n", name)
	fmt.Println("Concurrent function finished")
}

// Channels are a typed, thread-safe queue. Channels allow different goroutines to communicate with each other.
// using channel for waiting

// NOTE: be careful of the deadlocks. A deadlock is when a group of goroutines are all blocking so none of them can continue.

func greetConcurrentlyChannel(name string, ch chan string) {
	// simulating some work
	i := 0
	sum := 0
	for i < 10000000000 {
		sum += i
		i++
	}

	fmt.Printf("Hello, %s!\n", name)
	fmt.Println("Concurrent function finished. returning data into channel")

	// passing data into channel
	ch <- "done"
	// we can also close the channel (if sending multiple values to channel and recieving values in for loop to indicate that the the goroutine is done sending more values to the channel)
	// close(ch)
	// we need to check if channel is closed or not where we are recieving data from channel
	// data, ok := <-ch
}

// BUFFERED CHANNELS
// normal channels work like if there is sender there must be a reciever. channels themselves cannot store any
// data into them. This is where buffer channels comes in. They can store data inside them

// ch := make(chan int, 10)

// SELECT (SWTICH STATEMENT FOR CHANNELS)
// Sometimes we have a single goroutine listening to multiple channels and want to process data in the order it comes through each channel.
// A select statement is used to listen to multiple channels at the same time. It is similar to a switch statement but for channels.

// select {
// case i, ok := <- chInts:
//   fmt.Println(i)
// case s, ok := <- chStrings:
//   fmt.Println(s)
// }

// select statement can also have a default case which is executes immediately when no other channel has value to read. A default case stops the select statement from blocking.

// MARKING CHANNELS AS READ ONLY OR WRITE ONLY

// following marks channel as read only
// func main(){
// 	ch := make(chan int)
// 	readCh(ch)
//   }

//   func readCh(ch <-chan int) {
// 	// ch can only be read from
// 	// in this function
//   }

// following marks channel as write only
// func writeCh(ch chan<- int) {
// 	// ch can only be written to
// 	// in this function
//   }
