package main

import "fmt"

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	// 遍历网格的每个起点
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] && dfs(board, word, i, j, 0) {
				return true
			}
		}
	}

	return false
}

func dfs(board [][]byte, word string, i, j int, index int) bool {
	// 成功匹配整个单词
	if index == len(word) {
		return true
	}

	// 边界检查或字符串不匹配
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[index] {
		return false
	}

	// 标记当前单元格(避免重复访问)
	tmp := board[i][j]
	board[i][j] = '*'

	// 四向递归搜索(上下左右)
	found := dfs(board, word, i-1, j, index+1) || // 上
		dfs(board, word, i+1, j, index+1) || // 下
		dfs(board, word, i, j-1, index+1) || // 左
		dfs(board, word, i, j+1, index+1) // 右

	// 回溯: 恢复原始字符串
	board[i][j] = tmp
	return found
}

func main() {
	board1 := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board1, "ABCCED")) // true (示例1)
	fmt.Println(exist(board1, "SEE"))    // true (示例2)
	fmt.Println(exist(board1, "ABCB"))   // false (示例3)

	board2 := [][]byte{{'a'}}
	fmt.Println(exist(board2, "a")) // true (单字符)
}
