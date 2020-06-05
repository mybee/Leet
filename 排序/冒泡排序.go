package main

import "fmt"

func main() {
	// 可以看有没有交换发生, 提前结束
	s := []int{2, 5, 7, 3, 12}
	for i :=0; i< len(s); i++ {
		for j :=0; j < len(s)-i -1; j++ {
			if s[j+1] < s[j] {
				s[j+1], s[j] = s[j], s[j+1]
			}
		}
	}

	fmt.Println(s)

}

