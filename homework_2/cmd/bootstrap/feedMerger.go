package bootstrap

import "code-cadets-2021/homework_2/internal/domain/services"

func NewFeedMerger(feeds ...services.Feed) *services.FeedMerger {
	return services.NewFeedMerger(feeds)
}
