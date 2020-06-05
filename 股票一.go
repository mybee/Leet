package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{2, 5, 6, 8}
	fmt.Println(maxProfit(a))
}

func maxProfit(prices []int) int {
	minPrice := math.MaxInt64
	maxPrice := 0
	for i:=0; i<len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i] - minPrice > maxPrice {
			maxPrice = prices[i] - minPrice
		}
	}
	return maxPrice
}
