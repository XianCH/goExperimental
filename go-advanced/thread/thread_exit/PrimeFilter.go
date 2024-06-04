package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 生成自然数序列的管道
func GenerateNatural(ctx context.Context, wg *sync.WaitGroup) chan int {
	defer wg.Done()
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
			}
		}
	}()
	return ch
}

// 管道过滤器：删除能被素数整除的数
func PrimFiltter(ctx context.Context, in <-chan int, prime int, wg *sync.WaitGroup) chan int {
	defer wg.Done()
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				select {
				case <-ctx.Done():
					return
				case out <- i:
				}
			}
		}
	}()
	return out
}

func main() {
	wg := sync.WaitGroup{}
	ctx, cancal := context.WithCancel(context.Background())
	wg.Add(1)
	ch := GenerateNatural(ctx, &wg)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		prime := <-ch //generate new prime
		fmt.Println("%v: %v\n", i+1, prime)
		ch = PrimFiltter(ctx, ch, prime, &wg)
	}
	time.Sleep(time.Second)
	wg.Wait()
	cancal()
}
