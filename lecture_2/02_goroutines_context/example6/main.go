package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	// ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		t := time.Now()
		fmt.Println("separate goroutine started")
		<-ctx.Done()
		fmt.Println("separate goroutine done", time.Since(t))
	}()

	// cancel()
	fmt.Println("main goroutine sleeping")
	time.Sleep(time.Second * 5)
	fmt.Println("main goroutine wake up")
}
