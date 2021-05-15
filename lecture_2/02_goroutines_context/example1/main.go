package main

import (
	"fmt"
	"sync"
)

const (
	size = 10
)

func main() {
	wg := &sync.WaitGroup{}
	// we need to initialise wait group with number of goroutines
	wg.Add(size)

	for i := 0; i < size; i++ {
		go func(a int) {
			// this subtracts "1" from wait group
			defer wg.Done()

			fmt.Printf("hi, im %d. goroutine\n", a)
		}(i)
	}

	// main goroutine is waiting here, while other routines are saying hello to you
	wg.Wait()
	fmt.Println("main program finished")
}
