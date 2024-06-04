package lru

import (
	"sync"
)

type Cache struct {
	ll       *DoublyLinkedList
	maxBytes int64 //Cache max bytes
	nBytes   int64 // Cache current bytes
	cache    map[string]*Node
	mu       sync.Mutex

	OnEvicted func(key string, value Value) //executed when an entry is purged
}

type Value interface {
	Len() int
}

type entry struct {
	key   string
	value Value
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		OnEvicted: onEvicted,
		cache:     make(map[string]*Node),
		ll:        NewList(),
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}

func (c *Cache) Get(key string) (Value, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if node, ok := c.cache[key]; ok {
		c.ll.MoveToFront(node)
		if entry, ok := node.Value.(*entry); ok {
			// log.Printf("lru Cache get %d", entry.value)
			return entry.value, true
		}
	}
	return nil, false
}

func (c *Cache) Add(key string, value Value) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if node, ok := c.cache[key]; ok {
		c.ll.MoveToFront(node)
		kv := node.Value.(*entry)
		c.nBytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushNodeToFront(&entry{key, value})
		c.cache[key] = ele
		c.nBytes += int64(len(key)) + int64(value.Len())
	}

	// 如果超出了最大字节数，执行清理操作
	for c.maxBytes != 0 && c.maxBytes < c.nBytes {
		last, _ := c.ll.GetBack()
		if last == nil {
			break
		}
		c.ll.RemoveNode(last)
		kv := last.Value.(*entry)
		delete(c.cache, kv.key)
		c.nBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

func (c *Cache) RemoveOldest() {
	c.mu.Lock()
	defer c.mu.Unlock()
	ele, _ := c.ll.GetBack()
	if ele != nil {
		c.ll.RemoveNode(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}
