package main

import "fmt"

func main() {
	buffer := 10

	ch := make(chan int, buffer)

	// - - - - - - - -
	// PLAIN ITERATION

	// writing to channel
	for i := 0; i < buffer; i++ {
		ch <- i
	}

	// classic for loop iteration
	for i := 0; i < buffer; i++ {
		fmt.Print(<-ch)
	}
	fmt.Println("done first")

	// - - - - - - - - -
	// RANGE OVER CHANNEL

	// writing to channel
	for i := 0; i < buffer; i++ {
		ch <- i
	}

	// close(ch)

	// range iteration
	for x := range ch {
		fmt.Print(x)
	}
	fmt.Println("done second")

	// - - - - - - - - - - - - - - -
	// zero values & close inspection
	ch2 := make(chan int, buffer)
	ch2 <- 1
	ch2 <- 2
	ch2 <- 3

	fmt.Println(<-ch2)

	v1, ok1 := <-ch2
	fmt.Println(v1, ok1)

	// close(ch2)

	v2, ok2 := <-ch2
	fmt.Println(v2, ok2)

	v3, ok3 := <-ch2
	fmt.Println(v3, ok3)

	v4, ok4 := <-ch2
	fmt.Println(v4, ok4)

	// close(ch2)
}
