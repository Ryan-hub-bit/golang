package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// create a go routine to run newTask()
	go newTask()

	i := 0

	for {
		i++
		fmt.Printf("main Goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
