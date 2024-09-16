package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type pair struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	if ele, found := c.cache[key]; found {
		c.list.MoveToFront(ele)
		return ele.Value.(pair).value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if ele, found := c.cache[key]; found {
		c.list.MoveToFront(ele)
		ele.Value = pair{key, value}
		return
	}
	if c.list.Len() == c.capacity {
		ele := c.list.Back()
		if ele != nil {
			delete(c.cache, ele.Value.(pair).key)
			c.list.Remove(ele)
		}
	}
	ele := c.list.PushFront(pair{key, value})
	c.cache[key] = ele
}

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}