package main

import "fmt"

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	count := 0

	// 定义DFS淹没函数
	var dfs func(int, int)
	dfs = func(i, j int) {
		// 边界检查: 越界或遇到水则返回
		if i < 0 || j < 0 || i >= rows || j >= cols || grid[i][j] == '0' {
			return
		}

		// 淹没当前陆地
		grid[i][j] = '0'

		// 递归淹没四个方向
		dfs(i+1, j) // 下
		dfs(i-1, j) // 上
		dfs(i, j+1) // 右
		dfs(i, j-1) // 左
	}

	// 主循环: 遍历网格
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				count++   // 发现新岛屿
				dfs(i, j) // 淹没整个岛屿
			}
		}
	}

	return count
}

func main() {
	// 示例1：1个岛屿
	grid1 := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid1)) // 输出: 1

	// 示例2：3个岛屿
	grid2 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(grid2)) // 输出: 3

	// 边界测试：空网格
	fmt.Println(numIslands([][]byte{})) // 输出: 0

	// 单岛屿测试
	grid3 := [][]byte{{'1'}}
	fmt.Println(numIslands(grid3)) // 输出: 1
}
