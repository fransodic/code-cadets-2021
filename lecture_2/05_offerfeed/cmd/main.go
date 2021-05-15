package main

import (
	"code-cadets-2021/lecture_2/05_offerfeed/cmd/bootstrap"
	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/services"
	"code-cadets-2021/lecture_2/05_offerfeed/internal/tasks"
	"fmt"
)

func main() {

	signalHandler := bootstrap.NewSignalHandler()

	feed := bootstrap.NewAxilisOfferFeed()
	queue := bootstrap.NewOrderedQueue()
	processingService := services.NewFeedProcessorService(queue, feed)

	tasks.RunTasks(signalHandler, feed, queue, processingService)

	fmt.Println("program finished gracefully")
}

//func testThirdTask() error {
//	wg := &sync.WaitGroup{}
//
//	wg.Add(2)
//
//	orderedQueue := bootstrap.NewOrderedQueue()
//	axilisOfferFeed := bootstrap.NewAxilisOfferFeed()
//
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 7)
//	defer cancel()
//
//	go orderedQueue.Start(ctx, wg)
//	go axilisOfferFeed.Start(ctx, wg)
//
//	processorService := services.NewFeedProcessorService(orderedQueue, axilisOfferFeed)
//
//	err := processorService.Start(ctx)
//	if err != nil {
//		return err
//	}
//
//	wg.Wait()
//
//	return nil
//}

//func testSecondTask() error {
//	axilisOfferFeed := bootstrap.NewAxilisOfferFeed()
//
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
//	defer cancel()
//
//	axilisOfferFeed.Start(ctx)
//
//	for odd := range axilisOfferFeed.GetUpdates() {
//		fmt.Print("Jedan od rezultata je ")
//		fmt.Println(odd)
//	}
//
//	return nil
//}

//func testFirstTask() error {
//	oq := queue.NewOrderedQueue()
//	src := oq.GetSource()
//
//	src <- models.Odd{
//		Id:          "ID1234",
//		Name:        "Ime 123",
//		Match:       "Random Match",
//		Coefficient: 1.10,
//		Timestamp:   time.Time{},
//	}
//
//	src <- models.Odd{
//		Id:          "ID4321",
//		Name:        "Ime 456",
//		Match:       "New Match",
//		Coefficient: 2.20,
//		Timestamp:   time.Time{},
//	}
//
//	close(src)
//
//	err := oq.Start(nil)
//	if err != nil {
//		return errors.WithMessage(err, "error starting the ordered queue")
//	}
//
//	return nil
//}
