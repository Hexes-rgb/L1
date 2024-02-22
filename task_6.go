package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// example1()
	// example2()
	// example3()
	// example4()
	example5()
}

func example1() {
	stop := make(chan bool) //using channel for stop goroutine

	go goroutineForExample1(stop)

	time.Sleep(time.Second * 2)
	stop <- true //send stop signal
}

func goroutineForExample1(stopChannel chan bool) {
	for {
		select {
		case <-stopChannel: //the goroutine will be stopped when it receives the signal
			fmt.Println("stopped")
			return
		default:
			fmt.Println("working..")
			time.Sleep(time.Second)
		}
	}
}

func example2() {
	var stop bool //using flag
	var wg sync.WaitGroup

	wg.Add(1)
	go goroutineForExample2(&stop, &wg)

	time.Sleep(time.Second * 2)

	stop = true //change flag state
	wg.Wait()   //wait for goroutine is finished
}

func goroutineForExample2(stop *bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for !*stop { //check flag state
		fmt.Println("working..")
		time.Sleep(time.Second * 1)
	}
	fmt.Println("stopped")
}

func example3() {
	var wg sync.WaitGroup

	stop := make(chan struct{}) //using channel closure to stop a goroutine

	wg.Add(1)
	go goroutineForExample3(stop, &wg)

	time.Sleep(time.Second * 3)

	close(stop)
	wg.Wait() //wait for finish
}

func goroutineForExample3(stop chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-stop: //after the channel is closed, the goroutine will receive a null value and exit
			fmt.Println("stopped")
			return
		default:
			fmt.Println("working..")
			time.Sleep(time.Second * 2)
		}
	}

}

func example4() {
	ctx, cancel := context.WithCancel(context.Background()) //use context WithCancel

	var wg sync.WaitGroup

	wg.Add(1)
	go goroutineForExample4(ctx, &wg)

	time.Sleep(time.Second * 3)
	cancel() //cancel context

	wg.Wait() //wait for finish
}

func goroutineForExample4(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done(): //get stop signal
			fmt.Println("stopped")
			return
		default:
			fmt.Println("working..")
			time.Sleep(time.Second * 2)
		}
	}
}

func example5() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2) //use context WithTimeout
	defer cancel()                                                          //defer context cancellation

	var wg sync.WaitGroup

	wg.Add(1)
	go goroutineForExample4(ctx, &wg)

	time.Sleep(time.Second * 3)

	wg.Wait() //wait for finish
}

func goroutineForExample5(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done(): //get stop signal
			fmt.Println("stopped")
			return
		default:
			fmt.Println("working..")
			time.Sleep(time.Second * 2)
		}
	}
}
