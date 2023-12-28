package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cannal := context.WithCancel(context.Background())

	go func() {
		for {

			select {
			case <-ctx.Done():
				fmt.Println("get stop")
				return
			default:
				time.Sleep(1 * time.Second)
				fmt.Println("job 1")
			}
		}
	}()
	time.Sleep(5 * time.Second)
	cannal()
	fmt.Println("ALL JOB DOWN")
}
