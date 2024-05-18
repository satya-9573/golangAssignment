package main

import (
	"fmt"
	"sync"
)

// Generator function generates numbers and sends them to the out channel.
func generator(out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ { // Generate numbers from 1 to 10
		out <- i
	}
	close(out) // Close the channel when done
}

// Squarer function reads numbers from the in channel, squares them, and sends the results to the out channel.
func squarer(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range in {
		out <- num * num
	}
	close(out) // Close the channel when done
}

// Printer function reads squared numbers from the in channel and prints them.
func printNums(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for squared := range in {
		fmt.Println(squared)
	}
}

func main() {
	// Create channels
	nums := make(chan int)
	squaredNums := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(3)

	// Start the goroutines
	go generator(nums, wg)
	go squarer(nums, squaredNums, wg)
	go printNums(squaredNums, wg) // We can run this directly since it only needs to print

	// Wait for a moment to ensure all prints complete (not needed in more complex examples)
	wg.Wait()
}
