package cache_14n

import (
	"testing"
)

func TestList(t *testing.T) {
	dl := New()

	for i := 0; i < 10; i++ {
		dl.PushNodeToFront(i)
	}
	result, _ := dl.GetFront()
	expect := 9

	if expect != result.Value {
		t.Errorf("GetFront result:%d expect:%d", result.Value, expect)
	}
}

//	func TestDoublyLinkedListConcurrency(t *testing.T) {
//		dl := New()
//		var wg sync.WaitGroup
//
//		// Number of goroutines
//		const numGoroutines = 10
//		const numOperations = 100
//
//		// Adding elements to the front
//		for i := 0; i < numGoroutines; i++ {
//			wg.Add(1)
//			go func(id int) {
//				defer wg.Done()
//				for j := 0; j < numOperations; j++ {
//					dl.AddFront(fmt.Sprintf("Front %d-%d", id, j))
//				}
//			}(i)
//		}
//
//		// Adding elements to the back
//		for i := 0; i < numGoroutines; i++ {
//			wg.Add(1)
//			go func(id int) {
//				defer wg.Done()
//				for j := 0; j < numOperations; j++ {
//					dl.AddBack(fmt.Sprintf("Back %d-%d", id, j))
//				}
//			}(i)
//		}
//
//		// Removing elements from the front
//		for i := 0; i < numGoroutines; i++ {
//			wg.Add(1)
//			go func() {
//				defer wg.Done()
//				for j := 0; j < numOperations; j++ {
//					dl.RemoveFront()
//				}
//			}()
//		}
//
//		// Removing elements from the back
//		for i := 0; i < numGoroutines; i++ {
//			wg.Add(1)
//			go func() {
//				defer wg.Done()
//				for j := 0; j < numOperations; j++ {
//					dl.RemoveBack()
//				}
//			}()
//		}
//
//		wg.Wait()
//
//		// Check final state
//		finalLen := dl.Len()
//		if finalLen != 0 {
//			t.Errorf("Expected final length to be 0, got %d", finalLen)
//		}
//
//		if node, err := dl.GetFront(); err == nil {
//			t.Errorf("Expected front node to be nil, got %v", node.Value)
//		}
//
//		if node, err := dl.GetBack(); err == nil {
//			t.Errorf("Expected back node to be nil, got %v", node.Value)
//		}
//	}
