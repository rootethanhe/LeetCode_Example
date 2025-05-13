package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := left + (right-left)/2 // 防止溢出
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			right = middle - 1 // 缩小至左区间
		} else {
			left = middle + 1 // 缩小至右区间
		}
	}
	return left // 未找到时返回插入位置
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5)) // 2（示例1）
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2)) // 1（示例2）
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7)) // 4（示例3）
}
