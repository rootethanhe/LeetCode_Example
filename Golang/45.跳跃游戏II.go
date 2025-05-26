package main

import "fmt"

func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	jumps := 0      // 记录跳跃次数
	currentEnd := 0 // 当前跳跃的结束位置
	maxFar := 0     // 全局能到达的最远位置

	for i := 0; i < n; i++ {
		// 更新全局最远可达位置
		if i+nums[i] > maxFar {
			maxFar = i + nums[i]
		}
		// 到达当前跳跃边界时触发跳跃
		if i == currentEnd {
			jumps++
			currentEnd = maxFar
			// 若当前边界已覆盖终点，直接终止
			if currentEnd >= n-1 {
				break
			}
		}
	}
	return jumps
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4})) // 输出: 2
	fmt.Println(jump([]int{2, 3, 0, 1, 4})) // 输出: 2
	fmt.Println(jump([]int{5}))             // 输出: 0
	fmt.Println(jump([]int{1, 1, 1, 1, 1})) // 输出: 4
}
