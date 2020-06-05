package main

func main() {


}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	doTraversal(root, &res)
	return res
}

func doTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	doTraversal(root.Left, res)
	*res = append(*res, root.Val)
	doTraversal(root.Right, res)
}