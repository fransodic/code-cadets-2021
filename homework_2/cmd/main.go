package main

import (
	"code-cadets-2021/homework_2/cmd/bootstrap"
	"code-cadets-2021/homework_2/internal/tasks"
	"fmt"
)

func main() {

	signalHandler := bootstrap.NewSignalHandler()

	feed := bootstrap.NewAxilisOfferFeed()
	feedCSV := bootstrap.NewAxilisOfferFeedCSV()
	queue := bootstrap.NewOrderedQueue()
	processingService := bootstrap.NewFeedProcessorService(queue, feed, feedCSV)

	tasks.RunTasks(signalHandler, feedCSV, feed, queue, processingService)

	fmt.Println("program finished gracefully")
}
