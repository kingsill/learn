package main

import "fmt"

func main() {
	s := "abc"
	n := lengthOfLongestSubstring(s)
	fmt.Printf("%v", n)
}
func lengthOfLongestSubstring(s string) int {
	judg := make(map[byte]bool)
	left := 0
	right := 0
	maxlength := 0
	if len(s) <= 1 {
		return len(s)
	}
	// for n := 0; n < len(s); n++ {
	//不能这么使用，删除数据环节会使次序超出，因此使用right来进行当前遍历的数据记录
	for right < len(s) {
		if judg[s[right]] {
			delete(judg, s[left])
			left++

		} else {
			judg[s[right]] = true
			maxlength = max(maxlength, right-left+1)
			right++
		}

	}
	return maxlength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
