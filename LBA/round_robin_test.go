package lba

import (
	"sync"
	"testing"
)

func TestRoundRobinBalancer(t *testing.T) {
	servers := []string{"http://localhost:8081", "http://localhost:8082", "http://localhost:8083"}
	balancer := NewRoundRobinBalancer(servers)

	expectedOrder := []string{
		"http://localhost:8081",
		"http://localhost:8082",
		"http://localhost:8083",
		"http://localhost:8081",
		"http://localhost:8082",
		"http://localhost:8083",
	}

	for i, expected := range expectedOrder {
		actual := balancer.NextServer()
		if actual != expected {
			t.Errorf("Iteration %d: expected %s, got %s", i, expected, actual)
		}
	}
}

func TestRoundRobinBalancerConcurrency(t *testing.T) {
	servers := []string{"http://localhost:8081", "http://localhost:8082", "http://localhost:8083"}

	// 创建一个独立的 RoundRobinBalancer 实例
	balancer := NewRoundRobinBalancer(servers)

	var wg sync.WaitGroup
	numRequests := 100
	results := make([]string, numRequests)

	// 使用 channel 来确保顺序执行
	ch := make(chan struct{})

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			<-ch // 等待信号
			results[index] = balancer.NextServer()
		}(i)
	}

	// 依次发送信号以确保顺序执行
	for i := 0; i < numRequests; i++ {
		ch <- struct{}{}
	}

	wg.Wait()

	// 验证结果是否按照预期的顺序
	for i := 0; i < numRequests; i++ {
		expected := servers[i%len(servers)]
		if results[i] != expected {
			t.Errorf("Iteration %d: expected %s, got %s", i, expected, results[i])
		}
	}
}
