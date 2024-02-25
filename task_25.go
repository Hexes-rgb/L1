package main

import (
	"fmt"
	"time"
)

func main() {
	sleep(time.Second * 3)
	fmt.Println("Exiting..")
}

// Sleep function, based of time.After
func sleep(d time.Duration) {
	<-time.After(d)
}
