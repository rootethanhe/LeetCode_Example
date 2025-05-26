package main

import "fmt"

func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}

	// 初始化每个孩子至少1颗糖果
	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1
	}

	// 从左到右遍历，处理右侧比左侧高的情况
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// 从右到左遍历，处理左侧比右侧高的情况，并取最大值
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			// 避免覆盖左→右的结果，取左右遍历中的最大值
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	// 累加总糖果数
	total := 0
	for _, c := range candies {
		total += c
	}
	return total
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(candy([]int{1, 0, 2})) // 输出: 5
	fmt.Println(candy([]int{1, 2, 2})) // 输出: 4
}
