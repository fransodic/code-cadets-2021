package bootstrap

import "code-cadets-2021/lecture_2/06_offerfeed/internal/domain/services"

func FeedProcessingService(feed services.Feed, queue services.Queue) *services.FeedProcessorService {
	return services.NewFeedProcessorService(feed, queue)
}
