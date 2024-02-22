package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	// someFunc3()
	// someFunc2()
	someFunc1()
}

func someFunc3() {
	nums := []int{2, 4, 6, 8, 10}
	results := make(chan int, len(nums)) // Create a buffered channel

	// Run goroutines in a loop and pass them values from the slice
	for _, num := range nums {
		go func(x int) {
			square := x * x
			results <- square // send data to the results channel
		}(num)
	}

	for i := 0; i < len(nums); i++ {
		square := <-results // get and print data from the results channel
		fmt.Printf("Number:%.f Square:%d\n", math.Sqrt(float64(square)), square)
	}
}

func someFunc2() {
	nums := []int{2, 4, 6, 8, 10}
	results := make(chan int, len(nums)) //create buff channel

	var wg sync.WaitGroup //use WaitGroup to wait for multiple goroutines to finish
	wg.Add(len(nums))     //increase goroutines counter

	for _, num := range nums { //run routines
		go func(x int) {
			defer wg.Done() //decrease routines counter after routine finished
			square := x * x
			results <- square //send to channel
		}(num)
	}

	go func() { //run goroutine to wait all goroutines are finished
		wg.Wait()
		close(results) //close channel to avoid deadlock
	}()

	for square := range results {
		fmt.Printf("Number:%.f Square:%d\n", math.Sqrt(float64(square)), square)
	}
}

func someFunc1() {
	nums := []int{2, 4, 6, 8, 10}
	done := make(chan bool) //create a channel to signal the end of the goroutine

	for _, num := range nums { //run goroutines
		go func(x int) {
			square := x * x
			fmt.Printf("Number:%.f Square:%d\n", math.Sqrt(float64(square)), square)
			done <- true //send finish signal
		}(num)
	}

	for i := 0; i < len(nums); i++ {
		<-done // waiting for all goroutines to finish
	}

}
