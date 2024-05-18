package main

import (
	"fmt"
	"sync"
)

type Increment struct {
	Counter int
}

func main() {
	var (
		increment Increment
		wg        sync.WaitGroup
		mutex     sync.Mutex
	)

	const goroutines = 100
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {

		go func() {
			defer wg.Done()

			// Lock the mutex to synchronize access to the shared variable
			mutex.Lock()
			// Increment the counter
			increment.Counter++
			// Unlock the mutex
			mutex.Unlock()
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the final value of the counter
	fmt.Println("Final Counter Value:", increment.Counter)
}
