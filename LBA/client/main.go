package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	numRequests := 10
	frontendURL := "http://localhost:8080"

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			resp, err := http.Get(frontendURL)
			if err != nil {
				fmt.Printf("Request %d: Error: %v\n", i, err)
				return
			}
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("Request %d: Response: %s", i, string(body))
			resp.Body.Close()
		}(i)
	}

	wg.Wait()
}
