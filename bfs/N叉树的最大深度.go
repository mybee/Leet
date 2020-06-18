package main

import (
	"container/list"
	"fmt"
)

//Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func main() {
	fmt.Println(maxDepth(&Node{
		Val: 10,
		Children: []*Node{
			&Node{Val: 2,
				Children: []*Node{
					&Node{Val: 2,
						Children: []*Node{}},
				}},
		},
	}))
}

// bfs 算法
// 使用队列
// 模板
// 1. 新建 list
// 2. for len {
// 3. 把root 放进去
// 4.

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushFront(root)
	level := 0
	for queue.Len() > 0 {  // 直到没有数据放到队列中
		l := queue.Len()

		for i := 0; i < l; i++ {
			node := queue.Remove(queue.Back()).(*Node)
			for j, _ := range node.Children {
				queue.PushFront(node.Children[j])
			}
		}
		level++
	}

	return level
}
