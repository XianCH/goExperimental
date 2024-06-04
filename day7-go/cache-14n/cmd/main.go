package main

import (
// "container/list"
// "fmt"
// "sync"
//
// cache_14n "github.com/x14n/goExperimental/day7_go/cache-14n"
)

//
// func main() {
// 	dl := cache_14n.New()
// 	var wg sync.WaitGroup
// 	list := list.New()
// 	list.Len()
// 	list.PushFront()
// 	list.MoveToFront()
//
// 	// Number of goroutines
// 	const numGoroutines = 10
// 	const numOperations = 100
//
// 	// Adding elements to the front
// 	for i := 0; i < numGoroutines; i++ {
// 		wg.Add(1)
// 		go func(id int) {
// 			defer wg.Done()
// 			for j := 0; j < numOperations; j++ {
// 				dl.AddFront(fmt.Sprintf("Front %d-%d", id, j))
// 			}
// 		}(i)
// 	}
//
// 	// Adding elements to the back
// 	for i := 0; i < numGoroutines; i++ {
// 		wg.Add(1)
// 		go func(id int) {
// 			defer wg.Done()
// 			for j := 0; j < numOperations; j++ {
// 				dl.AddBack(fmt.Sprintf("Back %d-%d", id, j))
// 			}
// 		}(i)
// 	}
//
// 	// Removing elements from the front
// 	for i := 0; i < numGoroutines; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for j := 0; j < numOperations; j++ {
// 				dl.RemoveFront()
// 			}
// 		}()
// 	}
//
// 	// Removing elements from the back
// 	for i := 0; i < numGoroutines; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for j := 0; j < numOperations; j++ {
// 				dl.RemoveBack()
// 			}
// 		}()
// 	}
//
// 	wg.Wait()
//
// 	// Check final state
// 	fmt.Printf("Final length of list: %d\n", dl.Len())
// 	if node, err := dl.GetFront(); err == nil {
// 		fmt.Printf("Front node value: %v\n", node.Value)
// 	} else {
// 		fmt.Println("Failed to get front node:", err)
// 	}
// 	if node, err := dl.GetBack(); err == nil {
// 		fmt.Printf("Back node value: %v\n", node.Value)
// 	} else {
// 		fmt.Println("Failed to get back node:", err)
// 	}
// }
