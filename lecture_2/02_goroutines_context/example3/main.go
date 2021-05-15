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
	x := 0
	mutex := &sync.Mutex{}

	go count(mutex, &x)
	go count(mutex, &x)

	// we could/should use wait group here
	time.Sleep(time.Second * 2)

	mutex.Lock()
	fmt.Println(x)
	mutex.Unlock()

	fmt.Println("done")
}

func count(mutex *sync.Mutex, a *int) {
	for i := 0; i < num; i++ {
		mutex.Lock()
		*a += 1
		mutex.Unlock()
	}
	fmt.Println("routine done")
}
