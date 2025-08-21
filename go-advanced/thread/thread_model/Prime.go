package thread_model

import (
	"context"
	"fmt"
	"sync"
)

func GenerateNatural(ctx context.Context, wg *sync.WaitGroup) chan int {
	ch := make(chan int)
	defer wg.Done()
	go func() {
		for i := 2; ; i++ {
			select {
			case ch <- i:
			case <-ctx.Done():
				return
			}
		}
	}()
	return ch
}

func FilterPrime(ctx context.Context, in chan int, prime int, wg *sync.WaitGroup) chan int {
	defer wg.Done()
	out := make(chan int)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case i := <-in:
				if i%prime != 0 {
					select {
					case out <- i:
					case <-ctx.Done():
						return
					}
				}
			}

		}
	}()
	return out
}

func PrimeMain() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	ch := GenerateNatural(ctx, wg)
	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Printf("%v:%v \n", i+1, prime)
		wg.Add(1)
		ch = FilterPrime(ctx, ch, prime, wg)
	}
	cancel()
	wg.Wait()
}
