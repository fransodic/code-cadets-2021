package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	size    = 100
	threads = 5

	axilisFeedURL = "http://18.193.121.232/axilis-feed"
)

func main() {
	httpClient := &http.Client{}

	c1 := initChannel()
	httpRequests1(c1, httpClient)

	c2 := initChannel()
	httpRequests2(c2, httpClient)
}

func initChannel() chan int {
	c := make(chan int, size)
	defer close(c)

	for i := 0; i < size; i++ {
		c <- i
	}

	return c
}

func httpRequests1(c chan int, client *http.Client) {
	t := time.Now()

	for range c {
		client.Get(axilisFeedURL)
	}

	fmt.Println("httpRequests1 duration", time.Since(t))
}

func httpRequests2(c chan int, client *http.Client) {
	t := time.Now()

	wg := &sync.WaitGroup{}
	wg.Add(threads)

	for i := 0; i < threads; i++ {
		go func() {
			defer wg.Done()

			for range c {
				client.Get(axilisFeedURL)
			}
		}()
	}

	wg.Wait()

	fmt.Println("httpRequests2 duration", time.Since(t))
}
