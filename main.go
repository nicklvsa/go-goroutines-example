package main

import (
	"fmt"
	"sync"
)

func main() {

	// create the wait group and response channel
	var waiter sync.WaitGroup
	response := make(chan string)

	// add some workers
	for i := 0; i <= 5; i++ {
		waiter.Add(1)
		go example(fmt.Sprintf("Executing %d", i), response, &waiter)
	}

	// add some more workers
	for i := 0; i <= 10; i++ {
		waiter.Add(1)
		go example(fmt.Sprintf("Executing 2nd %d", i), response, &waiter)
	}

	// wait for data to appear in the response channel and display them
	go func() {
		for {
			select {
			case resp := <-response:
				fmt.Println(resp)
			}
		}
	}()

	// wait for the waiter to finish
	waiter.Wait()
	close(response)
}

// worker function to send the data to its channel and update the waiter
func example(data string, channel chan<- string, waiter *sync.WaitGroup) {
	defer waiter.Done()
	channel <- data
}
