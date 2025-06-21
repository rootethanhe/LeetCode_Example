package main

import "fmt"

func rob(nums []int) int {
	n := len(nums)
	// 边界处理
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}

	// 初始化状态
	prevPrev := nums[0]           // dp[0]
	prev := max(nums[0], nums[1]) // dp[1]

	// 状态转移(从第3间开始)
	for i := 2; i < n; i++ {
		current := max(prev, prevPrev+nums[i])
		prevPrev = prev
		prev = current
	}

	return prev
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 1}, 4},     // 示例1
		{[]int{2, 7, 9, 3, 1}, 12}, // 示例2
		{[]int{0}, 0},              // 边界：单元素
		{[]int{5, 2, 3}, 8},        // 测试：偷首尾
	}

	for _, test := range tests {
		res := rob(test.nums)
		fmt.Printf("输入: %v \t 输出: %d (期望: %d)\n", test.nums, res, test.want)
	}
}
