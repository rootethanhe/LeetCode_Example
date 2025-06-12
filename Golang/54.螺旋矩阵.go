package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	// 初始化边界
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	res := []int{}

	for top <= bottom && left <= right {
		// 1. 从左到右遍历上边界
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++ // 上边界下移

		// 2. 从上到下遍历右边界
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right-- // 右边界左移

		// 检查下边界是否有效
		if top <= bottom {
			// 3. 从右到左遍历下边界
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom-- // 下边界上移
		}

		// 检查左边界是否有效
		if left <= right {
			// 4. 从下到上遍历左边界
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++ // 左边界右移
		}
	}

	return res
}

func main() {
	// 示例 1
	matrix1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("示例1:", spiralOrder(matrix1)) // [1 2 3 6 9 8 7 4 5]

	// 示例 2
	matrix2 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println("示例2:", spiralOrder(matrix2)) // [1 2 3 4 8 12 11 10 9 5 6 7]

	// 测试单行矩阵
	matrix3 := [][]int{{5, 1, 9, 11}}
	fmt.Println("单行矩阵:", spiralOrder(matrix3)) // [5 1 9 11]

	// 测试单列矩阵
	matrix4 := [][]int{{2}, {4}, {6}}
	fmt.Println("单列矩阵:", spiralOrder(matrix4)) // [2 4 6]
}
