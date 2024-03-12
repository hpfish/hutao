package main

import "fmt"

func main() {
	s := "  c"
	fmt.Println(lengthOfLongestSubstring(s))
	fmt.Println(lengthOfLongestSubstringOffice(s))
}

func lengthOfLongestSubstring(s string) int {
	var (
		length      = 0
		j           = 0
		emptyString bool
		totalChars  = make(map[string]bool)
	)
	for i, c := range s {
		str := string(c)
		if str == " " {
			j = i
			if !emptyString {
				length += 1
				emptyString = true
				continue
			}
		}
		if !totalChars[str] {
			totalChars[str] = true
		} else {
			cLength := i - j
			j = i
			if cLength > length {
				length = cLength
			}
		}

	}
	return length
}

func lengthOfLongestSubstringOffice(s string) int {
	m := map[byte]int{}
	n := len(s)

	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		ans = max(ans, rk-i+1)
	}
	return ans
}
