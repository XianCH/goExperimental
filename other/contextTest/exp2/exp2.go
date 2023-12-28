package main

import (
	"fmt"
	"time"
)

func main() {

	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("job down")
				return

			default:
				time.Sleep(1 * time.Second)
				fmt.Println("still working")
			}
		}
	}()

	time.Sleep(5 * time.Second)
	stop <- true
	fmt.Println("well down")
}
