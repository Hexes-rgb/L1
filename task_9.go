package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numsChannel := make(chan int, len(numbers))
	resultsChannel := make(chan int, len(numbers)) //define channels
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, num := range numbers {
			numsChannel <- num //send data to the channel
		}
		close(numsChannel) //close channel when all data send
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range numsChannel {
			resultsChannel <- num * 2 //send data to the channel
		}
		close(resultsChannel) //close channel when all data send
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for res := range resultsChannel { //read results
			fmt.Println(res)
		}
	}()

	wg.Wait() //wait all goroutines
}
