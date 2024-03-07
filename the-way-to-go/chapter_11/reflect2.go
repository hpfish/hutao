package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	//v.SetFloat(3.1415) //panic: reflect: reflect.Value.SetFloat using unaddressable value
	//v = v.Elem() // panic: reflect: call of reflect.Value.Elem on float64 Value
	fmt.Println("set ability of v:", v.CanSet())
	v = reflect.ValueOf(&x)
	fmt.Println("type of v:", v.Type())
	fmt.Println("set ability of v:", v.CanSet())
	v = v.Elem()
	fmt.Println("The Elem of v is:", v)
	fmt.Println("set ability of v:", v.CanSet())
	v.SetFloat(3.1415)
	fmt.Println(v.Interface())
	fmt.Println(v)
}
