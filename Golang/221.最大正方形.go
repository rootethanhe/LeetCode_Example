package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, cols := len(matrix), len(matrix[0])
	dp := make([]int, cols+1) // 多一列避免边界检查
	maxSide, prev := 0, 0     // prev 保存左上角状态

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			temp := dp[j+1] // 保存当前状态供下一轮使用
			if matrix[i][j] == '0' {
				dp[j+1] = 0
			} else {
				// 状态转移：min(左, 上, 左上) + 1
				if i == 0 || j == 0 {
					dp[j+1] = 1
				} else {
					dp[j+1] = min(min(dp[j], dp[j+1]), prev) + 1
				}
				// 更新最大边长
				if dp[j+1] > maxSide {
					maxSide = dp[j+1]
				}
			}
			prev = temp // 更新左上角状态
		}
	}
	return maxSide * maxSide
}

func main() {
	// 示例 1
	matrix1 := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalSquare(matrix1)) // 输出: 4

	// 示例 2
	matrix2 := [][]byte{
		{'0', '1'},
		{'1', '0'},
	}
	fmt.Println(maximalSquare(matrix2)) // 输出: 1

	// 示例 3
	matrix3 := [][]byte{{'0'}}
	fmt.Println(maximalSquare(matrix3)) // 输出: 0

	// 全1矩阵
	matrix4 := [][]byte{
		{'1', '1', '1'},
		{'1', '1', '1'},
		{'1', '1', '1'},
	}
	fmt.Println(maximalSquare(matrix4)) // 输出: 9

	// 单行矩阵
	matrix5 := [][]byte{{'1', '1', '1', '1'}}
	fmt.Println(maximalSquare(matrix5)) // 输出: 1
}
