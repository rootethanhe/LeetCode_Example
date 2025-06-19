package main

import "fmt"

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2 // 防溢出计算中点
		if nums[mid] < nums[mid+1] {
			left = mid + 1 // 右侧爬坡
		} else {
			right = mid // 左侧爬坡
		}
	}

	return left // left == right 时为峰值位置
}

func main() {
	// 示例1：峰值在中间
	nums1 := []int{1, 2, 3, 1}
	fmt.Println(findPeakElement(nums1)) // 输出: 2 (nums[2]=3是峰值)

	// 示例2：多个峰值（返回任意一个）
	nums2 := []int{1, 2, 1, 3, 5, 6, 4}
	fmt.Println(findPeakElement(nums2)) // 可能输出: 1 (nums[1]=2) 或 5 (nums[5]=6)

	// 边界测试：峰值在首尾
	nums3 := []int{5, 4, 3, 2, 1}       // 峰值在开头 (nums[0]=5 > nums[1])
	fmt.Println(findPeakElement(nums3)) // 输出: 0

	nums4 := []int{1, 2, 3, 4, 5}       // 峰值在结尾 (nums[4]=5 > nums[3])
	fmt.Println(findPeakElement(nums4)) // 输出: 4
}
