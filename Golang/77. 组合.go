package main

import "fmt"

func combine(n int, k int) [][]int {
	var res [][]int
	var backtrack func(start int, path []int)

	backtrack = func(start int, path []int) {
		// 当路径长度等于k时，记录结果
		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path) // 复制当前路径避免后续修改
			res = append(res, tmp)
			return
		}

		// 计算当前可选的数字范围上限(减枝优化)
		// 剩余需要选的数量 = k - len(path)
		// 最大起始数字应满足：剩余数字数量 >= 剩余需要选的数量 → i <= n - (k - len(path)) + 1
		maxStart := n - (k - len(path)) + 1

		for i := start; i <= maxStart; i++ {
			path = append(path, i)    // 选择当前数字
			backtrack(i+1, path)      // 递归下一层，起始位置+1避免重复
			path = path[:len(path)-1] // 回溯，撤销选择
		}
	}

	backtrack(1, []int{}) // 从数字1开始回溯
	return res
}

func main() {
	fmt.Println(combine(4, 2)) // 输出：[[1 2] [1 3] [1 4] [2 3] [2 4] [3 4]]
	fmt.Println(combine(1, 1)) // 输出：[[1]]
}
