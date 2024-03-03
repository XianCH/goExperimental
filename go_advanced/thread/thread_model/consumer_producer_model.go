package thread_model

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Producer(v int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * v
	}
}

func Comsumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func Test_consumer_producer() {
	ch := make(chan int, 64)
	go Producer(3, ch)
	go Producer(5, ch)
	go Comsumer(ch)

	//ctrl c to quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}

//
// func test_main() {
// 	ch := make(chan int, 64)
// 	go Producer(3, ch)
// 	go Producer(5, ch)
// 	go Consumer(ch)
// 	sig := make(chan os.Signal, 1)
// 	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
// 	fmt.Printf("quit (%v)\n", <-sig)
// }
