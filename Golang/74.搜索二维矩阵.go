package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	// 空矩阵检查
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows, cols := len(matrix), len(matrix[0])
	left, right := 0, rows*cols-1 // 初始化一维索引边界

	for left <= right {
		mid := left + (right-left)/2 // 防溢出计算中点
		row := mid / cols            // 行索引 = 一维索引/列数
		col := mid % cols            // 列索引 = 一维索引%列数

		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			left = mid + 1 // 目标在右半区
		} else {
			right = mid - 1 // 目标在左半区
		}
	}

	return false
}

func main() {
	// 示例1：存在目标值
	matrix1 := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	fmt.Println(searchMatrix(matrix1, 3)) // true

	// 示例2：不存在目标值
	fmt.Println(searchMatrix(matrix1, 13)) // false

	// 单行矩阵
	matrix2 := [][]int{{1, 3, 5, 7}}
	fmt.Println(searchMatrix(matrix2, 5)) // true
	fmt.Println(searchMatrix(matrix2, 2)) // false

	// 单元素矩阵
	matrix3 := [][]int{{5}}
	fmt.Println(searchMatrix(matrix3, 5)) // true
	fmt.Println(searchMatrix(matrix3, 0)) // false

	// 包含重复值
	matrix4 := [][]int{
		{1, 3, 5, 7},
		{7, 8, 9, 10},
	}
	fmt.Println(searchMatrix(matrix4, 7)) // true
}
