package main

import "fmt"

type pair struct {
	key int
	val string
}

type arrayHashMap struct {
	buckets []*pair
}

func newArrayHashMap() *arrayHashMap {
	buckets := make([]*pair, 100)
	return &arrayHashMap{buckets: buckets}
}

func (a *arrayHashMap) hashFunc(key int) int {
	index := key % 100
	return index
}

func (a *arrayHashMap) get(key int) interface{} {
	index := a.hashFunc(key)
	pair := a.buckets[index]
	if pair == nil {
		return nil
	}
	return pair.val
}

func (a *arrayHashMap) put(key int, val string) {
	pair := &pair{key: key, val: val}
	index := a.hashFunc(key)
	a.buckets[index] = pair
}

func (a *arrayHashMap) remove(key int) {
	index := a.hashFunc(key)
	a.buckets[index] = nil
}

func (a *arrayHashMap) pairSet() []*pair {
	var pairs []*pair
	for i := range a.buckets {
		if pairs != nil {
			pairs = append(pairs, a.buckets[i])
		}
	}
	return pairs
}

func (a *arrayHashMap) keySet() []int {
	var keys []int
	for _, pair := range a.buckets {
		if pair != nil {
			keys = append(keys, pair.key)
		}
	}
	return keys
}

func (a *arrayHashMap) valueSet() []interface{} {
	var values []interface{}
	for _, pair := range a.buckets {
		if pair != nil {
			values = append(values, pair.val)
		}
	}
	return values
}

func (a *arrayHashMap) print() {
	for _, pair := range a.buckets {
		if pair != nil {
			fmt.Println(pair.key, "->", pair.val)
		}
	}
}
