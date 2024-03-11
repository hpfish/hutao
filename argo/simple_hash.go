package main

import "fmt"

func addHash(key string) int {
	var (
		hash    int64
		modules int64 = 1000000007
	)

	for _, b := range []byte(key) {
		hash = (hash + int64(b)) % modules
	}
	return int(hash)
}

func mulHash(key string) int {
	var (
		hash    int64
		modules int64 = 1000000007
	)
	for _, b := range []byte(key) {
		hash = (31*hash + int64(b)) % modules
	}
	return int(hash)
}

// 异或哈希
func xorHash(key string) int {
	var (
		hash    = 0
		modules = 1000000007
	)
	for _, b := range []byte(key) {
		hash ^= int(b)
		hash = (31*hash + int(b)) % modules
	}
	return hash & modules
}

// 旋转哈希
func rotHash(key string) int {
	var (
		hash    int64
		modules int64 = 1000000007
	)
	for _, b := range []byte(key) {
		hash = ((hash << 4) ^ (hash >> 28) ^ int64(b)) % modules
	}
	return int(hash)
}

func main() {
	fmt.Println(addHash("abc"))
	fmt.Println(addHash("abc"))
	fmt.Println(addHash("ccc"))
	fmt.Println(mulHash("ccc"))
	fmt.Println(xorHash("ccc"))
	fmt.Println(rotHash("ccc"))
}
