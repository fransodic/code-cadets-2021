package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	fmt.Print(<-ch)
	fmt.Print(<-ch)
	fmt.Println()

	ch <- 3
	ch <- 4
	ch <- 5
	fmt.Print(<-ch)
	fmt.Print(<-ch)
	fmt.Print(<-ch)
	fmt.Println()

	ch <- 6
	ch <- 7
	ch <- 8
	ch <- 9
	fmt.Print(<-ch)
	fmt.Print(<-ch)
	fmt.Print(<-ch)
	fmt.Print(<-ch)
	fmt.Println()
}
