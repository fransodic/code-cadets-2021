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

	for {
		select {
		case odd := <-upChan:
			updateCoefficient(odd, srcChan)
		case <-ctx.Done():
			return nil
		}
	}
}

func updateCoefficient(odd models.Odd, srcChan chan models.Odd) {
	odd.Coefficient = odd.Coefficient * 2
	srcChan <- odd
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
