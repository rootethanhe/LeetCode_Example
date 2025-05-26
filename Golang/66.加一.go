package main

import "fmt"

func plusOne(digits []int) []int {
	n := len(digits)

	// 从末位开始遍历
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits // 无进位直接返回
		}
		digits[i] = 0 // 进位后置零
	}

	// 处理全为9的特殊情况（如999→1000）
	return append([]int{1}, digits...)
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3})) // [1 2 4]
	fmt.Println(plusOne([]int{9, 9, 9})) // [1 0 0 0]
	fmt.Println(plusOne([]int{8, 9, 9})) // [9 0 0]
	fmt.Println(plusOne([]int{0}))       // [1]（题目保证输入非0开头，实际不会出现）
}
