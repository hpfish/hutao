package main

import "fmt"

func b() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return i
	//或者直接写成
	return
}
func main() {
	fmt.Println("return:", b())
	fmt.Println("return:", *(c()))
	a := new(map[string]string)
	fmt.Println(a)
	b := make(map[int]int, 2)
	fmt.Println(b)
}

func c() *int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return &i
}
