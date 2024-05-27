package lba

import (
	"sync/atomic"
)

// 实现一个轮询算法的负载均衡器
type RoundRobinBalancer struct {
	servers []string
	current uint64
}

func NewRoundRobinBalancer(servers []string) *RoundRobinBalancer {
	return &RoundRobinBalancer{
		servers: servers,
		current: 0,
	}
}

func (r *RoundRobinBalancer) NextServer() string {
	current := atomic.AddUint64(&r.current, 1) - 1 // 原子地增加计数器并获取当前值
	index := int(current) % len(r.servers)
	return r.servers[index]
}
