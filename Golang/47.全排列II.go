package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums) // 排序预处理(剪枝基础)
	n := len(nums)
	res := [][]int{}
	used := make([]bool, n)

	var backtrack func(path []int)
	backtrack = func(path []int) {
		if len(path) == n { // 满足条件时记录结果
			tmp := make([]int, n)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := 0; i < n; i++ {
			// 剪枝条件：已使用或与前一个相同且前一个未使用（去重核心逻辑）
			if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
				continue
			}

			used[i] = true
			path = append(path, nums[i])
			backtrack(path)
			path = path[:len(path)-1] // 回溯
			used[i] = false
		}
	}

	backtrack([]int{})
	return res
}

func main() {
	// 输入：nums = [1,1,2]  输出：[[1,1,2], [1,2,1], [2,1,1]]
	fmt.Println(permuteUnique([]int{1, 1, 2}))

	// 输入：nums = [1,2,3]  输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
	fmt.Println(permuteUnique([]int{1, 2, 3}))
}
