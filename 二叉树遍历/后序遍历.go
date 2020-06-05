package main

import "fmt"

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

func main() {
	//  3
	// 2  7
	//1 4
	//   5

	fmt.Println(postorderTraversal(&treeNode{
		Val: 3,
		Left: &treeNode{
			Val: 2,
			Left: &treeNode{
				Val: 1,
			},
			Right: &treeNode{
				Val: 4,
				Left: &treeNode{
					Val: 5,
				},
			},
		},
		Right: &treeNode{
			Val: 7,
		},
	}))
}

func postorderTraversal(root *treeNode) []int {
	return traversal(root, []int{})
}

func traversal(root *treeNode, ans []int) []int {
	if root == nil {
		return ans
	}
	ans = traversal(root.Left, ans)
	ans = traversal(root.Right, ans)
	ans = append(ans, root.Val)
	return ans
}
