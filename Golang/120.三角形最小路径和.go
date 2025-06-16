package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}

	// 初始化dp为最后一行
	dp := make([]int, n)
	for j := 0; j < n; j++ {
		dp[j] = triangle[n-1][j]
	}

	// 自底向上遍历
	for i := n - 2; i >= 0; i-- {
		// 更新当前行所有列
		for j := 0; j <= i; j++ {
			// 取下方两个位置的较小值 + 当前值
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}

	return dp[0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	// 示例1：路径2→3→5→1
	tri1 := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(tri1)) // 输出: 11

	// 示例2：单元素三角形
	tri2 := [][]int{{-10}}
	fmt.Println(minimumTotal(tri2)) // 输出: -10

	// 测试：全正三角形
	tri3 := [][]int{{1}, {2, 3}, {4, 5, 6}}
	fmt.Println(minimumTotal(tri3)) // 输出: 7 (1+2+4)
}
