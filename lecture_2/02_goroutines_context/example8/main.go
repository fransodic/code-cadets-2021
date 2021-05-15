package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	go sayHelloEachSecond(ctx, "axilis")
	go sayHelloEachSecond(ctx, "cadets")

	// cancel()
	fmt.Println("main goroutine sleeping")
	time.Sleep(time.Second * 5)
	fmt.Println("main goroutine wake up")
}

func sayHelloEachSecond(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("finished", name)
			return
		case <-time.After(time.Second):
			fmt.Println("hello", name)
		}
	}
}
