package main

import "math/rand"

func randomNumbers(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	return nums
}

func findOne(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			return i
		}
	}
	return -1
}
