package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print(time.Now().Format(time.RFC3339))
}



//这道题不太好懂的是题意。
//
//1、每一个位置都有 2 个阶梯，1 个阶梯上一层楼，另 1 个阶梯上两层楼；
//
//2、上两个阶梯的体力值耗费是一样的，但是在不同位置消耗的体力值是不一样的；
//
//3、楼层顶部在数组之外。如果数组长度为 len，那么楼顶就在下标为 len，注意 dp 数组开 len + 1 个空间。
//
//状态：dp[i] 表示到索引为 i 位置再选择向上爬一共需要的体力开销。
//
//状态转移方程：
//
//dp[i] = min(dp[i - 1], dp[i - 2]) + cost[i]
//输出： dp[len]。
//
//作者：liweiwei1419
//链接：https://leetcode-cn.com/problems/min-cost-climbing-stairs/solution/dong-tai-gui-hua-by-liweiwei1419-3/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
