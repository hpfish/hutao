package main

import "fmt"

func main() {
	fmt.Println(loss(" abcdadaterw a"))
}

func lengthOfLongestSubstring(s string) int {
	// 记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)

	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格， 移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断移动右指针
			m[s[rk+1]]++
			rk++
		}
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func loss(s string) int {
	var m = make(map[byte]int)
	n := len(s)

	lp, longest := -1, 0
	for i := 0; i < n; i++ {
		if i > 0 {
			delete(m, s[i-1])
		}

		for lp+1 < n && m[s[lp+1]] == 0 {
			m[s[lp+1]]++
			lp++
		}
		longest = max(longest, lp-i+1)
	}
	return longest
}
