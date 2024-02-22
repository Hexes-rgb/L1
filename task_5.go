package main

import (
	"fmt"
	"time"
)

func main() {
	var N int

	fmt.Print("Enter N: ")
	fmt.Scan(&N)

	mainChannel := make(chan int) // create channel for data

	go func() { // run goroutine-publisher
		for i := 0; ; i++ {
			time.Sleep(time.Millisecond * 200)
			mainChannel <- i
		}
	}()

	go func() { //run goroutine-subscriber
		for data := range mainChannel {
			fmt.Printf("Received data: %v\n", data)
		}
	}()

	time.Sleep(time.Second * time.Duration(N)) //sleep...
	fmt.Println("Exiting...")
}
