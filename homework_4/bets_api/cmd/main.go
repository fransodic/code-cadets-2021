package main

import (
	"log"

	"code-cadets-2021/homework_4/bets_api/cmd/bootstrap"
	"code-cadets-2021/homework_4/bets_api/cmd/config"
	"code-cadets-2021/homework_4/bets_api/internal/tasks"
)

func main() {
	log.Println("Bootstrap initiated")

	config.Load()

	db := bootstrap.Sqlite()
	signalHandler := bootstrap.SignalHandler()
	api := bootstrap.Api(db)

	log.Println("Bootstrap finished. Event API is starting")

	tasks.RunTasks(signalHandler, api)

	log.Println("Event API finished gracefully")
}
