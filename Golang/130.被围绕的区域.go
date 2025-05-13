package main

import "fmt"

// solve 函数用于捕获被围绕的区域
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	// 标记边缘 'O' 及其相连区域
	for i := 0; i < m; i++ {
		dfs(board, i, 0)
		dfs(board, i, n-1)
	}
	for j := 0; j < n; j++ {
		dfs(board, 0, j)
		dfs(board, m-1, j)
	}
	// 替换剩余 'O' 为 'X'，还原 '#' 为 'O'
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

// dfs 函数用于深度优先搜索标记相连的 'O' 区域
func dfs(board [][]byte, i, j int) {
	m, n := len(board), len(board[0])
	if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != 'O' {
		return
	}
	board[i][j] = '#'
	dfs(board, i+1, j)
	dfs(board, i-1, j)
	dfs(board, i, j+1)
	dfs(board, i, j-1)
}

func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
