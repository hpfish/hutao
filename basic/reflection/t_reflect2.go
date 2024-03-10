package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.ValueOf(1)
	i := v.Interface().(int)
	fmt.Println(i)

	b := 1
	v2 := reflect.ValueOf(&b)
	//v2.SetInt(10)
	v2.Elem().SetInt(10)
	fmt.Println(b)
}
