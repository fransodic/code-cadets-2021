package main

import (
	"code-cadets-2021/lecture_2/05_offerfeed/cmd/bootstrap"
	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/models"
	"code-cadets-2021/lecture_2/05_offerfeed/internal/infrastructure/queue"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"time"
)

func main() {

	//err := testFirstTask()
	//if err != nil {
	//	log.Fatal(err, "error while executing the first task")
	//}

	err := testSecondTask()
	if err != nil {
		log.Fatal(err, "error while executing the second task")
	}

	fmt.Println("program finished gracefully")
}

func testSecondTask() error {
	axilisOfferFeed := bootstrap.NewAxilisOfferFeed()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	go axilisOfferFeed.Start(ctx)

	for odd := range axilisOfferFeed.GetUpdates() {
		fmt.Print("Jedan od rezultata je ")
		fmt.Println(odd)
	}

	return nil
}

func testFirstTask() error {
	oq := queue.NewOrderedQueue()
	src := oq.GetSource()

	src <- models.Odd{
		Id:          "ID1234",
		Name:        "Ime 123",
		Match:       "Random Match",
		Coefficient: 1.10,
		Timestamp:   time.Time{},
	}

	src <- models.Odd{
		Id:          "ID4321",
		Name:        "Ime 456",
		Match:       "New Match",
		Coefficient: 2.20,
		Timestamp:   time.Time{},
	}

	close(src)

	err := oq.Start(nil)
	if err != nil {
		return errors.WithMessage(err, "error starting the ordered queue")
	}

	return nil
}
