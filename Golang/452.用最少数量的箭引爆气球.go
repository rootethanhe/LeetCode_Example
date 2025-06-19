package main

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	// 按每个气球的结束坐标 xend 升序排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	arrowCount := 1         // 至少需要一支箭
	prevEnd := points[0][1] // 第一支箭的位置设为第一个气球的结束位置

	for i := 1; i < len(points); i++ {
		// 当前气球的开始位置 > 上一支箭的位置，说明无重叠
		if points[i][0] > prevEnd {
			arrowCount++           // 新增一支箭
			prevEnd = points[i][1] // 箭设在当前气球的结束位置
		}
		// 否则（points[i][0] <= prevEnd）：当前气球被上一支箭覆盖，无需操作
	}
	return arrowCount
}

func main() {
	examples := [][][]int{
		{{10, 16}, {2, 8}, {1, 6}, {7, 12}}, // 输出 2
		{{1, 2}, {3, 4}, {5, 6}, {7, 8}},    // 输出 4
		{{1, 2}, {2, 3}, {3, 4}, {4, 5}},    // 输出 2
	}
	for _, points := range examples {
		fmt.Println(findMinArrowShots(points))
	}
}
