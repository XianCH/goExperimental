package thread_model

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Hello() {
	done := make(chan int)
	go func() {
		fmt.Println("Hello, thread!")
		<-done
	}()
	done <- 1
	fmt.Println("MAIN")
}

func Hello2() {
	done := make(chan int, 1)

	go func() {
		fmt.Println("Hello, thread!")
		done <- 1
	}()
	<-done
	fmt.Println("MAIN")
}

func Hello3() {
	done := make(chan int, 10)

	for i := 0; i < cap(done); i++ {
		go func(int) {
			fmt.Printf("Hello, thread :%d!\n", i)
			done <- 1
		}(i)
	}

	for i := 0; i < cap(done); i++ {
		<-done
		fmt.Printf("MAIN: %d\n", i)
	}
}

func SyncWait() {

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Printf("Hello, thread !\n")
			wg.Done()
		}()
	}
	defer wg.Wait()
}

func publishr(facter int, out chan<- int) {
	for i := 0; i < facter; i++ {
		out <- i * facter
	}
}

func consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func TestFactor() {
	out := make(chan int, 64)
	go publishr(3, out)
	go publishr(5, out)

	go consumer(out)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	// fmt.Printf("quit (%v)\n", <-sig)
}

func ThreadQuit() {
	c := make(chan int, 10)
	for i := 1; i < cap(c); i++ {
		c <- i
	}
	time.Sleep(1 * time.Second)

	for {
		select {
		case v := <-c:
			fmt.Printf("Received value: %d\n", v)
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout: No value received within 2 seconds")
			return
		}
	}
}

func ThreadQuit2() {
	ch := make(chan int, 10)
	defer close(ch)
	go func() {
		for {
			select {
			case ch <- 1:
			case ch <- 0:
			}
		}
	}()

	for c := range ch {
		fmt.Println(c)
	}
}

func ThreadQuit3() {
	cancel := make(chan bool)
	go worker(cancel)
	time.Sleep(2 * time.Second)
	// cancel <- true
	close(cancel)
}

func worker(cancel chan bool) {
	for {
		select {
		default:
			fmt.Println("Worker is running")
		case <-cancel:
		}
	}
}

func worker2(cancel chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		default:
			fmt.Println("Worker is running")
		case <-cancel:
		}
	}
}

func ThreadQuit4() {
	cancel := make(chan bool)
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker2(cancel, wg)
	}

	time.Sleep(2 * time.Second)
	close(cancel)
	wg.Wait()
}

func ThreadQuit5() {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker5(wg, ctx)
	}
	time.Sleep(2 * time.Second)
	cancel()
	wg.Wait()
}

func worker5(wg *sync.WaitGroup, ctx context.Context) error {
	defer wg.Done()
	for {
		select {
		default:
			fmt.Println(1)
		case <-ctx.Done():
			fmt.Println("Worker is stopping")
			return ctx.Err()
		}
	}
}
