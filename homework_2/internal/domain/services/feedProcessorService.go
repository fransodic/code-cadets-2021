package services

import (
	"code-cadets-2021/homework_2/internal/domain/models"
	"context"
)

type FeedProcessorService struct {
	feed  Feed
	queue Queue
}

func NewFeedProcessorService(orderedQueue Queue, feed Feed) *FeedProcessorService {
	return &FeedProcessorService{
		feed:  feed,
		queue: orderedQueue,
	}
}

func (f *FeedProcessorService) Start(ctx context.Context) error {
	// initially:
	// - get updates channel from feed interface
	// - get source channel from queue interface
	//
	// repeatedly:
	// - range over updates channel
	// - multiply each odd with 2
	// - send it to source channel
	//
	// finally:
	// - when updates channel is closed, exit
	// - when exiting, close source channel

	upChan := f.feed.GetUpdates()
	srcChan := f.queue.GetSource()
	defer close(srcChan)

	for odd := range upChan {
		//fmt.Printf("stari koef %g\n",  odd.Coefficient)
		odd.Coefficient = odd.Coefficient * 2.0
		//fmt.Printf("novi koef %g\n",  odd.Coefficient)
		srcChan <- odd
	}
	return nil
}

func (f *FeedProcessorService) String() string {
	return "feed processor service"
}

// define feed interface here
type Feed interface {
	GetUpdates() chan models.Odd
}

// define queue interface here
type Queue interface {
	GetSource() chan models.Odd
}
