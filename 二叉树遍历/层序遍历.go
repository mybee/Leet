package main


func main() {


}


func levelOrder(root *TreeNode) (result [][]int) {
	var dfs func(*TreeNode, int)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level >= len(result) {
			result = append(result, []int{})
		}
		result[level] = append(result[level], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 0)
	return result
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil{
		return res
	}
	var queue = []*TreeNode{root}
	var level int
	for len(queue) > 0 {
		counter := len(queue)
		res = append(res,[]int{})
		for 0 < counter {
			counter--
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			res[level] = append(res[level], queue[0].Val)
			queue = queue[1:]
		}
		level++
	}
	return res
}

//作者：linbingyuan
//链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal/solution/ceng-ci-bian-li-by-linbingyuan/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。