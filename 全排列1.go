package main

import "fmt"

func main()  {
	a := []int{1, 2, 3, 4}
	fmt.Println(permute(a))
}

func permute(nums []int) [][]int {

	var res [][]int
	if len(nums) == 0 {
		return res
	}

	path := []int{}
	used := make([]bool, len(nums))

	res = dfs(nums, path, used, res)
	return  res
}

func dfs(nums []int, path []int, used []bool, res[][]int) [][]int {
	fmt.Println(nums, path, used)

	if len(path) == len(nums) {
		// 需要copy 不然会有重复问题
		tmp := make([]int, len(nums))
		copy(tmp, path)
		res = append (res, tmp)
		fmt.Println("qeqeq", res)
		return res
	}

	for i := 0; i< len(nums); i++ {
		if used[i]  {
			continue
		}

		used[i] = true
		path = append(path, nums[i])
		res = dfs(nums, path, used, res)
		used[i] = false
		path = path[:len(path)-1]
	}
	return res
}