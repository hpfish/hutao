package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var arr [5]int
	fmt.Println(arr)
	var arr2 []int
	fmt.Println(arr2)
	var nums = []int{1, 2, 3, 4, 5}
	insert(nums, 4, 2)
	fmt.Println(nums)
	//var i = -1
	//fmt.Println(arr[i])
	fmt.Println(nums[:])

}

func randomAccess(nums []int) int {
	randomIndex := rand.Intn(len(nums))
	randomNum := nums[randomIndex]
	return randomNum
}

func insert(nums []int, num int, index int) {
	for i := len(nums) - 1; i < index; i-- {
		nums[i] = nums[i-1]
	}
	nums[index] = num
}

func remove(nums []int, index int) {
	for i := index; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
}

func traverse(nums []int) {
	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
	}
	count = 0
	for _, num := range nums {
		count += num
	}
	count = 0
	for i, num := range nums {
		count += nums[i]
		count += num
	}
}

func find(nums []int, target int) int {
	var index = 0
	for i, num := range nums {
		if num == target {
			index = i
			break
		}
	}
	return index
}

func extend(nums []int, enlarge int) []int {
	res := make([]int, len(nums)+enlarge)
	for i, num := range nums {
		res[i] = num
	}
	return res
}
