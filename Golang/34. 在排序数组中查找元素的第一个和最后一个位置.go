package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	// 查找左边界
	left := findLeft(nums, target)
	if left == -1 {
		return []int{-1, -1}
	}
	// 查找右边界
	right := findRight(nums, target)
	return []int{left, right}
}

// findLeft 查找左边界(第一个等于target的位置)
func findLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2 // 避免溢出
		if nums[mid] >= target {     // 关键：等于时继续向左收缩
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	// 检查是否越界或值不匹配
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

// findRight 查找右边界(最后一个等于target的位置)
func findRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target { // 关键：等于时继续向右收缩
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// 检查是否越界或值不匹配
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}

func main() {
	// 示例 1
	nums1 := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(nums1, 8)) // 输出: [3 4]

	// 示例 2
	nums2 := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(nums2, 6)) // 输出: [-1 -1]

	// 示例 3
	nums3 := []int{}
	fmt.Println(searchRange(nums3, 0)) // 输出: [-1 -1]
}
