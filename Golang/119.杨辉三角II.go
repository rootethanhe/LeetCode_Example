package main

import "fmt"

// 组合数公式法
func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	res[0] = 1

	for i := 1; i <= rowIndex; i++ {
		// 递推公式: 当前项 = 前一项 * (行号 - 当前索引 + 1) / 当前索引
		res[i] = res[i-1] * (rowIndex - i + 1) / i
	}

	return res
}

/*
动态规划(倒序更新)
func getRow(rowIndex int) []int {
    res := make([]int, rowIndex+1)
    res[0] = 1
    for i := 1; i <= rowIndex; i++ {
        // 倒序更新防止覆盖前一行数据
        for j := i; j > 0; j-- {
            res[j] += res[j-1]  // DP状态转移方程
        }
    }
    return res
}
*/

func main() {
	// 输入: rowIndex = 3  输出: [1,3,3,1]
	fmt.Println(getRow(3))

	// 输入: rowIndex = 0  输出: [1]
	fmt.Println(getRow(0))

	// 输入: rowIndex = 1  输出: [1,1]
	fmt.Println(getRow(1))
}
