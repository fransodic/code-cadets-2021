package tasks

import (
	"context"
	"log"
	"sync"
)

func RunTasks(tasks ...Task) {
	wg := &sync.WaitGroup{}
	wg.Add(len(tasks))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i, task := range tasks {
		go func(i int, task Task) {
			defer wg.Done()
			defer cancel()

			err := task.Start(ctx)
			log.Printf(`"%s" finished with "%v" error`, task, err)
		}(i, task)
	}

	log.Print("all tasks running, waiting")
	log.Print("- - - - - - - - - - - - - -")
	wg.Wait()
	log.Print("all tasks finished")
}

type Task interface {
	Start(ctx context.Context) error
}
