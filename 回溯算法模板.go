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

// 所有的排列组和

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(permute(a))
}
// 塔罗兔 回溯 模板
// 时间复杂度 和 空间复杂度都是 n 的阶乘
// 结果集 也是 n的 阶乘

// TODO  swap 操作

func dfs(nums []int, depth int, path []int, used []bool, res [][]int) [][]int {
	fmt.Println(depth, path, used)
	//0 [] [false false false false false]
	//1 [1] [true false false false false]
	//2 [1 2] [true true false false false]
	//3 [1 2 3] [true true true false false]
	//4 [1 2 3 4] [true true true true false]
	//5 [1 2 3 4 5] [true true true true true]
	//4 [1 2 3 5] [true true true false true]
	//5 [1 2 3 5 4] [true true true true true]
	//3 [1 2 4] [true true false true false]
	if len(nums) == depth { // 达到最大深度
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
		// 流程:  先一次执行到底, 然后慢慢的往上浮
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




