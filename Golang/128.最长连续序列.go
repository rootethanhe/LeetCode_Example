package main

import "fmt"

func longestConsecutive(nums []int) int {
	// 边界处理：空数组直接返回0
	if len(nums) == 0 {
		return 0
	}

	// 初始化哈希集合
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	maxLength := 0
	// 遍历哈希集合中的每个数字
	for num := range numSet {
		// 检查是否为序列起点(前驱不存在)
		if !numSet[num-1] {
			currentNum := num
			currentLength := 1

			// 向后扩展序列
			for numSet[currentNum+1] {
				currentNum++
				currentLength++
			}

			// 更新最大长度
			if currentLength > maxLength {
				maxLength = currentLength
			}
		}
	}

	return maxLength
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))         // 4
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})) // 9
	fmt.Println(longestConsecutive([]int{1, 0, 1, 2}))                   // 3
}
