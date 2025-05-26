package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	// 计算前缀积（左侧乘积）
	answer[0] = 1
	for i := 1; i < n; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	// 计算后缀积（右侧乘积）并更新答案
	right := 1
	for i := n - 2; i >= 0; i-- {
		right *= nums[i+1] // 累积右侧乘积
		answer[i] *= right // 左乘积 * 右乘积
	}

	return answer
}

func main() {
	nums1 := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums1)) // 输出: [24 12 8 6]

	nums2 := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums2)) // 输出: [0 0 9 0 0]
}
