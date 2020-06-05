package main

import "fmt"

func main() {

	fmt.Println(combine(4, 2))
}

func combine(n int, k int) [][]int {

	var res [][]int

	var list []int

	return dfs(n, k, 1, list, res)

}

func dfs(n, k int, pos int, list []int, res [][]int) [][]int {
	if len(list) == k {
		tmp := make([]int, k)
		copy(tmp, list)
		res = append(res, tmp)
		//fmt.Println(list)
		return res
	}

	for i := pos; i <= n; i++ {
		list = append(list, i)
		res = dfs(n, k, i+1, list, res)
		list = list[:len(list) -1 ]
	}
	return res
}