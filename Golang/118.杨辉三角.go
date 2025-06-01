package main

import "fmt"

func generate(numRows int) [][]int {
	// 处理边界情况
	if numRows == 0 {
		return [][]int{}
	}

	// 初始化结果二维切片
	triangle := make([][]int, numRows)

	for i := range triangle {
		// 每行有i+1个元素
		triangle[i] = make([]int, i+1)
		// 收尾元素设为1
		triangle[i][0], triangle[i][i] = 1, 1

		// 计算中间元素(j从1到i-1)
		for j := 1; j < i; j++ {
			// 核心递推公式
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	return triangle
}

func main() {
	// 输入: numRows = 5
	// 输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
	fmt.Println(generate(5))

	// 输入: numRows = 1
	// 输出: [[1]]
	fmt.Println(generate(1))
}
