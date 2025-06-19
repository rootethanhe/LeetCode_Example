package main

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2 // 防溢出计算中点
		if nums[mid] > nums[right] {
			left = mid + 1 // 最小元素在右侧
		} else {
			right = mid // 最小元素在左侧(含中点)
		}
	}

	return nums[left]
}

func main() {
	// 示例1：最小元素在旋转点
	nums1 := []int{3, 4, 5, 1, 2}
	fmt.Println(findMin(nums1)) // 输出: 1

	// 示例2：最小元素在数组后半段
	nums2 := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(findMin(nums2)) // 输出: 0

	// 示例3：无旋转（完全升序）
	nums3 := []int{11, 13, 15, 17}
	fmt.Println(findMin(nums3)) // 输出: 11

	// 边界测试：单元素数组
	nums4 := []int{5}
	fmt.Println(findMin(nums4)) // 输出: 5

	// 特殊测试：最小元素为首元素（旋转后与原数组相同）
	nums5 := []int{2, 3, 4, 5, 1}
	fmt.Println(findMin(nums5)) // 输出: 1
}
