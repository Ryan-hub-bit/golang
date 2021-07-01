package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)

	go func() {
		defer fmt.Println("go routing end")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("subprocess is running element: ", i, "len(c) = ", len(c), "capacity = ", cap(c))
		}
	}()
	time.Sleep(2 * time.Second)
	for i := 0; i < 4; i++ {

		num := <-c
		fmt.Println("num = ", num)
	}
	fmt.Println("main routine end!")
}
