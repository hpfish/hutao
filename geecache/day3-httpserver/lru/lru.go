package lru

import "container/list"

type Cache struct {
	maxBytes int64                    // 最大内存
	nbytes   int64                    // 当前已使用内存
	ll       *list.List               // 标准库实现的双向链表
	cache    map[string]*list.Element // 值是双向链表中对应节点的指针

	OnEvicted func(key string, value Value) // 某条记录被移除时的回调函数，可以为nil
}

// entry 是双向链表节点的数据类型，在链表中仍保存每个值对应的key的好处在于，淘汰队首节点时
// 需要用key从字典中删除对应的映射
type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int // 用户返回值所占用的内存大小
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		// 双向链表作为队列， 队首队尾是相对的，这里约定 front 为队尾
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		// 从字典 c.cache 中删除该节点的映射关系
		delete(c.cache, kv.key)
		// 更新当前使用的内存： 减去key和value的内存
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}

	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
