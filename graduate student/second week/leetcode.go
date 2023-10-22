package main

import "fmt"

func main() {
	s := "abcabcaa"
	n := lengthOfLongestSubstring(s)
	fmt.Printf("%v", n)
}
func lengthOfLongestSubstring(s string) int {
	n := len(s) // 获得字符串长度
	if n <= 1 { // 对于字符串长度为0,1的可以直接返回了
		return n
	}

	// 使用哈希集合存储字符的出现情况
	// rune类型可以表示任意Unicode字符，并能存储一个32位的Unicode字符编码
	// byte和rune可以直接转换
	// 用map会比数组慢,毕竟要计算
	set := make(map[byte]bool)

	left := 0
	right := 0
	maxLength := 0

	for right < n {
		// 如果当前字符已经在集合中存在，则移动左指针，并将该字符从集合中移除
		if set[s[right]] {
			fmt.Printf("set1: %v\n", set)

			// 知道left=right，左右指针重合，将set清空 ×××大错特错，这条思路大错特错

			// 将第一个重复的数据排除后，继续进行循环
			delete(set, s[left])
			// set = make(map[byte]bool)  大错特错

			fmt.Printf("set2: %v\n", set)
			left++
			// left = right 大错特错
		} else {
			// 如果当前字符不在集合中，则将其加入集合，并更新最大长度
			set[s[right]] = true
			maxLength = max(maxLength, right-left+1)
			right++
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
