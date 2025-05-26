package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	// 创建哈希表，记录元素最后出现的索引
	lastIndex := make(map[int]int)

	for i, num := range nums {
		// 检查当前元素是否存在于哈希表中
		if j, exists := lastIndex[num]; exists {
			// 若索引差符合要求，直接返回 true
			if i-j <= k {
				return true
			}
		}
		// 更新当前元素的最后出现位置（无论是否满足条件）
		lastIndex[num] = i
	}
	return false
}

func main() {
	// 示例测试
	example1 := []int{1, 2, 3, 1}
	fmt.Println(containsNearbyDuplicate(example1, 3)) // 输出: true
	example2 := []int{1, 0, 1, 1}
	fmt.Println(containsNearbyDuplicate(example2, 1)) // 输出: true
	example3 := []int{1, 2, 3, 1, 2, 3}
	fmt.Println(containsNearbyDuplicate(example3, 2)) // 输出: false
	// 边界测试：k=0
	fmt.Println(containsNearbyDuplicate([]int{1, 1}, 0)) // 输出: false
	// 边界测试：所有元素相同
	fmt.Println(containsNearbyDuplicate([]int{5, 5, 5, 5}, 1)) // 输出: true
}
