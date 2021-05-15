package main

import (
	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/models"
	"code-cadets-2021/lecture_2/05_offerfeed/internal/infrastructure/queue"
	"fmt"
	"log"
	"time"
)

func main() {

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
		log.Fatal(err, "error starting the ordered queue")
	}

	fmt.Println("program finished gracefully")
}
