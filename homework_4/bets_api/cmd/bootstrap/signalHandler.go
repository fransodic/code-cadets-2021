package bootstrap

import "code-cadets-2021/homework_4/bets_api/internal/tasks"

// SignalHandler bootstraps the signal handler.
func SignalHandler() *tasks.SignalHandler {
	return tasks.NewSignalHandler()
}
