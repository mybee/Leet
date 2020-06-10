package main

import (
	"fmt"
)

func quickSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		fmt.Println(key, (start+end)/2, start, end)
		fmt.Println(i, j)
		for i <= j {
			// 注意: 这里是个 for, 不是 if
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			fmt.Println("----", i, j)
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
				fmt.Println(arr)
			}

		}

		if start < j {
			quickSort(arr, start, j)
		}
		if end > i {
			quickSort(arr, i, end)
		}
	}
}

// 时间复杂度:O(nlogn)

// 最坏情况: O(n2)

// 空间复杂度: O(logn)

func main() {
	arr := []int{80, 2, 5, 7, 9 ,8}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}