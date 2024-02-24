package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Counter struct { //define structure for counter
	count int
	mu    sync.Mutex //use a mutex to avoid simultaneous value changes by several goroutines
}

func NewCounter(initialCount int) *Counter {
	return &Counter{
		count: initialCount,
	}
}

func (c *Counter) Run(seconds, workersCount int) {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(seconds)) //use context with timeout
	defer cancel()                                                                               //cancel context after timeout

	wg.Add(workersCount)
	for i := 0; i < workersCount; i++ { //create workers
		go worker(&wg, c, ctx)
	}

	wg.Wait() //wait all workers
}

func worker(wg *sync.WaitGroup, c *Counter, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done(): //wait stop signal
			return
		default:
			c.mu.Lock()   //lock access to counter data
			c.count++     //increment
			c.mu.Unlock() //unlock access
		}
	}

}

func main() {
	counter := NewCounter(0)
	counter.Run(2, 2)
	fmt.Println(counter.count)
}
