package main

import "fmt"

func main() {
	var f = Adder2()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
}

func Adder2() func(int) int {
	var x int // 闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量。
	fmt.Printf("x=%d\n", x)
	return func(delta int) int {
		x += delta
		return x
	}
}
