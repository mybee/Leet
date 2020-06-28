package main

import (
	"encoding/json"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	l1 := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:2,
			Next:&ListNode{
				Val:4,
				Next:nil,
			},
		},
	}

	l2 := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:3,
			Next:&ListNode{
				Val:5,
				Next:nil,
			},
		},
	}

	l := mergeTwoList(l1, l2)

	fmt.Println(Json(l))

}

func mergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {


	// 理解思路:

	// 尾部是: 谁到头, 返回另一个
	// 递归部分 只是将 节点连起来

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val > l2.Val {

		l2.Next = mergeTwoList(l1, l2.Next)
		fmt.Println(Json(l2.Next))
		return l2

	} else {

		l1.Next = mergeTwoList(l1.Next, l2)
		fmt.Println(Json(l1.Next))
		return l1

	}

}

func Json(v interface{}) string  {
	s, _ := json.Marshal(v)
	return string(s)
}