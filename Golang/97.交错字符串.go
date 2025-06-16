package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}

	dp := make([]bool, n+1)
	dp[0] = true

	// 初始化首行(仅用s2)
	for j := 1; j <= n; j++ {
		dp[j] = dp[j-1] && (s2[j-1] == s3[j-1])
	}

	// 动态规划更新
	for i := 1; i <= m; i++ {
		// 更新当前行首列(仅用s1)
		dp[0] = dp[0] && (s1[i-1] == s3[i-1])

		for j := 1; j <= n; j++ {
			// 状态转移: 用s1匹配 或 用s2匹配
			useS1 := dp[j] && (s1[i-1] == s3[i+j-1])
			useS2 := dp[j-1] && (s2[j-1] == s3[i+j-1])
			dp[j] = useS1 || useS2
		}
	}

	return dp[n]
}

func main() {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac")) // true (示例1)
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbbaccc")) // false (示例2)
	fmt.Println(isInterleave("", "", ""))                     // true (示例3)
	fmt.Println(isInterleave("a", "b", "ab"))                 // true
	fmt.Println(isInterleave("a", "b", "ba"))                 // true
	fmt.Println(isInterleave("a", "", "a"))                   // true
}
