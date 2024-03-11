package main

import (
	"fmt"
	"unsafe"
)

type Object1 struct {
	b byte
	s struct{}
}

type Object2 struct {
	n int64
	s struct{}
}

type Object3 struct {
	n int16
	m int16
	s struct{}
}

type Object4 struct {
	n int16
	m int64
	s struct{}
}

var zerobase uintptr

func main() {
	o1 := Object1{}
	fmt.Println(unsafe.Sizeof(o1))
	o2 := Object2{}
	fmt.Println(unsafe.Sizeof(o2))
	o3 := Object3{}
	fmt.Println(unsafe.Sizeof(o3))
	o4 := Object4{}
	fmt.Println(unsafe.Sizeof(o4))

	fmt.Println(unsafe.Sizeof(zerobase))
}
