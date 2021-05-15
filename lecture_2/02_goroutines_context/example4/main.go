package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	num = 10000
)

func main() {
	mutex := &sync.Mutex{}
	isDone := false

	data := make([]int, 0, num)

	go fillSlice(mutex, &isDone, &data)
	go sumSlice(mutex, &isDone, &data)

	// we should use wait group here
	time.Sleep(time.Second * 2)
	fmt.Println("program done")
}

func fillSlice(mutex *sync.Mutex, isDone *bool, a *[]int) {
	for i := 0; i < num; i++ {
		mutex.Lock()
		*a = append(*a, 1)
		mutex.Unlock()
	}

	mutex.Lock()
	*isDone = true
	mutex.Unlock()

	fmt.Println("fill slice done")
}

func sumSlice(mutex *sync.Mutex, isDone *bool, a *[]int) {
	counter := 0

	for {
		mutex.Lock()
		if *isDone && len(*a) == 0 {
			mutex.Unlock()
			break

		} else if len(*a) == 0 {
			mutex.Unlock()
			continue
		}

		counter += (*a)[0]
		*a = (*a)[1:]
		mutex.Unlock()
	}

	fmt.Println("sum slice done", counter)
}
