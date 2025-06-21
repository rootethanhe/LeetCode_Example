package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	// 初始化四个状态
	buy1 := -prices[0]
	sell1 := 0
	buy2 := -prices[0]
	sell2 := 0

	// 从第二天开始遍历
	for i := 1; i < len(prices); i++ {
		// 更新第一次交易状态
		buy1 = max(buy1, -prices[i])       // 买入或不操作
		sell1 = max(sell1, buy1+prices[i]) // 卖出或不操作

		// 更新第二次交易状态
		buy2 = max(buy2, sell1-prices[i])  // 用第一次利润买入
		sell2 = max(sell2, buy2+prices[i]) // 第二次卖出
	}

	return sell2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 示例1：输出6（交易：[4-6天], [7-8天]）
	prices1 := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println(maxProfit(prices1))

	// 示例2：输出4（交易：[1-5天]）
	prices2 := []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit(prices2))

	// 示例3：输出0（无交易）
	prices3 := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices3))

	// 示例4：输出0（单元素数组）
	prices4 := []int{1}
	fmt.Println(maxProfit(prices4))
}
