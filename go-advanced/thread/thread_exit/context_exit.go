package main

import (
	"context"
	"fmt"
	"sync"
)

func worker1(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello worker")
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

//
// func main() {
// 	ctx, cancal := context.WithTimeout(context.Background(), 10*time.Second)
// 	var wg sync.WaitGroup
// 	for i := 0; i <= 10; i++ {
// 		wg.Add(1)
// 		go worker1(ctx, &wg)
// 	}
//
// 	time.Sleep(time.Second)
// 	cancal()
// 	wg.Wait()
// }
