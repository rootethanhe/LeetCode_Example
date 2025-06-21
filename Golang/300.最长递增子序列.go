package main

import (
	"fmt"
	"sort"
)

func lengthOfLIS_DP(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n) // dp[i]: 以 nums[i] 结尾的 LTS 长度
	maxLen := 1

	for i := 0; i < n; i++ {
		dp[i] = 1 // 初始化: 每个元素自身是长度为1的子序列
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1 // 状态转移
				}
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i] // 更新全局最大值
		}
	}
	return maxLen
}

func lengthOfLIS_Greedy(nums []int) int {
	d := make([]int, 0) // d[k]: 长度为 k 的 LIS 的最小末尾值
	for _, num := range nums {
		// 若 num 大于所有 d 中的值，则扩展序列
		if len(d) == 0 || num > d[len(d)-1] {
			d = append(d, num)
		} else {
			// 二分查找第一个 ≥ num 的位置
			pos := sort.SearchInts(d, num)
			d[pos] = num // 替换 d[pos]，保持最小末尾值
		}
	}
	return len(d) // d 的长度即 LIS 长度
}

func main() {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4}, // 示例1: [2,3,7,101]
		{[]int{0, 1, 0, 3, 2, 3}, 4},           // 示例2: [0,1,2,3]
		{[]int{7, 7, 7, 7, 7, 7, 7}, 1},        // 示例3: [7]
	}
	for _, test := range tests {
		res1 := lengthOfLIS_DP(test.nums)
		res2 := lengthOfLIS_Greedy(test.nums)
		fmt.Printf("输入: %v\n动态规划: %d\n贪心+二分: %d (期望: %d)\n\n",
			test.nums, res1, res2, test.want)
	}
}
