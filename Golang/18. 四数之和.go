package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	var result [][]int
	n := len(nums)
	if n < 4 {
		return result
	}
	// 先排序
	sort.Ints(nums)

	for i := 0; i < n-3; i++ {
		// 跳过重复的i
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 剪枝：若当前i对于的最小四位数已超过target，则后续无需遍历
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		// 剪枝：若当前i对应的最大四位数和仍不足target，跳过本次循环
		if nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			// 跳过重复的j
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left, right := j+1, n-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					// 跳过重复的left和right
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return result
}

func main() {
	// [-2,-1,0,0,1,2]
	// 输入：nums = [1,0,-1,0,-2,2], target = 0  输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	// 输入：nums = [2,2,2,2,2], target = 8  输出：[[2,2,2,2]]
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}
