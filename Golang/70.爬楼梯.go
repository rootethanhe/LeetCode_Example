package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	// 初始化前两个状态: 1阶和2阶的方法数
	prev, current := 1, 2
	for i := 3; i <= n; i++ {
		// 计算当前状态，并更新前两个状态
		prev, current = current, prev+current
	}
	return current
}

func main() {
	// 输入：n = 2  输出：2
	fmt.Println(climbStairs(2))

	// 输入：n = 3  输出：3
	fmt.Println(climbStairs(3))
}
