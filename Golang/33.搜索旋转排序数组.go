package main

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2 // 防止溢出
		if nums[mid] == target {
			return mid
		}
		// 判断左半部分是否有序
		if nums[left] <= nums[mid] {
			// 检查目标值是否在左半部分的范围内
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右半部分有序
			// 检查目标值是否在右半部分的范围内
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {
	// 输入：nums = [4,5,6,7,0,1,2], target = 0  输出：4
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))

	// 输入：nums = [4,5,6,7,0,1,2], target = 3  输出：-1
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))

	// 输入：nums = [1], target = 0  输出：-1
	fmt.Println(search([]int{1}, 0))
}
