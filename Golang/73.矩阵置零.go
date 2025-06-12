package main

import "fmt"

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	firstRowZero, firstColZero := false, false

	// 1. 检查第一行/原始零值
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			firstRowZero = true
			break
		}
	}
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstColZero = true
			break
		}
	}

	// 2. 使用第一行/列记录标记
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0 // 列标记存首列
				matrix[0][j] = 0 // 行标记存首行
			}
		}
	}

	// 3. 倒序更新非首行/列
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 4. 处理首行/列
	if firstRowZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if firstColZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

func main() {
	// 示例1
	matrix1 := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix1)
	fmt.Println(matrix1) // [[1 0 1] [0 0 0] [1 0 1]]

	// 示例2
	matrix2 := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix2)
	fmt.Println(matrix2) // [[0 0 0 0] [0 4 5 0] [0 3 1 0]]

	// 边界测试：首行首列含零
	matrix3 := [][]int{{0, 2, 3}, {4, 5, 6}, {7, 0, 9}}
	setZeroes(matrix3)
	fmt.Println(matrix3) // [[0 0 0] [0 0 6] [0 0 0]]
}
