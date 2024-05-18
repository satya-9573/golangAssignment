package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Generator function generates numbers and sends them to the out channel.
func main() {
	wg := sync.WaitGroup{}
	startTime := time.Now()
	wg.Add(1)
	go func() {
		defer wg.Done()
		randomTimeGenarator()
	}()
	wg.Wait()
	duration := time.Until(startTime)
	timeTaken := duration * -1

	if timeTaken >= 3*time.Second {
		fmt.Println("Current time is at least 3 seconds greater than the given time.")
	} else {
		fmt.Println("Current time is less than 3 seconds greater than the given time.")
	}

}

func randomTimeGenarator() {
	randomInt := rand.Int63n(10)
	fmt.Printf("Program will sleep for %v seconds \n", randomInt)
	time.Sleep(time.Duration(randomInt) * time.Second)
}
