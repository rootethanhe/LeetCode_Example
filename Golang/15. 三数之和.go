package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	// 先排序
	sort.Ints(nums)
	result := [][]int{}
	n := len(nums)

	for i := 0; i < n-2; i++ {
		// 跳过重复的第一个数
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 剪枝：若第一个数已大于0，后续无解
		if nums[i] > 0 {
			break
		}

		left, right := i+1, n-1
		target := -nums[i]
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// 跳过重复的左右指针元素
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

	return result
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // [[-1,-1,2],[-1,0,1]]
	fmt.Println(threeSum([]int{0, 1, 1}))             // []
	fmt.Println(threeSum([]int{0, 0, 0}))             // [[0,0,0]]
}
