package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, cancal chan bool) {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello worker!")
		case <-cancal:
			return
		}
	}
}

//
// func main() {
// 	cancal := make(chan bool)
// 	var wg sync.WaitGroup
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go worker(&wg, cancal)
// 	}
// 	time.Sleep(time.Second)
// 	close(cancal)
// 	wg.Wait()
// }
