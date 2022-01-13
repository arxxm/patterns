package strategy

import (
	"fmt"
)

type evictionAlgo interface {
	Evict(c *cache)
}

type Fifo struct { //first in first out
}

func (l *Fifo) Evict(c *cache) {
	fmt.Println("Evicting by fifo strategy")
}

type Lru struct { // last recently used
}

func (l *Lru) Evict(c *cache) {
	fmt.Println("Evicting by lru strategy")
}

type Lfu struct {
}

func (l *Lfu) Evict(c *cache) {
	fmt.Println("Evicting by lfu strategy")
}

type cache struct {
	storage      map[string]string
	evictionAlgo evictionAlgo
	capacity     int
	maxCapacity  int
}

func InitCache(e evictionAlgo) *cache {
	storage := make(map[string]string)
	return &cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *cache) SetEvicionAlgo(e evictionAlgo) {
	c.evictionAlgo = e
}

func (c *cache) Add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *cache) Set(key string) {
	delete(c.storage, key)
}

func (c *cache) evict() {
	c.evictionAlgo.Evict(c)
	c.capacity--
}
