package main

import "fmt"

// 输入整数数组 arr ，找出其中最小的 k 个数。例如，
// 输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

func main() {
	// 前k小, 用大根堆
	// 前k大, 用小根堆
	fmt.Println(getLeastNumbers([]int{1, 1, 1, 2, 4,5,3,2}, 3))

}

// 堆化 函数
func heapify(root int, length int, arr []int) {
	fmt.Println(root, length, arr)
	for root < length {
		fmt.Println(arr)
		//2, 5, 6
		min, left, right := root, 2*root+1, 2*root+2
		fmt.Println(min, left, right)
		if left < length && arr[left] < arr[min] {
			min = left
		}
		if right < length && arr[right] < arr[min] {
			min = right
		}
		if min != root {
			// 如果 不符合子孩子 和 父的关系, 则调换位置
			arr[min], arr[root] = arr[root], arr[min]
			root = min
		}else{
			break
		}
	}
}

func getLeastNumbers(arr []int, k int) []int {
	// 不用这部分
	//for i := len(arr) / 2; i >= 0; i-- {
	//	heapify(i, len(arr), arr)
	//}

	// 如有疑问:
	// https://www.cnblogs.com/chengxiao/p/6129630.html  这个很详细
	// 自下而上下滤
	// 思路
	// 1. 将最大的过滤到顶部
	// 2. 将顶部的最后一个(每次减一)交换
	for i := len(arr) - 1; i >= len(arr)-k; i-- {
		// 将 浮到最上面的 root当前最大值 和 后面的 交换
		arr[0], arr[i] = arr[i], arr[0]
		// 将最大值浮到root
		heapify(0, i, arr)
	}
	fmt.Println(arr)
	return arr[len(arr)-k:]
}




