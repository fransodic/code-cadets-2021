package services

import (
	"code-cadets-2021/homework_2/internal/domain/models"
	"context"
	"sync"
)

type FeedMerger struct {
	updates chan models.Odd
	feeds   []Feed
}

func NewFeedMerger(feeds []Feed) *FeedMerger {
	return &FeedMerger{
		updates: make(chan models.Odd),
		feeds:   feeds,
	}
}

func (f *FeedMerger) Start(ctx context.Context) error {
	dest := f.GetUpdates()
	defer close(dest)

	wg := &sync.WaitGroup{}

	wg.Add(len(f.feeds))

	for _, feed := range f.feeds {
		go func(src chan models.Odd, dest chan models.Odd) {
			for v := range src {
				dest <- v
			}
			wg.Done()
		}(feed.GetUpdates(), dest)
	}

	wg.Wait()

	return nil
}

func (f *FeedMerger) GetUpdates() chan models.Odd {
	return f.updates
}

func (f *FeedMerger) String() string {
	return "feed merger"
}
