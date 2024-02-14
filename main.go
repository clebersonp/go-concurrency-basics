package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	// communicate when chan has changed
	// this arrow in this way means that we are passing data to the channel, so we are changing it to communicate outside
	doneChan <- true

	// to for loop works fine we must close the channel somehow
	// close it only make sense if we now where in the code will perform more time than other places
	// because, when it's closed we can't use it again
	close(doneChan)
}

func main() {
	// channel is used to communicate between goroutines, we can pass or receive data from channels
	// passing channel to function to communicate from there
	done := make(chan bool)

	// keyword go means that go will run code as goroutines (parallel)
	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done)

	// this arrow in this way means we are receiving data from channel when its has some changed
	// receiving all channel values that was communicated before
	//<-done
	//<-done
	//<-done
	//<-done

	// using for loop to waiting for receive values from channel
	for range done {
	}
}
