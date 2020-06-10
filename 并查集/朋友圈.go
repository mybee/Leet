package main

import "fmt"

func main() {
	fs := [][]int{[]int{1, 1, 0},
		[]int{0, 1, 0},
		[]int{1, 0, 1}}
	fmt.Print(findCircleNum(fs))
}

func findCircleNum(M [][]int) int {
	unions := 0
	number := len(M)
	stu := make([]int, number)
	tempInit := new(inits)
	tempInit.unions = unions
	tempInit.stu = stu
	for i := 0; i < number; i++ {
		stu[i] = i //类似 武当、峨眉、少林 123
	}
	for i := 0; i < number; i++ {
		for j := i + 1; j < number; j++ {
			if M[i][j] == 1 {
				tempInit.union(i, j) //看是不是一个阵营的 eg：武当+峨眉都为剑修
			}
		}
	}
	return number - tempInit.unions
}

type inits struct {
	unions int
	stu    []int
}

func (this *inits) find(i int) int {
	for true {
		if i == this.stu[i] {
			break
		}
		this.stu[i] = this.stu[this.stu[i]]
		i = this.stu[i]
	}
	return i
}

func (this *inits) union(p, q int) {
	i := this.find(p)
	j := this.find(q)
	if i == j {
		return
	}
	// 合并
	this.stu[i] = j
	this.unions++
}
