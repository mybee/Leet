package main

import "sort"

// 可转为两数之和问题

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	nSum(nums, target, 0, 4, &res, []int{})
	return res
}

func nSum(nums []int, target int, pos int, n int, res *[][]int, cur []int) {
	if n == 2 { // 只有两个
		left, right := pos, len(nums)-1
		for left < right { // 终止条件
			if sum := nums[left] + nums[right]; sum < target { // 偏小
				left++ // 移动左边
			} else if sum > target { // 偏大
				right-- // 移动右边
			} else { // 相等
				t := make([]int, len(cur)+2)
				copy(t, cur) // 复制游标
				t[len(t)-2] = nums[left]
				t[len(t)-1] = nums[right]
				*res = append(*res, t) // 保存游标
				left++
				right--
				for left < right && nums[left] == nums[left-1] { // 跳过相同数字
					left++
				}
				for left < right && nums[right] == nums[right+1] { // 跳过相同数字
					right--
				}
			}
		}
		return
	}
	for i := pos; i < len(nums)-n+1; i++ {
		if target < nums[i]*n || target > nums[len(nums)-1]*n { // 目标数字在当前有序数组边界之外
			break
		}
		if i > pos && nums[i] == nums[i-1] { // 跳过相同数字
			continue
		}
		cur = append(cur, nums[i])                     // 增加游标
		nSum(nums, target-nums[i], i+1, n-1, res, cur) // 递归
		cur = cur[:len(cur)-1]                         // 移除游标
	}
}

