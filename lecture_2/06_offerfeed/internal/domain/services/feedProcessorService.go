package services

import (
	"context"
	"log"

	"code-cadets-2021/lecture_2/06_offerfeed/internal/domain/models"
)

type FeedProcessorService struct {
	feed  Feed
	queue Queue
}

func NewFeedProcessorService(
	feed Feed,
	queue Queue,
) *FeedProcessorService {
	return &FeedProcessorService{
		feed:  feed,
		queue: queue,
	}
}

func (f *FeedProcessorService) Start(ctx context.Context) error {
	updates := f.feed.GetUpdates()
	source := f.queue.GetSource()

	defer close(source)
	defer log.Printf("shutting down %s", f)

	for update := range updates {
		update.Coefficient *= 2
		source <- update
	}

	return nil
}

func (f *FeedProcessorService) String() string {
	return "feed processor service"
}

type Feed interface {
	GetUpdates() chan models.Odd
}

type Queue interface {
	GetSource() chan models.Odd
}
