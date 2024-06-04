package atomic

import (
	"sync"
	"sync/atomic"
)

var total uint64

func workder(wg *sync.WaitGroup) {
	defer wg.Done()

	var i uint64
	for i = 0; i <= 100; i++ {
		atomic.AddUint64(&total, i)
	}
}

func TestWorker() {
	var wg sync.WaitGroup
	wg.Add(2)
	go workder(&wg)
	go workder(&wg)

	wg.Wait()
}
