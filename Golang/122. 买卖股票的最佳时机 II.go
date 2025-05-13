package main

import "fmt"

func maxProfit(prices []int) int {
	/* 1、贪心算法
	profit := 0
	for i := 1; i < len(prices); i++ {
		if diff := prices[i] - prices[i-1]; diff > 0 {
			profit += diff // 累加所有相邻正利润
		}
	}

	return profit
	*/

	// 2、动态规划
	if len(prices) < 2 {
		return 0
	}

	hold := -prices[0] // 初始持有状态（第一天买入）
	cash := 0          // 初始不持有状态

	for i := 1; i < len(prices); i++ {
		prevCash := cash
		cash = max(cash, hold+prices[i])     // 卖出操作
		hold = max(hold, prevCash-prices[i]) // 买入操作（需用前一天的现金）
	}
	return cash
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // 7（贪心/动态规划均正确）
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))    // 4（连续上涨）
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))    // 0（持续下跌）
}
