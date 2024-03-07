package main

import (
	"fmt"
	"reflect"
	"strings"
)

// Map示例 -- 根据规则处理完再返回

func MapStrToStr(arr []string, fn func(s string) string) []string {
	var newArray []string
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

func MapStrToInt(arr []string, fn func(s string) int) []int {
	var newArray []int
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

// Reduce 示例 根据规则合并成一个返回
func Reduce(arr []string, fn func(s string) int) int {
	sum := 0
	for _, it := range arr {
		sum += fn(it)
	}
	return sum
}

// Filter 示例 根据规则规律
func Filter(arr []int, fn func(n int) bool) []int {
	var newArray = []int{}
	for _, it := range arr {
		if fn(it) {
			newArray = append(newArray, it)
		}
	}
	return newArray
}

func Map(data interface{}, fn interface{}) []interface{} {
	vfn := reflect.ValueOf(fn)
	vdata := reflect.ValueOf(data)
	result := make([]interface{}, vdata.Len())

	for i := 0; i < vdata.Len(); i++ {
		result[i] = vfn.Call([]reflect.Value{vdata.Index(i)})[0].Interface()
	}

	return result
}

func main() {
	var list = []string{"Hao", "Chen", "MegaEase"}

	x := MapStrToStr(list, func(s string) string {
		return strings.ToUpper(s)
	})
	fmt.Printf("%v\n", x) //[HAO CHEN MEGAEASE]

	y := MapStrToInt(list, func(s string) int {
		return len(s)
	})
	fmt.Printf("%v\n", y) //[3 4 8]

	z := Reduce(list, func(s string) int {
		return len(s)
	})
	fmt.Printf("%v\n", z) // 15

	var intset = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	out := Filter(intset, func(n int) bool { return n%2 == 1 })

	fmt.Printf("%v\n", out) //[1 3 5 7 9]

	out = Filter(intset, func(n int) bool {
		return n > 9
	})
	fmt.Printf("%v\n", out)

	square := func(x int) int {
		return x * x
	}
	nums := []int{1, 2, 3, 4}
	squaredArr := Map(nums, square)
	fmt.Println(squaredArr)
}

func transform(slice, function interface{}, inPlace bool) interface{} {

	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("transform： not slice")
	}

	fn := reflect.ValueOf(function)
	// 返回元素类型
	elemType := sliceInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, nil) {
		panic("trasform: function must be of type func(" + sliceInType.Type().Elem().String() + ") outputElemType")
	}

	sliceOutType := sliceInType
	if !inPlace {
		sliceOutType = reflect.MakeSlice(reflect.SliceOf(fn.Type().Out(0)), sliceInType.Len(), sliceInType.Len())
	}

	for i := 0; i < sliceOutType.Len(); i++ {
		sliceOutType.Index(i).Set(fn.Call([]reflect.Value{sliceInType.Index(i)})[0])
	}
	return sliceOutType.Interface()
}

func verifyFuncSignature(fn reflect.Value, types ...reflect.Type) bool {
	if fn.Kind() != reflect.Func {
		return false
	}

	if (fn.Type().NumIn() != len(types)-1) || (fn.Type().NumOut() != 1) {
		return false
	}

	for i := 0; i < len(types)-1; i++ {
		if fn.Type().In(i) != types[i] {
			return false
		}
	}

	outType := types[len(types)-1]
	if outType != nil && fn.Type().Out(0) != outType {
		return false
	}
	return true
}
