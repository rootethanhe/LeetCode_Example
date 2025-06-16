package main

import "fmt"

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	dp := make([]int, n)

	// 初始化第一行
	dp[0] = grid[0][0]
	for j := 1; j < n; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}

	// 动态规划主循环
	for i := 1; i < m; i++ {
		// 更新当前行首列(只能从上方来)
		dp[0] += grid[i][0]

		// 更新当前行其他列
		for j := 1; j < n; j++ {
			dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// 示例 1
	grid1 := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(grid1)) // 输出: 7 (路径: 1→3→1→1→1)

	// 示例 2
	grid2 := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(minPathSum(grid2)) // 输出: 12 (路径: 1→2→3→6)

	// 边界测试：单行单列
	grid3 := [][]int{{5}}
	fmt.Println(minPathSum(grid3)) // 输出: 5
}
