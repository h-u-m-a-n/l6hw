package l6hw

import (
	"context"
	"fmt"
	"sync"
)
func Execute(tasks []func(ctx context.Context) error, E int) error {
	var (
		wg              sync.WaitGroup
		mu              sync.Mutex
		err             error
		errCount        = 0
		ctx, cancelFunc = context.WithCancel(context.Background())
	)
	defer cancelFunc()
	for _, task := range tasks {
		wg.Add(1)
		go func(task func(ctx context.Context) error) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			if locErr := task(ctx); locErr != nil {
				errCount++
				err = locErr
			}
			fmt.Println("Number of errors: ", errCount)
			if errCount >= E {
				cancelFunc()
			}
		}(task)
	}
	wg.Wait()
	if errCount >= E {
		return err
	}
	return nil

}
