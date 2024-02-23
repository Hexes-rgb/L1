package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
)

type MyMap struct { //definition structure for map
	data          map[uuid.UUID]int
	writerChannel chan KeyValue //channel for data
}

func (m *MyMap) RunReader(wg *sync.WaitGroup, ctx context.Context) { //goroutine-reader
	defer wg.Done()
	defer close(m.writerChannel)
	for {
		select {
		case <-ctx.Done(): //wait stop signal
			fmt.Println("Reader goroutine stopped")
			return
		case data := <-m.writerChannel: //store data in the map
			fmt.Printf("Reader received key: %v value: %d\n", data.key, data.value)
			m.data[data.key] = data.value
		}
	}
}

func (m *MyMap) Set(key uuid.UUID, value int) { //method to send KeyValue data to data channel
	m.writerChannel <- KeyValue{key, value}
}

func NewMyMap() *MyMap {
	return &MyMap{
		data:          make(map[uuid.UUID]int),
		writerChannel: make(chan KeyValue, 5),
	}
}

type KeyValue struct { //structure for data
	key   uuid.UUID
	value int
}

func main() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) //use context to manage the life cycle of goroutines
	defer cancel()

	wg.Add(1)
	myMap := NewMyMap()
	go myMap.RunReader(&wg, ctx) //run reader

	for i := 1; i <= 5; i++ { //run writers
		wg.Add(1)
		go writer(&wg, ctx, myMap, i)
	}

	wg.Wait()
	time.Sleep(time.Second)
	fmt.Println("Exiting...")
}

func writer(wg *sync.WaitGroup, ctx context.Context, m *MyMap, id int) { //goroutine-writer
	defer wg.Done()
	for {
		select {
		case <-ctx.Done(): //wait stop signal
			fmt.Printf("Writer #%v stopped\n", id)
			return
		default: //generate new data, send it to writer cahnnel, sleep
			key := uuid.New()
			value := rand.Int()
			m.Set(key, value)
			fmt.Printf("Writer #%d send key: %v value: %d\n", id, key, value)
			time.Sleep(time.Millisecond * 500)
		}
	}
}
