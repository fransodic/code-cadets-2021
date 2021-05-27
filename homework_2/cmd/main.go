package main

import (
	"code-cadets-2021/homework_2/cmd/bootstrap"
	"code-cadets-2021/homework_2/internal/tasks"
	"fmt"
)

func main() {

	signalHandler := bootstrap.NewSignalHandler()

	client := bootstrap.NewHttpClient()

	feedJSON := bootstrap.NewAxilisOfferFeedJSON(*client)
	feedCSV := bootstrap.NewAxilisOfferFeedCSV(*client)
	feed := bootstrap.NewFeedMerger(feedJSON, feedCSV)

	queue := bootstrap.NewOrderedQueue()
	processingService := bootstrap.NewFeedProcessorService(queue, feed)

	tasks.RunTasks(signalHandler, feedJSON, feedCSV, feed, queue, processingService)

	fmt.Println("program finished gracefully")
}
