package main

import (
	"flag"
	"fmt"
)

var ngoroutine = flag.Int("n", 100000, "how many goroutines")

func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	flag.Parse()
	leftmost := make(chan int)
	var left, right chan int = nil, leftmost
	for i := 0; i < *ngoroutine; i++ {
		left, right = right, make(chan int)
		go f(left, right)
	}
	// 不是最初循环的那个right， 而是最终循环的 right
	// 当主线程 right <- 0 执行时，类似于递归函数在最内层产生返回值一般
	right <- 0
	x := <-leftmost
	fmt.Println(x)
}
