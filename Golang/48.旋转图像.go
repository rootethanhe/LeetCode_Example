package main

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix)

	// 1. 矩阵转置(只遍历上三角)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// 交换对称元素
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 2. 水平翻转每一行
	for i := 0; i < n; i++ {
		left, right := 0, n-1
		for left < right {
			// 双指针首位交换
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
			left++
			right--
		}
	}
}

func main() {
	// 示例 1
	matrix1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix1)
	fmt.Println("示例1:", matrix1) // [[7 4 1] [8 5 2] [9 6 3]]

	// 示例 2
	matrix2 := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(matrix2)
	fmt.Println("示例2:", matrix2) // [[15 13 2 5] [14 3 4 1] [12 6 8 9] [16 7 10 11]]
}
