package main

import "container/list"

type entry struct {
	key   int
	value int
}

type LRUCache struct {
	capacity  int
	list      *list.List
	keyToNode map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, list.New(), map[int]*list.Element{}}
}

func (c *LRUCache) Get(key int) int {
	node := c.keyToNode[key]
	if node == nil {
		return -1
	}
	c.list.MoveToFront(node)
	return node.Value.(entry).value
}

func (c *LRUCache) Push(key, value int) {
	if node := c.keyToNode[key]; node != nil {
		node.Value = entry{key: value}
		c.list.MoveToFront(node)
	}

	c.keyToNode[key] = c.list.PushFront(entry{key: value})
	if len(c.keyToNode) > c.capacity {
		delete(c.keyToNode, c.list.Remove(c.list.Back()).(entry).key)
	}
}
