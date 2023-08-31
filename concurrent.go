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
}
