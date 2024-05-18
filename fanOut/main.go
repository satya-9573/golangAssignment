/*
	The program runs on fan-out concurrency pattern.
	It reads data via multiple Go routines and prints them.
*/

package main

import (
	"fmt"
	"sync"
)

func InsertDataInChannel(nums []int, channel chan int) {

	go func() {
		//iterate the nums data and sends it to channel
		for _, val := range nums {
			channel <- val
		}
		close(channel)
	}()
}

func main() {
	dataToInsertInChannel1 := []int{100, 99, 98, 97, 95}
	dataToInsertInChannel2 := []int{10, 20, 30, 40, 50}

	channel1 := make(chan int)
	channel2 := make(chan int)
	var wg sync.WaitGroup

	//it receives a "receive-only" directional channel
	go InsertDataInChannel(dataToInsertInChannel1, channel1)
	go InsertDataInChannel(dataToInsertInChannel2, channel2)
	wg.Add(2)

	//we will loop through both the channels till all data is sent and marked as close
	go func() {
		for val := range channel1 {
			fmt.Printf("Data from Channel 1: %v\n", val)
		}
		wg.Done()
	}()

	go func() {
		for val := range channel2 {
			fmt.Printf("Data from Channel 2: %v\n", val)
		}
		wg.Done()
	}()

	wg.Wait() //will wait till the above goroutines are marked as done
}
