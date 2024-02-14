package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	// communicate when chan has changed
	// this arrow in this way means that we are passing data to the channel, so we are changing it to communicate outside
	doneChan <- true
}

func main() {
	// keyword go means that go will run code as goroutines (parallel)
	//go greet("Nice to meet you!")
	//go greet("How are you?")

	// channel is used to communicate between goroutines, we can pass or receive data from channels
	done := make(chan bool)
	// passing channel to function to communicate from there
	go slowGreet("How ... are ... you ...?", done)

	// this arrow in this way means we are receiving data from channel when its has some changed
	receivedChannelData := <-done
	fmt.Println(receivedChannelData)

	//go greet("I hope you're liking the course!")
}
