package main

import "fmt"

func gameOfLife(board [][]int) {
	// 1. 定义八邻域方向向量
	dirs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	m, n := len(board), len(board[0])

	// 2. 状态标记阶段
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 统计存活邻居（包括标记为2的存活细胞）
			liveCount := 0
			for _, d := range dirs {
				ni, nj := i+d[0], j+d[1]
				if ni >= 0 && ni < m && nj >= 0 && nj < n {
					// 当前状态为存活：1或2
					if board[ni][nj] == 1 || board[ni][nj] == 2 {
						liveCount++
					}
				}
			}

			// 应用生存定律
			switch {
			case board[i][j] == 1 && (liveCount < 2 || liveCount > 3):
				board[i][j] = 2 // 存活→死亡
			case board[i][j] == 0 && liveCount == 3:
				board[i][j] = 3 // 死亡→存活
			}
		}
	}

	// 3. 状态转换阶段
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 2 {
				board[i][j] = 0
			} else if board[i][j] == 3 {
				board[i][j] = 1
			}
		}
	}
}

func main() {
	// 示例1
	board1 := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(board1)
	fmt.Println(board1) // [[0 0 0] [1 0 1] [0 1 1] [0 1 0]]

	// 示例2
	board2 := [][]int{{1, 1}, {1, 0}}
	gameOfLife(board2)
	fmt.Println(board2) // [[1 1] [1 1]]

	// 滑翔机测试
	board3 := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
	}
	gameOfLife(board3)
	fmt.Println(board3) // [[0 0 0] [1 0 1] [0 1 1]]
}
