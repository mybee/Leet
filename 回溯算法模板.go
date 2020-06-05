package main

import "fmt"

//function trackback(...arguments){
////1满足条件return
//if(xxx){
////满足了 咋整
////return
//}
//
////2剩余情况继续找
//for(var i=0;i<n;i++){
////3需要的地方枝解（就是跳过循环的意思）
//if(枝解条件)
//continue
//
////4满足情况操作
//操作【例如push/增加字符串末位/哈希表修改值】
////5进行相应操作后继续递归寻找
//trackback(...arguments)
////回溯（退回上级操作）
//与第4步相反【例如pop()/删除字符串末位/哈希表修回值】
//}
//}



func main() {
	a := []int{1, 2, 3}
	fmt.Println(permute(a))
}
// 塔罗兔 回溯 模板
// 时间复杂度 和 空间复杂度都是 n 的阶乘

// TODO  swap 操作

func dfs(nums []int, depth int, path []int, used []bool, res [][]int) [][]int {
	fmt.Println(nums, depth, path, used)
	if len(nums) == depth {
		//tmp := make([]int, len(path))
		//copy(tmp, path)
		res = append(res, path)
		//fmt.Println(res)
		return res
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		// 下面 回溯部分对称操作
		used[i] = true
		path = append(path, nums[i])
		res = dfs(nums, depth+1, path, used, res)
		used[i] = false
		path = path[:len(path)-1]
	}
	//fmt.Println(res)
	return res
}

func permute(nums []int) [][]int {

	res := [][]int{}
	if len(nums) == 0 {
		return res
	}

	path := []int{}
	used := make([]bool, len(nums))
	res = dfs(nums, 0, path, used, res)

	return res
}