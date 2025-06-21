package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	curSum := nums[0] // 当前子序和(以当前元素结尾)
	maxSum := nums[0] // 全局最大和(处理全负数情况)

	// 从第二个元素开始遍历
	for i := 1; i < len(nums); i++ {
		// 若当前子序和为负，重置当前元素(负数拖累后续和)
		if curSum <= 0 {
			curSum = nums[i]
		} else {
			curSum += nums[i]
		}

		// 更新全局最大和
		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
}

func main() {
	// 示例1: 输出6 (子数组 [4,-1,2,1])
	nums1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums1))

	// 示例2: 输出1 (子数组 [1])
	nums2 := []int{1}
	fmt.Println(maxSubArray(nums2))

	// 示例3: 输出23 (子数组 [5,4,-1,7,8])
	nums3 := []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubArray(nums3))

	// 边界测试: 全负数数组，输出-1
	nums4 := []int{-3, -1, -5, -2}
	fmt.Println(maxSubArray(nums4))
}
