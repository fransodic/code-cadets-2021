package main

import (
	"fmt"
	"time"
)

const (
	num = 10000
)

func main() {
	data := make(chan int)

	go fillChan(data)
	// main goroutine can perform sum function
	// in that case, we do not need sleep call
	// main goroutine can execute relevant work!
	go sumChan(data)

	time.Sleep(time.Second * 2)
	fmt.Println("program done")
}

func fillChan(data chan int) {
	defer close(data)

	for i := 0; i < num; i++ {
		data <- 1
	}

	fmt.Println("fill chan done")
}

func sumChan(data chan int) {
	counter := 0

	for x := range data {
		counter += x
	}

	fmt.Println("sum chan done", counter)
}
