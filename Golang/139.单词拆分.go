package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	// 预处理字典：存储单词并记录最大长度
	wordSet := make(map[string]bool)
	maxLen := 0
	for _, word := range wordDict {
		wordSet[word] = true
		if len(word) > maxLen {
			maxLen = len(word)
		}
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true // 空字符串初始化为true

	for i := 1; i <= n; i++ {
		// 检查所有可能的子串长度(1到maxLen)
		start := max(0, i-maxLen) // 防止越界
		for j := start; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break // 找到一个可行解即可跳出
			}
		}
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 输入: s = "leetcode", wordDict = ["leet", "code"]
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	// 输出: true

	// 输入: s = "applepenapple", wordDict = ["apple", "pen"]
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	// 输出: true

	// 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	// 输出: false
}
