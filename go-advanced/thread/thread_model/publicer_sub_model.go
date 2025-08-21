package thread_model

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type topicFunc func(any) bool

type subscriber chan any

type Publisher struct {
	rw          *sync.RWMutex
	buffer      int
	timeout     time.Duration
	subscribers map[subscriber]topicFunc
}

func NewPublisher(b int, t time.Duration) *Publisher {
	return &Publisher{
		rw:          &sync.RWMutex{},
		buffer:      b,
		timeout:     t,
		subscribers: make(map[subscriber]topicFunc),
	}
}

// subscribe all
func (p *Publisher) subscribeAll() chan any {
	return p.subscribeTopic(nil)
}

func (p *Publisher) subscribeTopic(topic topicFunc) chan any {
	p.rw.Lock()
	defer p.rw.Unlock()

	c := make(subscriber, p.buffer)
	p.subscribers[c] = topic
	return c
}

func (p *Publisher) Close() {
	p.rw.Lock()
	defer p.rw.Unlock()
	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

func (p *Publisher) Publishe(v any) {
	wg := sync.WaitGroup{}
	for sub, topic := range p.subscribers {
		wg.Add(1)
		publisherTopic(v, &wg, sub, topic, p.timeout)
	}
	wg.Wait()
}

func publisherTopic(v any, wg *sync.WaitGroup, sub subscriber, topic topicFunc, timeout time.Duration) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}
	select {
	case sub <- v:
	case <-time.After(timeout):
	}
}

func TestModel() {
	p := NewPublisher(10, 2*time.Second)

	all := p.subscribeAll()
	golang := p.subscribeTopic(func(v any) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publishe("hello world")
	p.Publishe("hello golang")
	p.Publishe("hello golang")
	p.Publishe("hello golang")
	p.Publishe("hello golang")

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done() // 确保 goroutine 结束时调用 Done
		for msg := range all {
			fmt.Println("all subscriber received:", msg)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done() // 确保 goroutine 结束时调用 Done
		for msg := range golang {
			fmt.Println("golang subscriber received:", msg)
		}
	}()
	p.Close()
	wg.Wait()
}
