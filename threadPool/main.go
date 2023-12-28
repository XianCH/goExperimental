package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	conVar := sync.NewCond(&mu)
	for i := 1; i <= 3; i++ {
		go worker(i, &mu, conVar)
	}
	time.Sleep(time.Second)
	fmt.Println("Setting signal to Workers...")
	mu.Lock()
	conVar.Broadcast()
	mu.Unlock()

	time.Sleep(time.Second)
}

func worker(id int, mu *sync.Mutex, convar *sync.Cond) {
	mu.Lock()
	defer mu.Unlock()

	fmt.Printf("Worker %d is waiting...\n", id)
	convar.Wait()
	fmt.Printf("Workder %d is received signal!\n", id)
}
