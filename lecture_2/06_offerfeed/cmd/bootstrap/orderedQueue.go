package bootstrap

import "code-cadets-2021/lecture_2/06_offerfeed/internal/infrastructure/queue"

func OrderedQueue() *queue.OrderedQueue {
	return queue.NewOrderedQueue()
}
