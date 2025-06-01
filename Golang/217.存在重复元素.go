package main

import "fmt"

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool) // 创建哈希表记录元素出现状态

	for _, num := range nums { // 遍历数组
		if seen[num] { // 若元素已存在
			return true // 立即返回true
		}
		seen[num] = true // 标记当前元素已出现
	}
	return false // 遍历结束无重复
}

func main() {
	// 输入：nums = [1,2,3,1]  输出：true
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))

	// 输入：nums = [1,2,3,4]  输出：false
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))

	// 输入：nums = [1,1,1,3,3,4,3,2,4,2]  输出：true
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
