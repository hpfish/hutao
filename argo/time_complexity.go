package main

import "fmt"

func main() {
	nums := []int{10, 8, 9, 7, 11}
	count := bubbleSort(nums)
	fmt.Println(count)
	fmt.Println(nums)
}

// 冒泡排序
func bubbleSort(nums []int) int {
	count := 0
	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				tmp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = tmp
				count += 3
			}
		}
	}
	return count
}

func logarithmic(n float64) int {
	count := 0
	for n > 1 {
		n = n / 2
		count++
	}
	return count
}
