/*
Assignment 2: Producer-Consumer with Channel Problem: Implement the producer-consumer problem using goroutines and channels.
The producer should generate numbers from 1 to 100 and send them to a channel, and the consumer should print those numbers.
*/

package main

import (
	"fmt"
	"sync"
)

// producer generates numbers from 1 to 100 and sends them to the channel.
// It takes a channel of integers as input.
func producer(ch chan int) {
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch) // Close the channel after sending all numbers
}

// consumer reads numbers from the channel and prints them.
// It takes a channel of integers and a channel of boolean values as input.
func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println(num)
	}
}

func main() {
	ch := make(chan int) // Create a channel for communication between producer and consumer

	wg := sync.WaitGroup{} // Create a wait group for untill all the consumers are printed
	wg.Add(1)
	go producer(ch) // Start the producer goroutine

	go consumer(ch, &wg) // Start the consumer goroutine

	wg.Wait() // Wait for the consumer to finish
}
