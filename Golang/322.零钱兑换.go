package main

import "fmt"

func coinChange(coins []int, amount int) int {
	// 边界处理: 金额为0时无需硬币
	if amount == 0 {
		return 0
	}

	// 初始化dp数组, dp[i]表示凑成金额i所需的最少硬币数
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1 // 初始化为不可达值
	}
	dp[0] = 0 // 金额0不需要硬币

	// 动态规划填表
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			// 若硬币面值不超过当前金额, 且减去该面值后的状态可达
			if coin <= i && dp[i-coin] != amount+1 {
				// 更新dp[i]为更优解
				if dp[i-coin]+1 < dp[i] {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}

	// 返回结果
	if dp[amount] > amount {
		return -1 // 无法凑成
	}
	return dp[amount]
}

func main() {
	tests := []struct {
		coins  []int
		amount int
		want   int
	}{
		{[]int{1, 2, 5}, 11, 3},       // 11 = 5+5+1
		{[]int{2}, 3, -1},             // 无法凑出3
		{[]int{1}, 0, 0},              // 金额0
		{[]int{25, 21, 10, 1}, 63, 3}, // 最优解：21 * 3（贪心会失效）
	}

	for _, test := range tests {
		res := coinChange(test.coins, test.amount)
		fmt.Printf("coins=%v\tamount=%d\t→ %d (期望: %d)\n",
			test.coins, test.amount, res, test.want)
	}
}
