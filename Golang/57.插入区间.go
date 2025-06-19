package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	i := 0
	n := len(intervals)

	// 1. 添加所有在新区间之前的区间(不重叠)
	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	// 2. 合并所有与新区间重叠的区间
	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(intervals[i][0], newInterval[0])
		newInterval[1] = max(intervals[i][1], newInterval[1])
		i++
	}
	result = append(result, newInterval) // 添加合并后的区间

	// 3. 添加剩余的区间(在新区间之后)
	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 示例1
	intervals1 := [][]int{{1, 3}, {6, 9}}
	newInterval1 := []int{2, 5}
	fmt.Println(insert(intervals1, newInterval1)) // [[1 5] [6 9]]

	// 示例2
	intervals2 := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval2 := []int{4, 8}
	fmt.Println(insert(intervals2, newInterval2)) // [[1 2] [3 10] [12 16]]

	// 空列表测试
	fmt.Println(insert([][]int{}, []int{5, 7})) // [[5 7]]

	// 完全包含测试
	fmt.Println(insert([][]int{{1, 5}}, []int{2, 3})) // [[1 5]]

	// 右扩展测试
	fmt.Println(insert([][]int{{1, 5}}, []int{2, 7})) // [[1 7]]
}
