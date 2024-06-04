package pubsub

import (
	"sync"
	"time"
)

type (
	subscriber chan interface{}
	topicFunc  func(v interface{}) bool
)

type Publisher struct {
	buffer     int
	timeout    time.Duration
	subscribes map[subscriber]topicFunc
	m          sync.RWMutex
}

func NewPublisher(buffer int, timeout time.Duration) *Publisher {
	return &Publisher{
		buffer:     buffer,
		timeout:    timeout,
		subscribes: make(map[subscriber]topicFunc),
	}
}

func (p *Publisher) SubscribeAll() chan interface{} {
	return p.SubscribeTopic(nil)
}

func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{})
	p.m.Lock()
	p.subscribes[ch] = topic
	p.m.Unlock()
	return ch
}

func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	for sub := range p.subscribes {
		delete(p.subscribes, sub)
		close(sub)
	}
}

func (p *Publisher) Publishe(v interface{}) {
	var wg sync.WaitGroup
	for sub := range p.subscribes {
		wg.Add(1)
		go p.sendTopic(v, sub, p.subscribes[sub], &wg)
	}
	wg.Wait()
}

func (p *Publisher) sendTopic(v interface{}, sub subscriber, topic topicFunc, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}

	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}
