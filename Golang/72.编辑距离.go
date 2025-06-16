package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 初始化边界条件
	for i := 0; i <= m; i++ {
		dp[i][0] = i // 删除操作积累
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j // 插入操作积累
	}

	// 填充DP表
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] // 字符匹配, 无需操作
			} else {
				insert := dp[i][j-1] + 1    // 插入操作代价
				delete := dp[i-1][j] + 1    // 删除操作代价
				replace := dp[i-1][j-1] + 1 // 替换操作代价
				dp[i][j] = min(insert, delete, replace)
			}
		}
	}
	return dp[m][n]
}

func min(a, b, c int) int {
	m := a
	if b < m {
		m = b
	}
	if c < m {
		m = c
	}
	return m
}

func main() {
	// 示例测试
	fmt.Println(minDistance("horse", "ros"))           // 输出: 3
	fmt.Println(minDistance("intention", "execution")) // 输出: 5
}
