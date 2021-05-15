package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	// close(ch1)

	ch2 := make(chan string, 3)
	ch2 <- "a"
	ch2 <- "b"
	ch2 <- "c"
	// close(ch2)

	for {
		select {
		case v1, ok := <-ch1:
			if !ok {
				fmt.Println("ch1 returned")
				return
			}
			fmt.Println(v1, ok)

		case v2, ok := <-ch2:
			if !ok {
				fmt.Println("ch2 returned")
				return
			}
			fmt.Println(v2, ok)
		}
	}
}
