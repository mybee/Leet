package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//  3
	// 2  7
	//1 4
	//   5

	fmt.Println(preorderTraversal(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 7,
		},
	}))
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	doTraversal(root, &res)
	return res
}

func doTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	fmt.Println(res)
	doTraversal(root.Left, res)
	doTraversal(root.Right, res)
}
