package main

import (
	"fmt"
)

// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
//push(x) —— 将元素 x 推入栈中。
//pop() —— 删除栈顶的元素。
//top() —— 获取栈顶元素。
//getMin() —— 检索栈中的最小元素。
//
//链接：https://leetcode-cn.com/problems/min-stack


// 使用一个栈
type Node struct {
	Min int
	Val int
}

// 不使用辅助栈
type MinStack struct{
	stack []*Node
}

func Constructor() MinStack{
	return MinStack{[]*Node{}}
}

func (this *MinStack) Push(x int){
	if len(this.stack) == 0 {
		this.stack = append(this.stack, &Node{Min: x, Val:x})
		return
	}
	min := this.GetMin()
	if min > x {
		min = x
	}

	this.stack = append(this.stack, &Node{Min: min, Val:x})

}

func (this *MinStack) Pop(){ // 删除栈顶元素
	// 题设是pop、top、getMin总是在非空栈上操作
	// 这里是减一
	this.stack = this.stack[:len(this.stack) - 1]
}

func (this *MinStack) Top() int{ // 获取栈顶元素
	return this.stack[len(this.stack)- 1].Val
}

func (this *MinStack) GetMin() int{ // 获取最小值

	return this.stack[len(this.stack) - 1].Min
}

func main() {

	minStack := &MinStack{}
	minStack.Push(-2);
	minStack.Push(0);
	minStack.Push(-3);
	for _, v :=range minStack.stack {
		fmt.Println(v.Min, v.Val)
	}
	fmt.Println(minStack.GetMin()); //  --> 返回 -3.
	minStack.Pop()
	fmt.Println(minStack.Top());      //--> 返回 0.
	fmt.Println(minStack.GetMin());   //--> 返回 -2.
	fmt.Println(minStack.stack)

}


//作者：xusanduo08
//链接：https://leetcode-cn.com/problems/min-stack/solution/goban-ben-bu-shi-yong-fu-zhu-zhan-ji-shi-yong-fu-z/

// 使用两个栈

// type MinStack struct {
//	values []int
//	min    []int
//}
//
///*
//利用栈的先进后出的特性
//使用两个栈，values栈保持我们的元素，min栈保持我们的最小元素
//--------------------------------------
//这里需要注意的是再进行push操作时
//我们需要注意两点
//1、min是否为空  2、x元素是否小于栈顶元素
//满意以上任意一条我们就在push操作时
//将x元素push道values的同时也push到min
//这样会使最小的元素始终是min栈的栈顶元素
//取最小值时，直接取min栈的栈顶元素即可
//--------------------------------------
//还需要注意的是，在进行pop操作时
//我们需要判断values栈顶元素和min的栈顶元素是否一样
//如果一样也需要把min的栈顶元素pop掉
//--------------------------------------
//理论上除了push操作外，其他操作都是要判断栈是否为空的
//但是我们是在做题目，所以这里先略过，以让大家理解思路为主
// */
//func Constructor() MinStack {
//	return MinStack{values: make([]int, 0), min: make([]int, 0)}
//}
//
//func (this *MinStack) Push(x int) {
//	this.values = append(this.values, x)
//	if len(this.min) == 0 || x <= this.GetMin() {
//		this.min = append(this.min, x)
//	}
//}
//
//func (this *MinStack) Pop() {
//	if this.Top() == this.GetMin() {
//		this.min = this.min[:len(this.min)-1]
//	}
//	this.values = this.values[:len(this.values)-1]
//}
//
//func (this *MinStack) Top() int {
//	return this.values[len(this.values)-1]
//}
//
//func (this *MinStack) GetMin() int {
//	return this.min[len(this.min)-1]
//}
