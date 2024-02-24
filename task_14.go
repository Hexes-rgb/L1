package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	variables := []interface{}{1, "str", true, make(chan struct{})} //create slice with vars
	writerChannel := make(chan interface{}, len(variables))         // create buff channel for writer

	wg.Add(2)
	go writer(&wg, writerChannel, variables) //run writer
	go typeChecker(&wg, writerChannel)       // run goroutine-typechecker
	// go typeChecker2(&wg, writerChannel)   //alternate variant

	wg.Wait()
}

func typeChecker(wg *sync.WaitGroup, writerChannel chan interface{}) {
	defer wg.Done()
	for variable := range writerChannel {
		switch variable.(type) { //use switch for check type
		case int:
			fmt.Printf("Variable %v, type: int\n", variable)
		case string:
			fmt.Printf("Variable %v, type: string\n", variable)
		case bool:
			fmt.Printf("Variable %v, type: boolean\n", variable)
		case chan struct{}:
			fmt.Printf("Variable %v, type: chan strunct{}\n", variable)
		}
	}
}

func writer(wg *sync.WaitGroup, writerChannel chan interface{}, variables []interface{}) {
	defer wg.Done()
	for _, elem := range variables { //write data to the channel
		writerChannel <- elem
	}
	close(writerChannel) //close channel after all data send
}

func typeChecker2(wg *sync.WaitGroup, writerChannel chan interface{}) {
	defer wg.Done()
	for variable := range writerChannel {
		fmt.Printf("Variable %v, type: %T\n", variable, variable) //use formatted output to get the type of a variable
	}
}
