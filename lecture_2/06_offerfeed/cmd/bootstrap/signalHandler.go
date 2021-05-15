package bootstrap

import "code-cadets-2021/lecture_2/06_offerfeed/internal/tasks"

func SignalHandler() *tasks.SignalHandler {
	return tasks.NewSignalHandler()
}
