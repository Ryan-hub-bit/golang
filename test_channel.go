package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("go routing end")

		fmt.Println("go routine running")
		c <- 555
	}()
	num := <-c
	fmt.Printf("num: %d", num)
	fmt.Println("main routine end!")
}
