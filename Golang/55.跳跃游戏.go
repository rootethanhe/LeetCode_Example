package main

import "fmt"

func canJump(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true // 单元素数组无需跳跃
	}

	maxReach := 0 // 当前能到达的最远位置
	for i := 0; i < n; i++ {
		if i > maxReach { // 当前位置不可达
			return false
		}
		maxReach = max(maxReach, i+nums[i]) // 更新最远覆盖范围
		if maxReach >= n-1 {                // 已覆盖终点
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4})) // true
	fmt.Println(canJump([]int{3, 2, 1, 0, 4})) // false
}
