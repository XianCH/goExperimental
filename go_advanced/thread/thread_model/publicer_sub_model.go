package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type (
	subscriber chan any         //订阅通道
	topicFunc  func(v any) bool //订阅过滤方法
)

// 发布者对象
type Publisher struct {
	m           *sync.RWMutex
	buffer      int
	timeOut     time.Duration
	subscribers map[subscriber]topicFunc
}

// 创建发布者
func NewPublisher(publishTimeOut time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeOut:     publishTimeOut,
		subscribers: make(map[subscriber]topicFunc),
	}
}

// 添加一个新的订阅者,订阅所有主题
func (p *Publisher) Subscribe() chan any {
	return p.SubscribeTopic(nil)
}

func (p *Publisher) SubscribeTopic(topic topicFunc) chan any {
	ch := make(chan any, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

// 退出订阅
func (p *Publisher) Evict(sub chan any) {
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscribers, sub)
	close(sub)
}

// 发布一个主题
func (p *Publisher) Publish(v any) {
	p.m.RLock()
	defer p.m.RUnlock()

	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

// 发送一个主题
func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v any, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}

	select {
	case sub <- v:
	case <-time.After(p.timeOut):
	}
}

// 关闭发布者对象，同时关闭所有发布者通道
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(
			p.subscribers, sub)
		close(sub)
	}
}

func main() {
	p := NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	all := p.Subscribe()
	golang := p.SubscribeTopic(func(v any) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})
	p.Publish("hello,world")
	p.Publish("hello.golang")

	go func() {
		for msg := range all {
			fmt.Println("all", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang", msg)
		}
	}()
}
