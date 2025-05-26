package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	minPrice := math.MaxInt32 // 初始化最小价格为一个极大值
	maxProfit := 0            // 初始化最大利润为0

	for _, price := range prices {
		if price < minPrice {
			// 更新历史最低价
			minPrice = price
		} else if profit := price - minPrice; profit > maxProfit {
			// 计算当前价格与历史最低价的利润，更新最大利润
			maxProfit = profit
		}
	}
	return maxProfit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // 输出:5（符合示例1）
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))    // 输出:0（符合示例2）
	fmt.Println(maxProfit([]int{2, 4, 1}))          // 输出:2（最低价1，最高价4）
	fmt.Println(maxProfit([]int{3}))                // 输出:0（无法交易）
}
