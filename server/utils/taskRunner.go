package utils

import (
	"context"
	"sync"
)

type Task interface {
	Start(ctx context.Context)
}

func Run(tasks ...Task) {
	wg := sync.WaitGroup{}
	wg.Add(len(tasks))

	ctx, cancel := context.WithCancel(context.Background())

	for _, task := range tasks {
		go func(task Task) {
			defer wg.Done()
			defer cancel()
			defer CapturePanic(false, nil)

			task.Start(ctx)
		}(task)
	}
	wg.Wait()
	cancel()
}
