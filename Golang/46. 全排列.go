package main

import "fmt"

func permute(nums []int) [][]int {
	var res [][]int
	visited := make([]bool, len(nums))
	var path []int

	var backtrack func()
	backtrack = func() {
		// 终止条件: 路径长度等于数组长度
		if len(path) == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		// 遍历选择列表
		for i := 0; i < len(nums); i++ {
			if visited[i] { // 跳过已使用的元素
				continue
			}

			// 做选择
			visited[i] = true
			path = append(path, nums[i])

			// 进入下一层决策树
			backtrack()

			// 撤销选择
			visited[i] = false
			path = path[:len(path)-1]
		}
	}

	backtrack()
	return res
}

func main() {
	// 示例1：[1,2,3]
	fmt.Println(permute([]int{1, 2, 3})) // 输出6种排列组合

	// 示例2：[0,1]
	fmt.Println(permute([]int{0, 1})) // 输出[[0,1],[1,0]]

	// 示例3：[1]
	fmt.Println(permute([]int{1})) // 输出[[1]]
}
