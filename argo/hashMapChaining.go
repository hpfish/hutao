package main

import (
	"fmt"
	"strconv"
	"strings"
)

type hashMapChaining struct {
	size        int      // 键值对数量
	capacity    int      // 哈希表容量
	loadThres   float64  // 触发扩容的负载因子阈值
	extendRatio int      // 扩容倍数
	buckets     [][]pair // 桶数组
}

func newHashMapChaining() *hashMapChaining {
	buckets := make([][]pair, 4)
	for i := 0; i < 4; i++ {
		buckets[i] = make([]pair, 0)
	}

	return &hashMapChaining{
		size:        0,
		capacity:    4,
		loadThres:   2.0 / 3.0,
		extendRatio: 2,
		buckets:     buckets,
	}
}

func (m *hashMapChaining) hashFunc(key int) int {
	return key % m.capacity
}

func (m *hashMapChaining) loadFactor() float64 {
	return float64(m.size) / float64(m.capacity)
}

func (m *hashMapChaining) get(key int) string {
	idx := m.hashFunc(key)
	bucket := m.buckets[idx]
	for _, p := range bucket {
		if p.key == key {
			return p.val
		}
	}
	return ""
}

func (m *hashMapChaining) put(key int, val string) {
	if m.loadFactor() > m.loadThres {
		m.extend()
	}
	idx := m.hashFunc(key)
	for i := range m.buckets[idx] {
		if m.buckets[idx][i].key == key {
			m.buckets[idx][i].val = val
			return
		}
	}
	p := pair{
		key: key,
		val: val,
	}
	m.buckets[idx] = append(m.buckets[idx], p)
	m.size += 1
}

func (m *hashMapChaining) remove(key int) {
	idx := m.hashFunc(key)
	for i, p := range m.buckets[idx] {
		if p.key == key {
			m.buckets[idx] = append(m.buckets[idx][:i], m.buckets[idx][i+1:]...)
			m.size -= 1
			break
		}

	}
}

func (m *hashMapChaining) extend() {
	tmpBuckets := make([][]pair, len(m.buckets))
	for i := 0; i < len(m.buckets); i++ {
		tmpBuckets[i] = make([]pair, len(m.buckets[i]))
		copy(tmpBuckets[i], m.buckets[i])
	}

	m.capacity *= m.extendRatio
	m.buckets = make([][]pair, m.capacity)
	for i := 0; i < m.capacity; i++ {
		m.buckets[i] = make([]pair, 0)
	}
	m.size = 0
	for _, bucket := range tmpBuckets {
		for _, p := range bucket {
			m.put(p.key, p.val)
		}
	}
}

func (m *hashMapChaining) print() {
	var builder strings.Builder

	for _, bucket := range m.buckets {
		builder.WriteString("[")
		for _, p := range bucket {
			builder.WriteString(strconv.Itoa(p.key) + "->" + p.val + " ")
		}
		builder.WriteString("]")
		fmt.Println(builder.String())
		builder.Reset()
	}
}
