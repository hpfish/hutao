package main

import "fmt"

type StorageType int

const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemoryStorage
)

type Store interface {
	save()
}

func newTempStorage() tempStore {
	return tempStore{}
}

func newMemoryStore() memoryStore {
	return memoryStore{}
}

func newDiskStore() diskStore {
	return diskStore{}
}

type memoryStore struct {
}

func (s memoryStore) save() {
	fmt.Println("memoryStore")
}

type diskStore struct {
}

func (s diskStore) save() {
	fmt.Println("diskStore")
}

type tempStore struct {
}

func (s tempStore) save() {
	fmt.Println("tempStore")
}

func NewStore(t StorageType) Store {
	switch t {
	case MemoryStorage:
		return newMemoryStore()
	case DiskStorage:
		return newDiskStore()
	default:
		return newTempStorage()
	}
}

func main() {
	s := NewStore(MemoryStorage)
	s.save()
}
