package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	// 1. 按区间起始点排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 2. 初始化结果数组
	merged := [][]int{intervals[0]}

	// 3. 遍历合并结果区间
	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		current := intervals[i]
		if current[0] <= last[1] { // 存在重叠
			if current[1] > last[1] { // 更新右边界为较大值
				merged[len(merged)-1][1] = current[1]
			}
		} else {
			merged = append(merged, current)
		}
	}

	return merged
}

func main() {
	// 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]  输出：[[1,6],[8,10],[15,18]]
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))

	// 输入：intervals = [[1,4],[4,5]]  输出：[[1,5]]
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
}
