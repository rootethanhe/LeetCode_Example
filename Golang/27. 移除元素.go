package main

import "fmt"

func removeElement(nums []int, val int) int {
	slow := 0                                 // 慢指针标记有效位置
	for fast := 0; fast < len(nums); fast++ { // 快指针遍历所有元素
		if nums[fast] != val { // 发现有效元素
			nums[slow] = nums[fast] // 覆盖到慢指针位置
			slow++                  // 有效位置后移
		}
	}
	return slow // 返回有效元素数量
}

func main() {
	// 测试用例
	testCases := []struct {
		nums []int
		val  int
		want int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},             // 示例1
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5}, // 示例2
		{[]int{}, 1, 0},                       // 空数组
		{[]int{2, 2, 2}, 2, 0},                // 全需删除
	}

	for _, tc := range testCases {
		// 创建数组副本用于测试
		nums := append([]int{}, tc.nums...)

		// 测试双指针法
		k1 := removeElement(nums, tc.val)
		fmt.Printf("双指针法：输入%v => 返回%d, 数组前%d位:%v\n",
			tc.nums, k1, k1, nums[:k1])
	}
}
