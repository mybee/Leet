package main

//二分查找的常见痛点如下：
//
//1. 死循环问题
//2.循环结束条件到底是哪个：start<=end ？ start<end ? start+1<end
//3.指针变化到底是哪个: start=mid+1 ? start=mid-1 ? start=mid-1
//在此直接给出通用模板：模板的核心就是
//
//1、left+1<right (作为循环条件【保证了永远不出现死循环】)
//2、循环结束要单独判断left和right的情况
//3、中间不需要指针的变化不需要什么mid+1,mid-1这些，直接等于mid就行
//
//作者：shixuewei
//链接：https://leetcode-cn.com/problems/find-in-mountain-array/solution/fen-xiang-yi-tao-zi-ji-xi-huan-de-er-fen-cha-zhao-/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

//二分查找通用模板
func binarysearch(arr []int,left,right int,target int)int{
	//循环结束条件
	for left+1<right{ //1、条件一
		//中数(这样写为了防止left+right太大导致溢出)
		mid:=left+(right-left)/2
		//相等时：
		if arr[mid]==target{
			//执行相应操作
			return mid
		}else if arr[mid] > target{ //大于时
			right=mid  //3、条件三
		}else{//小于时
			left=mid
		}
	}
	//2、条件二
	//循环结束后单独判断left和right就行
	if arr[left]==target{
		//执行操作
		return left
	}
	if arr[right]==target{
		return right
	}
	return -1 //找不到
}

// 二分法 题目:
// 69,367,35,34(重点),74,240,153,154,162,33,81

