package main

import (
	"fmt"
	"sync"
	"time"
)

// A safe Counter
type Counter struct {
	mu    sync.RWMutex
	count uint64
}

// Readers and writers lock
func main() {
	var counter Counter
	for i := 0; i < 10; i++ { // 10 readers
		go func() {
			for {
				counter.Count() // Counter read
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for { // one writer
		counter.Incr() // increase
		time.Sleep(time.Second)
	}
}

// Use write lock protect
func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
	fmt.Println("Write counter:", c.count)
}

// Use read lock protect
func (c *Counter) Count() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}
