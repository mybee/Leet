package main

import "fmt"

func main() {
	// 插入排序
	// 最好: on 最差 on2
	// 稳定的排序算法

	s := []int{2, 5, 7, 4, 1, 7, 0, 1}

	for i := 1; i < len(s); i++ {
		for j := i; j > 0 && s[j] > s[j-1]; j-- {
			s[j-1], s[j] = s[j], s[j-1]
		}
	}

	fmt.Println(s)
}
