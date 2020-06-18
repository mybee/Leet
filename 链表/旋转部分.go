package main

import (
	"encoding/json"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 6月12日面试内容

func main() {
	s := reverseeBetween(&ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}, 1, 3)
	ss, _ := json.Marshal(s)
	fmt.Printf(string(ss))
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {

	if head == nil {
		return nil
	}

	var cur, prev *ListNode = head, nil
	for m > 1 {
		prev, cur = cur, cur.Next
		m--
		n--
	}

	con, tailf := prev, cur

	fmt.Println(prev)

	for n > 0 {
		cur.Next, prev, cur = prev, cur, cur.Next
		n--
	}

	//->-> <-<- ->->

	if con != nil {
		con.Next = prev
	} else {
		head = prev
	}

	tailf.Next = cur
	return head
}

func Json(v interface{}) string {
	s, _ := json.Marshal(v)
	return string(s)
}

func reverseeBetween(head *ListNode, m int, n int) *ListNode {

	if head == nil {
		return nil
	}

	var cur, prev *ListNode = head, nil

	// 循环到 m 点附近
	for m > 1 {
		cur, prev = cur.Next, cur
		m--
		n--
	}

	mPrev, mCur := prev, cur

	// 反转到 n 点附近
	for n > 0 {
		cur.Next, cur, prev = prev, cur.Next, cur
		n--
	}

	// 这里很重要
	mCur.Next = cur

	if mPrev == nil {
		head = prev
	} else  {
		mPrev.Next = prev
	}

	return head
}
