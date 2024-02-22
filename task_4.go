package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker(mainChannel <-chan string, id int) {
	for data := range mainChannel { //read data from main channel
		fmt.Printf("Worker #%d received: %v\n", id, data)
	}
	fmt.Printf("Worker #%d stopped: \n", id)
}

func main() {
	var numWorkers int

	fmt.Print("Enter the desired number of workers: ")
	fmt.Scan(&numWorkers)

	mainChannel := make(chan string)

	for i := 1; i <= numWorkers; i++ { //run workers
		go worker(mainChannel, i)
	}

	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Millisecond * 200)
			mainChannel <- fmt.Sprintf("msg %d", i) //send data to the channel
		}
	}()

	stop := make(chan os.Signal, 1)                    // define channel for stop signal
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM) // register stop signals

	<-stop //wait for stop signal

	close(mainChannel) //when a channel is closed, all reads operations from it immediately return to zero, allowing goroutines to exit
	fmt.Printf("\nExiting...\n")

}
