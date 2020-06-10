package main

import (
	"fmt"
	"sort"
)

func main() {

	// 合并a 和  b
	// a为起始, b为结尾
	// a为起始, a 为结尾
	fmt.Println(merge([][]int{{1,3}, {2,6}, {5,6}}))
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := [][]int{}
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		m := merged[len(merged)-1]  // 合并后集合中的最后一个数组
		c := intervals[i]  // 待合并的数组
		fmt.Println(c, m)
		fmt.Println(merged)

		if c[0] > m[1] {
			merged = append(merged, c)
			continue
		}

		if c[1] > m[1] {
			m[1] = c[1]
		}
	}

	return merged
}


