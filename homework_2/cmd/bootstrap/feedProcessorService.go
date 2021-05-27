package bootstrap

import "code-cadets-2021/homework_2/internal/domain/services"

func NewFeedProcessorService(queue services.Queue, feed services.Feed) *services.FeedProcessorService {
	return services.NewFeedProcessorService(queue, feed)
}
