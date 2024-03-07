package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		n, result := fibonacci2(i)
		fmt.Printf("fibonacci(%d) is: %d\n", n, result)
	}
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

func fibonacci2(n int) (i, res int) {
	i = n
	if n <= 1 {
		res = 1
	} else {
		_, res1 := fibonacci2(n - 1)
		_, res2 := fibonacci2(n - 2)
		res = res1 + res2
	}
	return
}
