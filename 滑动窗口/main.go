package main

import "fmt"

// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//示例 1:
//
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func lengthOfLongestSubstring(s string) int {
	res := 0
	window := [256]int{0}
	l,r := 0,0
	for ;r<len(s);r++ {
		fmt.Println(s[r])
		window[s[r]]++  // 字母计数加一
		for window[s[r]]==2 {
			fmt.Println(res, l, r, s, window)
			res = max(res,r-l)
			window[s[l]]--  // 字母计数减一
			// 窗口左边向右收缩一格
			l++
		}

	}
	return max(res,r-l)
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
	s := "abcabcbb"
	fmt.Print(lengthOfLongestSubstring(s))
}
