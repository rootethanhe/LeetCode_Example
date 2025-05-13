package main

import "fmt"

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	slow := 1 // 慢指针从第二个位置开始（第一个元素无需处理）
	for fast := 1; fast < n; fast++ {
		// 当发现新元素时，更新慢指针位置
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func main() {
	// 示例1
	nums1 := []int{1, 1, 2}
	k1 := removeDuplicates(nums1)
	fmt.Println("输出:", k1, nums1[:k1]) // 输出: 2 [1 2]

	// 示例2
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k2 := removeDuplicates(nums2)
	fmt.Println("输出:", k2, nums2[:k2]) // 输出: 5 [0 1 2 3 4]
}
