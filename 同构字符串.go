package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isIsomorphic("aag", "egg"))
}

//思路一
//
//常规解法，使用哈希映射，两个字符串相互映射。


// 思路二
//参考讨论区里的大佬，比较巧妙，对比两个字符串对应位置的字符在字符串内第一次出现的位置。

func isIsomorphic(s string, t string) bool {
	for i:=0; i< len(s); i++ {
		if strings.Index(s, s[i:i+1]) != strings.Index(t, t[i:i+1]) {
			return false
		}
	}
	return true
}