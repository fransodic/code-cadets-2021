package main

import (
	"log"

	"code-cadets-2021/lecture_3/03_project/calculator/cmd/bootstrap"
	"code-cadets-2021/lecture_3/03_project/calculator/cmd/config"
	"code-cadets-2021/lecture_3/03_project/calculator/internal/tasks"
)

func main() {
	log.Println("Bootstrap initiated")

	config.Load()

	rabbitMqChannel := bootstrap.RabbitMq()
	db := bootstrap.Sqlite()

	signalHandler := bootstrap.SignalHandler()
	engine := bootstrap.Engine(rabbitMqChannel, db)

	log.Println("Bootstrap finished. Engine is starting")

	tasks.RunTasks(signalHandler, engine)

	log.Println("Calculator service finished gracefully")
}
