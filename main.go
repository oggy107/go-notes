package main

func main() {
	// errorExample()
	// interfaceExample()
	// structsExample()
	// mapsExample()

	// fmt.Println(mystrings.Reverse("Hello"))

	////////////// CONCURRENCY AND CHANNELS

	// using waitgroups for waiting for goroutines to finish (prefrerable for multiple goroutines)
	// var wg sync.WaitGroup
	// wg.Add(1)

	// go greetConcurrently("oggy", &wg)
	// fmt.Println("program doing other stuff....")
	// wg.Wait()

	// using channel to wait for goroutines to finish
	// ch := make(chan string)

	// go greetConcurrentlyChannel("jack", ch)

	// fmt.Println("program doing some more other stuff....")
	// recieving data from channel (this will wait (block) until data is recieved)
	// data := <-ch

	// fmt.Println("data from channel:", data)

	///////////////////////////////////////////////

	genericsExample()
}
