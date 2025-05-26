package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s // 边界条件直接返回
	}

	// 初始化行存储数组
	rows := make([][]byte, numRows)
	row, direction := 0, 1 // 初始方向向下(行号递增)

	for i := 0; i < len(s); i++ {
		rows[row] = append(rows[row], s[i])
		// 方向反转条件(到达顶部或底部)
		if row == 0 {
			direction = 1
		} else if row == numRows-1 {
			direction = -1
		}
		row += direction
	}

	// 合并所有行的字符
	result := make([]byte, 0)
	for _, r := range rows {
		result = append(result, r...)
	}

	return string(result)
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3)) // "PAHNAPLSIIGYIR"
	fmt.Println(convert("PAYPALISHIRING", 4)) // "PINALSIGYAHRPI"
	fmt.Println(convert("A", 1))              // "A"
}
