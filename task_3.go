package main

import (
	"fmt"
	"sync"
)

func main() {
	// someFunc1()
	someFunc2()
}

func someFunc1() {
	nums := []int{2, 4, 6, 8, 10}
	var squaresSum int
	squaresChannel := make(chan int, len(nums)) //creation of the buffered channel
	var wg sync.WaitGroup                       //use WaitGroup to wait for multiple goroutines to finish

	wg.Add(len(nums))
	for _, num := range nums {
		go func(a int) { //run goroutines
			defer wg.Done()
			squaresChannel <- a * a // send squared to the channel
		}(num)
	}

	go func() { //run goroutine to wait all goroutines are finished
		wg.Wait()
		close(squaresChannel) //close channel
	}()

	for square := range squaresChannel { //sum calculation
		squaresSum += square
	}
	fmt.Println(squaresSum)
}

func someFunc2() {
	nums := []int{2, 4, 6, 8, 10}
	var mu sync.Mutex //use mutex to be sure only one goroutine can access a variable at a time
	var squaresSum int
	var wg sync.WaitGroup //use WaitGroup to wait for multiple goroutines to finish

	wg.Add(len(nums))
	for _, num := range nums { //run goroutines
		go func(a int) {
			defer wg.Done()
			mu.Lock() //define a block of code to be executed in mutual exclusion
			squaresSum += a * a
			defer mu.Unlock()
		}(num)
	}

	wg.Wait() //wait all goroutines are finished
	fmt.Println(squaresSum)
}
