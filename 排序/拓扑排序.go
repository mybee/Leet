package main

import "fmt"

func main() {
	// 使用条件:
	// 连通图
	// 无环
	// 1. 统计入度表, 入度为0的顶点

	// 你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。
	//
	//在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]
	//
	//给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？
	//
	//链接：https://leetcode-cn.com/problems/course-schedule

	fmt.Println(canFinish(3, [][]int{{1,0}, {1,2}, {0, 2}}))

}

// 出度和入度 逆序而已.

func canFinish(numCourses int, prerequisites [][]int) bool {
	//先找到初始状态下所有顶点的出度
	indegree := make([]int, numCourses, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		indegree[prerequisites[i][0]] += 1  // 算出度
		//indegree[prerequisites[i][1]] += 1  // 算入度
	}
	// 输出入度表
	fmt.Println(indegree)
	queue := []int{}
	for i, v := range indegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	if len(queue) == 0 {
		return false
	}
	res := []int{}
	// 循环
	for len(queue) > 0 {
		fmt.Println("queue", queue)
		cur := queue[0]
		queue = queue[1:]
		res = append(res, cur)
		//找到该顶点连接的顶点，将其出度减1，如果减到0了，就入队
		for i := 0; i < len(prerequisites); i++ {
			if prerequisites[i][1] == cur {
				indegree[prerequisites[i][0]] -= 1
				if indegree[prerequisites[i][0]] == 0 {
					queue = append(queue, prerequisites[i][0])
				}
			}
		}

	}
	return len(res) == numCourses
}