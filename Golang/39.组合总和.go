package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates) // 排序以便剪枝
	var res [][]int
	var path []int

	var backtrack func(start, remain int) // start: 起始索引  remain: 当前剩余值
	backtrack = func(start, remain int) {
		if remain == 0 { // 找到合法组合
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := start; i < len(candidates); i++ {
			if candidates[i] > remain { // 剪枝：后续元素更大，直接终止
				break
			}
			path = append(path, candidates[i]) // 选择当前元素
			backtrack(i, remain-candidates[i]) // 递归(允许重复选择)
			path = path[:len(path)-1]          // 回溯
		}
	}

	backtrack(0, target) // 从索引0开始，目标为完整target
	return res
}

func main() {
	// 示例1
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7)) // [[2 2 3] [7]]
	// 示例2
	fmt.Println(combinationSum([]int{2, 3, 5}, 8)) // [[2 2 2 2] [2 3 3] [3 5]]
	// 示例3
	fmt.Println(combinationSum([]int{2}, 1)) // []
}
