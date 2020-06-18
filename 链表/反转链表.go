package main

import (
	"encoding/json"
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	s := reverseList(&ListNode{Val:2, Next:&ListNode{Val:3, Next:&ListNode{Val:4, Next:&ListNode{Val:5}}}})
	ss, _ := json.Marshal(s)
	fmt.Printf(string(ss))
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for (cur != nil) {
		// prev 在cur后面, 擦
		// 是未来链表的 prev
		// 第一个是 反指向
		// 后面两个是 推进
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	fmt.Println(Json(cur))
	fmt.Println(Json(prev))
	return prev
}

func Json(v interface{}) string {
	s, _ := json.Marshal(v)
	return string(s)
}