package main

import "fmt"

func main() {
	// 可以看有没有交换发生, 提前结束
	s := []int{2, 5, 7, 3, 12}

	fmt.Println(MergeSort(s))

}

// 时间复杂度: nlog(n)
// 空间复杂度: n
func MergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	m := len(a) / 2
	l := MergeSort(a[:m])
	r := MergeSort(a[m:])
	return merge(l, r)
}
// 合并两个有序数组的代码
// 只有一个元素的数组也是有序数组
func merge(a, b []int) (c []int) {
	fmt.Println(a, b)
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}
	c = append(c, a[i:]...)
	c = append(c, b[j:]...)
	return
}

