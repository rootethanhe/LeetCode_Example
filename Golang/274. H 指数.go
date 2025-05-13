package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	n := len(citations)

	// 降序排序
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	maxH := 0
	for i := 0; i < n; i++ {
		if citations[i] >= i+1 {
			maxH = i + 1
		} else {
			break // 后续元素更小，无需继续判断
		}
	}
	return maxH
}

func main() {
	// 示例1
	citations1 := []int{3, 0, 6, 1, 5}
	fmt.Println(hIndex(citations1)) // 输出: 3

	// 示例2
	citations2 := []int{1, 3, 1}
	fmt.Println(hIndex(citations2)) // 输出: 1
}
