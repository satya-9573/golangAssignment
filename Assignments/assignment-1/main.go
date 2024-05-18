/*
Assignment 1: Goroutine with Channel Problem: Write a Go program that calculates the
sum of numbers from 1 to N concurrently using goroutines and channels.
The program should take the value of N as input from the user.
*/
package main

import (
	"fmt"
)

func main() {
	var num int

	ch := make(chan int)
	// Scanner for user to input the number
	fmt.Print("Please give the input number: ")
	fmt.Scan(&num)
	// go routine for calculating the sum
	go getSumUntilNum(int(num), ch)
	sumOfNum := <-ch
	//Printing the sum of numbers
	fmt.Printf("The sum of numbers from 1 to %d is: %d\n", num, sumOfNum)

}

func getSumUntilNum(N int, ch chan int) {
	// initialising sum to 0
	sum := 0
	for i := 1; i <= N; i++ {
		sum += i
	}
	ch <- sum
	close(ch)
}
