package main

import "fmt"

func main() {
	returnAndDefer()
	fmt.Println("returnAndDeferValue", returnAndDeferValue())
	fmt.Println("returnAndDeferReferenceValue", returnAndDeferReferenceValue())
	fmt.Println("returnAndDeferNameValue", returnAndDeferNameValue())
	fmt.Println("f2", f2())
}

func deferFunc() int {
	fmt.Println("defer func called")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called")
	return 0
}

func returnAndDefer() int {

	defer deferFunc()

	return returnFunc()
}

func returnAndDeferValue() int {
	var a int
	defer func() {
		a = 10
	}()
	a = 2
	return a
}

func returnAndDeferReferenceValue() interface{} {
	var a interface{}
	defer func() {
		a = 10
	}()
	a = 2
	return a
}

func returnAndDeferNameValue() (t int) {
	t = 1
	defer func() {
		t = t * 10
	}()
	return t
}

func f2() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)

	return 1
}
